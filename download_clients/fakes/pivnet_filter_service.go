// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/om/download_clients"
)

type PivnetFilter struct {
	ProductFileKeysByGlobsStub        func([]pivnet.ProductFile, []string) ([]pivnet.ProductFile, error)
	productFileKeysByGlobsMutex       sync.RWMutex
	productFileKeysByGlobsArgsForCall []struct {
		arg1 []pivnet.ProductFile
		arg2 []string
	}
	productFileKeysByGlobsReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFileKeysByGlobsReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	ReleasesByVersionStub        func([]pivnet.Release, string) ([]pivnet.Release, error)
	releasesByVersionMutex       sync.RWMutex
	releasesByVersionArgsForCall []struct {
		arg1 []pivnet.Release
		arg2 string
	}
	releasesByVersionReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	releasesByVersionReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PivnetFilter) ProductFileKeysByGlobs(arg1 []pivnet.ProductFile, arg2 []string) ([]pivnet.ProductFile, error) {
	var arg1Copy []pivnet.ProductFile
	if arg1 != nil {
		arg1Copy = make([]pivnet.ProductFile, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.productFileKeysByGlobsMutex.Lock()
	ret, specificReturn := fake.productFileKeysByGlobsReturnsOnCall[len(fake.productFileKeysByGlobsArgsForCall)]
	fake.productFileKeysByGlobsArgsForCall = append(fake.productFileKeysByGlobsArgsForCall, struct {
		arg1 []pivnet.ProductFile
		arg2 []string
	}{arg1Copy, arg2Copy})
	fake.recordInvocation("ProductFileKeysByGlobs", []interface{}{arg1Copy, arg2Copy})
	fake.productFileKeysByGlobsMutex.Unlock()
	if fake.ProductFileKeysByGlobsStub != nil {
		return fake.ProductFileKeysByGlobsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productFileKeysByGlobsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetFilter) ProductFileKeysByGlobsCallCount() int {
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	return len(fake.productFileKeysByGlobsArgsForCall)
}

func (fake *PivnetFilter) ProductFileKeysByGlobsCalls(stub func([]pivnet.ProductFile, []string) ([]pivnet.ProductFile, error)) {
	fake.productFileKeysByGlobsMutex.Lock()
	defer fake.productFileKeysByGlobsMutex.Unlock()
	fake.ProductFileKeysByGlobsStub = stub
}

func (fake *PivnetFilter) ProductFileKeysByGlobsArgsForCall(i int) ([]pivnet.ProductFile, []string) {
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	argsForCall := fake.productFileKeysByGlobsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PivnetFilter) ProductFileKeysByGlobsReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.productFileKeysByGlobsMutex.Lock()
	defer fake.productFileKeysByGlobsMutex.Unlock()
	fake.ProductFileKeysByGlobsStub = nil
	fake.productFileKeysByGlobsReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetFilter) ProductFileKeysByGlobsReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.productFileKeysByGlobsMutex.Lock()
	defer fake.productFileKeysByGlobsMutex.Unlock()
	fake.ProductFileKeysByGlobsStub = nil
	if fake.productFileKeysByGlobsReturnsOnCall == nil {
		fake.productFileKeysByGlobsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFileKeysByGlobsReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetFilter) ReleasesByVersion(arg1 []pivnet.Release, arg2 string) ([]pivnet.Release, error) {
	var arg1Copy []pivnet.Release
	if arg1 != nil {
		arg1Copy = make([]pivnet.Release, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.releasesByVersionMutex.Lock()
	ret, specificReturn := fake.releasesByVersionReturnsOnCall[len(fake.releasesByVersionArgsForCall)]
	fake.releasesByVersionArgsForCall = append(fake.releasesByVersionArgsForCall, struct {
		arg1 []pivnet.Release
		arg2 string
	}{arg1Copy, arg2})
	fake.recordInvocation("ReleasesByVersion", []interface{}{arg1Copy, arg2})
	fake.releasesByVersionMutex.Unlock()
	if fake.ReleasesByVersionStub != nil {
		return fake.ReleasesByVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releasesByVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetFilter) ReleasesByVersionCallCount() int {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return len(fake.releasesByVersionArgsForCall)
}

func (fake *PivnetFilter) ReleasesByVersionCalls(stub func([]pivnet.Release, string) ([]pivnet.Release, error)) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = stub
}

func (fake *PivnetFilter) ReleasesByVersionArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	argsForCall := fake.releasesByVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PivnetFilter) ReleasesByVersionReturns(result1 []pivnet.Release, result2 error) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = nil
	fake.releasesByVersionReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetFilter) ReleasesByVersionReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = nil
	if fake.releasesByVersionReturnsOnCall == nil {
		fake.releasesByVersionReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.releasesByVersionReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PivnetFilter) recordInvocation(key string, args []interface{}) {
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

var _ download_clients.PivnetFilter = new(PivnetFilter)
