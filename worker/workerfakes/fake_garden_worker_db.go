// This file was generated by counterfeiter
package workerfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/lock"
	"github.com/concourse/atc/worker"
)

type FakeGardenWorkerDB struct {
	CreateContainerStub        func(container db.Container, ttl time.Duration, maxLifetime time.Duration, volumeHandles []string) (db.SavedContainer, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		container     db.Container
		ttl           time.Duration
		maxLifetime   time.Duration
		volumeHandles []string
	}
	createContainerReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	UpdateContainerTTLToBeRemovedStub        func(container db.Container, ttl time.Duration, maxLifetime time.Duration) (db.SavedContainer, error)
	updateContainerTTLToBeRemovedMutex       sync.RWMutex
	updateContainerTTLToBeRemovedArgsForCall []struct {
		container   db.Container
		ttl         time.Duration
		maxLifetime time.Duration
	}
	updateContainerTTLToBeRemovedReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	GetContainerStub        func(handle string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		handle string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	ReapContainerStub        func(string) error
	reapContainerMutex       sync.RWMutex
	reapContainerArgsForCall []struct {
		arg1 string
	}
	reapContainerReturns struct {
		result1 error
	}
	GetPipelineByIDStub        func(pipelineID int) (db.SavedPipeline, error)
	getPipelineByIDMutex       sync.RWMutex
	getPipelineByIDArgsForCall []struct {
		pipelineID int
	}
	getPipelineByIDReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	AcquireVolumeCreatingLockStub        func(lager.Logger, int) (lock.Lock, bool, error)
	acquireVolumeCreatingLockMutex       sync.RWMutex
	acquireVolumeCreatingLockArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	acquireVolumeCreatingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGardenWorkerDB) CreateContainer(container db.Container, ttl time.Duration, maxLifetime time.Duration, volumeHandles []string) (db.SavedContainer, error) {
	var volumeHandlesCopy []string
	if volumeHandles != nil {
		volumeHandlesCopy = make([]string, len(volumeHandles))
		copy(volumeHandlesCopy, volumeHandles)
	}
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		container     db.Container
		ttl           time.Duration
		maxLifetime   time.Duration
		volumeHandles []string
	}{container, ttl, maxLifetime, volumeHandlesCopy})
	fake.recordInvocation("CreateContainer", []interface{}{container, ttl, maxLifetime, volumeHandlesCopy})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(container, ttl, maxLifetime, volumeHandles)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) CreateContainerArgsForCall(i int) (db.Container, time.Duration, time.Duration, []string) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].container, fake.createContainerArgsForCall[i].ttl, fake.createContainerArgsForCall[i].maxLifetime, fake.createContainerArgsForCall[i].volumeHandles
}

func (fake *FakeGardenWorkerDB) CreateContainerReturns(result1 db.SavedContainer, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) UpdateContainerTTLToBeRemoved(container db.Container, ttl time.Duration, maxLifetime time.Duration) (db.SavedContainer, error) {
	fake.updateContainerTTLToBeRemovedMutex.Lock()
	fake.updateContainerTTLToBeRemovedArgsForCall = append(fake.updateContainerTTLToBeRemovedArgsForCall, struct {
		container   db.Container
		ttl         time.Duration
		maxLifetime time.Duration
	}{container, ttl, maxLifetime})
	fake.recordInvocation("UpdateContainerTTLToBeRemoved", []interface{}{container, ttl, maxLifetime})
	fake.updateContainerTTLToBeRemovedMutex.Unlock()
	if fake.UpdateContainerTTLToBeRemovedStub != nil {
		return fake.UpdateContainerTTLToBeRemovedStub(container, ttl, maxLifetime)
	} else {
		return fake.updateContainerTTLToBeRemovedReturns.result1, fake.updateContainerTTLToBeRemovedReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) UpdateContainerTTLToBeRemovedCallCount() int {
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	return len(fake.updateContainerTTLToBeRemovedArgsForCall)
}

func (fake *FakeGardenWorkerDB) UpdateContainerTTLToBeRemovedArgsForCall(i int) (db.Container, time.Duration, time.Duration) {
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	return fake.updateContainerTTLToBeRemovedArgsForCall[i].container, fake.updateContainerTTLToBeRemovedArgsForCall[i].ttl, fake.updateContainerTTLToBeRemovedArgsForCall[i].maxLifetime
}

func (fake *FakeGardenWorkerDB) UpdateContainerTTLToBeRemovedReturns(result1 db.SavedContainer, result2 error) {
	fake.UpdateContainerTTLToBeRemovedStub = nil
	fake.updateContainerTTLToBeRemovedReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) GetContainer(handle string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("GetContainer", []interface{}{handle})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(handle)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
	}
}

