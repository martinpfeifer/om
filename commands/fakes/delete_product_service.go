// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type DeleteProductService struct {
	DeleteAvailableProductsStub        func(api.DeleteAvailableProductsInput) error
	deleteAvailableProductsMutex       sync.RWMutex
	deleteAvailableProductsArgsForCall []struct {
		arg1 api.DeleteAvailableProductsInput
	}
	deleteAvailableProductsReturns struct {
		result1 error
	}
	deleteAvailableProductsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeleteProductService) DeleteAvailableProducts(arg1 api.DeleteAvailableProductsInput) error {
	fake.deleteAvailableProductsMutex.Lock()
	ret, specificReturn := fake.deleteAvailableProductsReturnsOnCall[len(fake.deleteAvailableProductsArgsForCall)]
	fake.deleteAvailableProductsArgsForCall = append(fake.deleteAvailableProductsArgsForCall, struct {
		arg1 api.DeleteAvailableProductsInput
	}{arg1})
	fake.recordInvocation("DeleteAvailableProducts", []interface{}{arg1})
	fake.deleteAvailableProductsMutex.Unlock()
	if fake.DeleteAvailableProductsStub != nil {
		return fake.DeleteAvailableProductsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteAvailableProductsReturns
	return fakeReturns.result1
}

func (fake *DeleteProductService) DeleteAvailableProductsCallCount() int {
	fake.deleteAvailableProductsMutex.RLock()
	defer fake.deleteAvailableProductsMutex.RUnlock()
	return len(fake.deleteAvailableProductsArgsForCall)
}

func (fake *DeleteProductService) DeleteAvailableProductsCalls(stub func(api.DeleteAvailableProductsInput) error) {
	fake.deleteAvailableProductsMutex.Lock()
	defer fake.deleteAvailableProductsMutex.Unlock()
	fake.DeleteAvailableProductsStub = stub
}

func (fake *DeleteProductService) DeleteAvailableProductsArgsForCall(i int) api.DeleteAvailableProductsInput {
	fake.deleteAvailableProductsMutex.RLock()
	defer fake.deleteAvailableProductsMutex.RUnlock()
	argsForCall := fake.deleteAvailableProductsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeleteProductService) DeleteAvailableProductsReturns(result1 error) {
	fake.deleteAvailableProductsMutex.Lock()
	defer fake.deleteAvailableProductsMutex.Unlock()
	fake.DeleteAvailableProductsStub = nil
	fake.deleteAvailableProductsReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeleteProductService) DeleteAvailableProductsReturnsOnCall(i int, result1 error) {
	fake.deleteAvailableProductsMutex.Lock()
	defer fake.deleteAvailableProductsMutex.Unlock()
	fake.DeleteAvailableProductsStub = nil
	if fake.deleteAvailableProductsReturnsOnCall == nil {
		fake.deleteAvailableProductsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAvailableProductsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeleteProductService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteAvailableProductsMutex.RLock()
	defer fake.deleteAvailableProductsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeleteProductService) recordInvocation(key string, args []interface{}) {
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
