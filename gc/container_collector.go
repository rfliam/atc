package gc

import (
	"context"
	"fmt"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/metric"
	"github.com/concourse/atc/worker"
	multierror "github.com/hashicorp/go-multierror"
)

const HijackedContainerTimeout = 5 * time.Minute

type containerCollector struct {
	containerRepository db.ContainerRepository
	jobRunner           WorkerJobRunner
}

func NewContainerCollector(
	containerRepository db.ContainerRepository,
	jobRunner WorkerJobRunner,
) Collector {
	return &containerCollector{
		containerRepository: containerRepository,
		jobRunner:           jobRunner,
	}
}

type job struct {
	JobName string
	RunFunc func(worker.Worker)
}

func (j *job) Name() string {
	return j.JobName
}

func (j *job) Run(w worker.Worker) {
	j.RunFunc(w)
}

func (c *containerCollector) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("container-collector")

	logger.Debug("start")
	defer logger.Debug("done")

	var errs error

	err := c.cleanupOrphanedContainers(logger.Session("orphaned-containers"))
	if err != nil {
		errs = multierror.Append(errs, err)
		logger.Error("failed-to-clean-up-orphaned-containers", err)
	}

	err = c.cleanupFailedContainers(logger.Session("failed-containers"))
	if err != nil {
		errs = multierror.Append(errs, err)
		logger.Error("failed-to-clean-up-failed-containers", err)
	}

	return errs
}

func (c *containerCollector) cleanupFailedContainers(logger lager.Logger) error {
	failedContainers, err := c.containerRepository.FindFailedContainers()
	if err != nil {
		logger.Error("failed-to-find-failed-containers-for-deletion", err)
		return err
	}

	failedContainerHandles := []string{}
	var failedContainerstoDestroy = []destroyableContainer{}

	if len(failedContainers) > 0 {
		for _, container := range failedContainers {
			failedContainerHandles = append(failedContainerHandles, container.Handle())
			failedContainerstoDestroy = append(failedContainerstoDestroy, container)
		}
	}

	if len(failedContainerHandles) > 0 {
		logger.Debug("found-failed-containers-for-deletion", lager.Data{
			"failed-containers": failedContainerHandles,
		})
	}

	metric.FailedContainersToBeGarbageCollected{
		Containers: len(failedContainerHandles),
	}.Emit(logger)

	destroyDBContainers(logger, failedContainerstoDestroy)

	return nil
}

func (c *containerCollector) cleanupOrphanedContainers(logger lager.Logger) error {
	creatingContainers, createdContainers, destroyingContainers, err := c.containerRepository.FindOrphanedContainers()

	if err != nil {
		logger.Error("failed-to-get-orphaned-containers-for-deletion", err)
		return err
	}

	creatingContainerHandles := []string{}
	createdContainerHandles := []string{}
	destroyingContainerHandles := []string{}

	if len(creatingContainers) > 0 {
		for _, container := range creatingContainers {
			creatingContainerHandles = append(creatingContainerHandles, container.Handle())
		}
	}

	if len(createdContainers) > 0 {
		for _, container := range createdContainers {
			createdContainerHandles = append(createdContainerHandles, container.Handle())
		}
	}

	if len(destroyingContainers) > 0 {
		for _, container := range destroyingContainers {
			destroyingContainerHandles = append(destroyingContainerHandles, container.Handle())
		}
	}

	if len(createdContainerHandles) > 0 || len(createdContainerHandles) > 0 || len(destroyingContainerHandles) > 0 {
		logger.Debug("found-orphaned-containers-for-deletion", lager.Data{
			"creating-containers":   creatingContainerHandles,
			"created-containers":    createdContainerHandles,
			"destroying-containers": destroyingContainerHandles,
		})
	}

	metric.CreatingContainersToBeGarbageCollected{
		Containers: len(creatingContainerHandles),
	}.Emit(logger)

	metric.CreatedContainersToBeGarbageCollected{
		Containers: len(createdContainerHandles),
	}.Emit(logger)

	metric.DestroyingContainersToBeGarbageCollected{
		Containers: len(destroyingContainerHandles),
	}.Emit(logger)

	var workerCreatedContainers = make(map[string][]db.CreatedContainer)

	for _, createdContainer := range createdContainers {
		containers, ok := workerCreatedContainers[createdContainer.WorkerName()]
		if ok {
			// update existing array
			containers = append(containers, createdContainer)
			workerCreatedContainers[createdContainer.WorkerName()] = containers
		} else {
			// create new array
			workerCreatedContainers[createdContainer.WorkerName()] = []db.CreatedContainer{createdContainer}
		}
	}

	for worker, createdContainers := range workerCreatedContainers {
		// prevent closure from capturing last value of loop
		c.jobRunner.Try(logger,
			worker,
			&job{
				JobName: fmt.Sprintf("created-containers-%d", len(createdContainers)),
				RunFunc: destroyCreatedContainers(logger, createdContainers),
			},
		)
	}

	var workerContainers = make(map[string][]db.DestroyingContainer)

	for _, destroyingContainer := range destroyingContainers {
		containers, ok := workerContainers[destroyingContainer.WorkerName()]
		if ok {
			// update existing array
			containers = append(containers, destroyingContainer)
			workerContainers[destroyingContainer.WorkerName()] = containers
		} else {
			// create new array
			workerContainers[destroyingContainer.WorkerName()] = []db.DestroyingContainer{destroyingContainer}
		}
	}

	for worker, destroyingContainers := range workerContainers {
		c.jobRunner.Try(logger,
			worker,
			&job{
				JobName: fmt.Sprintf("destroying-containers-%d", len(destroyingContainers)),
				RunFunc: destroyDestroyingContainers(logger, destroyingContainers),
			},
		)
	}
	return nil
}

