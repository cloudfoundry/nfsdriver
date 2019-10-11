// Code generated by counterfeiter. DO NOT EDIT.
package time_fake

import (
	"sync"
	"time"

	"code.cloudfoundry.org/goshims/timeshim"
)

type FakeTime struct {
	NowStub        func() time.Time
	nowMutex       sync.RWMutex
	nowArgsForCall []struct {
	}
	nowReturns struct {
		result1 time.Time
	}
	nowReturnsOnCall map[int]struct {
		result1 time.Time
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTime) Now() time.Time {
	fake.nowMutex.Lock()
	ret, specificReturn := fake.nowReturnsOnCall[len(fake.nowArgsForCall)]
	fake.nowArgsForCall = append(fake.nowArgsForCall, struct {
	}{})
	fake.recordInvocation("Now", []interface{}{})
	fake.nowMutex.Unlock()
	if fake.NowStub != nil {
		return fake.NowStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nowReturns
	return fakeReturns.result1
}

func (fake *FakeTime) NowCallCount() int {
	fake.nowMutex.RLock()
	defer fake.nowMutex.RUnlock()
	return len(fake.nowArgsForCall)
}

func (fake *FakeTime) NowCalls(stub func() time.Time) {
	fake.nowMutex.Lock()
	defer fake.nowMutex.Unlock()
	fake.NowStub = stub
}

func (fake *FakeTime) NowReturns(result1 time.Time) {
	fake.nowMutex.Lock()
	defer fake.nowMutex.Unlock()
	fake.NowStub = nil
	fake.nowReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeTime) NowReturnsOnCall(i int, result1 time.Time) {
	fake.nowMutex.Lock()
	defer fake.nowMutex.Unlock()
	fake.NowStub = nil
	if fake.nowReturnsOnCall == nil {
		fake.nowReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.nowReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeTime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nowMutex.RLock()
	defer fake.nowMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTime) recordInvocation(key string, args []interface{}) {
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

var _ timeshim.Time = new(FakeTime)