func (fake *FakeGardenWorkerDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].handle
}

func (fake *FakeGardenWorkerDB) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.recordInvocation("UpdateExpiresAtOnContainer", []interface{}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	} else {
		return fake.updateExpiresAtOnContainerReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) ReapContainer(arg1 string) error {
	fake.reapContainerMutex.Lock()
	fake.reapContainerArgsForCall = append(fake.reapContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReapContainer", []interface{}{arg1})
	fake.reapContainerMutex.Unlock()
	if fake.ReapContainerStub != nil {
		return fake.ReapContainerStub(arg1)
	} else {
		return fake.reapContainerReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) ReapContainerCallCount() int {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return len(fake.reapContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) ReapContainerArgsForCall(i int) string {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.reapContainerArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) ReapContainerReturns(result1 error) {
	fake.ReapContainerStub = nil
	fake.reapContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) GetPipelineByID(pipelineID int) (db.SavedPipeline, error) {
	fake.getPipelineByIDMutex.Lock()
	fake.getPipelineByIDArgsForCall = append(fake.getPipelineByIDArgsForCall, struct {
		pipelineID int
	}{pipelineID})
	fake.recordInvocation("GetPipelineByID", []interface{}{pipelineID})
	fake.getPipelineByIDMutex.Unlock()
	if fake.GetPipelineByIDStub != nil {
		return fake.GetPipelineByIDStub(pipelineID)
	} else {
		return fake.getPipelineByIDReturns.result1, fake.getPipelineByIDReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) GetPipelineByIDCallCount() int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return len(fake.getPipelineByIDArgsForCall)
}

func (fake *FakeGardenWorkerDB) GetPipelineByIDArgsForCall(i int) int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return fake.getPipelineByIDArgsForCall[i].pipelineID
}

func (fake *FakeGardenWorkerDB) GetPipelineByIDReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByIDStub = nil
	fake.getPipelineByIDReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) AcquireVolumeCreatingLock(arg1 lager.Logger, arg2 int) (lock.Lock, bool, error) {
	fake.acquireVolumeCreatingLockMutex.Lock()
	fake.acquireVolumeCreatingLockArgsForCall = append(fake.acquireVolumeCreatingLockArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcquireVolumeCreatingLock", []interface{}{arg1, arg2})
	fake.acquireVolumeCreatingLockMutex.Unlock()
	if fake.AcquireVolumeCreatingLockStub != nil {
		return fake.AcquireVolumeCreatingLockStub(arg1, arg2)
	} else {
		return fake.acquireVolumeCreatingLockReturns.result1, fake.acquireVolumeCreatingLockReturns.result2, fake.acquireVolumeCreatingLockReturns.result3
	}
}

func (fake *FakeGardenWorkerDB) AcquireVolumeCreatingLockCallCount() int {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return len(fake.acquireVolumeCreatingLockArgsForCall)
}

func (fake *FakeGardenWorkerDB) AcquireVolumeCreatingLockArgsForCall(i int) (lager.Logger, int) {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return fake.acquireVolumeCreatingLockArgsForCall[i].arg1, fake.acquireVolumeCreatingLockArgsForCall[i].arg2
}

func (fake *FakeGardenWorkerDB) AcquireVolumeCreatingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireVolumeCreatingLockStub = nil
	fake.acquireVolumeCreatingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGardenWorkerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGardenWorkerDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.GardenWorkerDB = new(FakeGardenWorkerDB)