func destroyCreatedContainers(logger lager.Logger, containers []db.CreatedContainer) func(worker.Worker) {
	return func(workerClient worker.Worker) {
		destroyingContainers := []db.DestroyingContainer{}

		for _, container := range containers {

			var destroyingContainer db.DestroyingContainer
			var err error

			if container.IsHijacked() {
				cLog := logger.Session("mark-hijacked-container", lager.Data{
					"container": container.Handle(),
					"worker":    workerClient.Name(),
				})

				destroyingContainer, err = markHijackedContainerAsDestroying(cLog, container, workerClient.GardenClient())
				if err != nil {
					cLog.Error("failed-to-transition", err)
					return
				}
			} else {
				cLog := logger.Session("mark-created-as-destroying", lager.Data{
					"container": container.Handle(),
					"worker":    workerClient.Name(),
				})

				destroyingContainer, err = container.Destroying()
				if err != nil {
					cLog.Error("failed-to-transition", err)
					return
				}
			}
			if destroyingContainer != nil {
				destroyingContainers = append(destroyingContainers, destroyingContainer)
			}

		}
		tryToDestroyContainers(logger.Session("destroy-containers"), destroyingContainers, workerClient)
	}
}

func destroyDestroyingContainers(logger lager.Logger, containers []db.DestroyingContainer) func(worker.Worker) {
	return func(workerClient worker.Worker) {
		cLog := logger.Session("destroy-containers-on-worker", lager.Data{
			"worker": workerClient.Name(),
		})
		tryToDestroyContainers(cLog, containers, workerClient)
	}
}

func markHijackedContainerAsDestroying(
	logger lager.Logger,
	hijackedContainer db.CreatedContainer,
	gardenClient garden.Client,
) (db.DestroyingContainer, error) {

	gardenContainer, found, err := findContainer(gardenClient, hijackedContainer.Handle())
	if err != nil {
		logger.Error("failed-to-lookup-container-in-garden", err)
		return nil, err
	}

	if !found {
		logger.Debug("hijacked-container-not-found-in-garden")

		destroyingContainer, err := hijackedContainer.Destroying()
		if err != nil {
			logger.Error("failed-to-mark-container-as-destroying", err)
			return nil, err
		}
		return destroyingContainer, nil
	}

	err = gardenContainer.SetGraceTime(HijackedContainerTimeout)
	if err != nil {
		logger.Error("failed-to-set-grace-time-on-hijacked-container", err)
		return nil, err
	}

	_, err = hijackedContainer.Discontinue()
	if err != nil {
		logger.Error("failed-to-mark-container-as-destroying", err)
		return nil, err
	}

	return nil, nil
}

func tryToDestroyContainers(
	logger lager.Logger,
	containers []db.DestroyingContainer,
	workerClient worker.Worker,
) {
	logger.Debug("start")
	defer logger.Debug("done")

	gardenDeleteHandles := []string{}
	gardenDeleteContainers := []destroyableContainer{}

	dbDeleteContainers := []destroyableContainer{}

	gardenClient := workerClient.GardenClient()
	reaperClient := workerClient.ReaperClient()

	for _, container := range containers {
		if container.IsDiscontinued() {
			cLog := logger.Session("discontinued", lager.Data{"handle": container.Handle()})

			_, found, err := findContainer(gardenClient, container.Handle())
			if err != nil {
				cLog.Error("failed-to-lookup-container-in-garden", err)
			}

			if found {
				cLog.Debug("still-present-in-garden")
			} else {
				cLog.Debug("container-no-longer-present-in-garden")
				dbDeleteContainers = append(dbDeleteContainers, container)
			}
		} else {
			gardenDeleteHandles = append(gardenDeleteHandles, container.Handle())
			gardenDeleteContainers = append(gardenDeleteContainers, container)
		}
	}

	if len(gardenDeleteHandles) > 0 {
		err := reaperClient.DestroyContainers(gardenDeleteHandles)
		if err != nil {
			logger.Error("failed-to-destroy-garden-containers", err, lager.Data{"handles": gardenDeleteHandles})
		} else {
			logger.Debug("completed-destroyed-in-garden", lager.Data{"handles": gardenDeleteHandles})
			dbDeleteContainers = append(dbDeleteContainers, gardenDeleteContainers...)
		}
	}

	destroyDBContainers(logger.Session("destroy-in-db"), dbDeleteContainers)
}

type destroyableContainer interface {
	Handle() string
	Destroy() (bool, error)
}

func destroyDBContainers(logger lager.Logger, dbContainers []destroyableContainer) {
	if len(dbContainers) == 0 {
		return
	}

	logger.Debug("start", lager.Data{"length": len(dbContainers)})
	defer logger.Debug("done")

	for _, dbContainer := range dbContainers {
		dLog := logger.Session("destroy-container", lager.Data{
			"container": dbContainer.Handle(),
		})

		destroyed, err := dbContainer.Destroy()
		if err != nil {
			dLog.Error("failed-to-destroy", err)
			continue
		}

		if !destroyed {
			dLog.Info("container-not-destroyed")
			continue
		}

		metric.ContainersDeleted.Inc()

		logger.Debug("destroyed")
	}
}

func findContainer(gardenClient garden.Client, handle string) (garden.Container, bool, error) {
	gardenContainer, err := gardenClient.Lookup(handle)
	if err != nil {
		if _, ok := err.(garden.ContainerNotFoundError); ok {
			return nil, false, nil
		}
		return nil, false, err
	}
	return gardenContainer, true, nil
}
