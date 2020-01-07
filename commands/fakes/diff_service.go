// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type BoshDiffService struct {
	DirectorDiffStub        func() (api.DirectorDiff, error)
	directorDiffMutex       sync.RWMutex
	directorDiffArgsForCall []struct {
	}
	directorDiffReturns struct {
		result1 api.DirectorDiff
		result2 error
	}
	directorDiffReturnsOnCall map[int]struct {
		result1 api.DirectorDiff
		result2 error
	}
	ListStagedProductsStub        func() (api.StagedProductsOutput, error)
	listStagedProductsMutex       sync.RWMutex
	listStagedProductsArgsForCall []struct {
	}
	listStagedProductsReturns struct {
		result1 api.StagedProductsOutput
		result2 error
	}
	listStagedProductsReturnsOnCall map[int]struct {
		result1 api.StagedProductsOutput
		result2 error
	}
	ProductDiffStub        func(string) (api.ProductDiff, error)
	productDiffMutex       sync.RWMutex
	productDiffArgsForCall []struct {
		arg1 string
	}
	productDiffReturns struct {
		result1 api.ProductDiff
		result2 error
	}
	productDiffReturnsOnCall map[int]struct {
		result1 api.ProductDiff
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BoshDiffService) DirectorDiff() (api.DirectorDiff, error) {
	fake.directorDiffMutex.Lock()
	ret, specificReturn := fake.directorDiffReturnsOnCall[len(fake.directorDiffArgsForCall)]
	fake.directorDiffArgsForCall = append(fake.directorDiffArgsForCall, struct {
	}{})
	fake.recordInvocation("DirectorDiff", []interface{}{})
	fake.directorDiffMutex.Unlock()
	if fake.DirectorDiffStub != nil {
		return fake.DirectorDiffStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.directorDiffReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BoshDiffService) DirectorDiffCallCount() int {
	fake.directorDiffMutex.RLock()
	defer fake.directorDiffMutex.RUnlock()
	return len(fake.directorDiffArgsForCall)
}

func (fake *BoshDiffService) DirectorDiffCalls(stub func() (api.DirectorDiff, error)) {
	fake.directorDiffMutex.Lock()
	defer fake.directorDiffMutex.Unlock()
	fake.DirectorDiffStub = stub
}

func (fake *BoshDiffService) DirectorDiffReturns(result1 api.DirectorDiff, result2 error) {
	fake.directorDiffMutex.Lock()
	defer fake.directorDiffMutex.Unlock()
	fake.DirectorDiffStub = nil
	fake.directorDiffReturns = struct {
		result1 api.DirectorDiff
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) DirectorDiffReturnsOnCall(i int, result1 api.DirectorDiff, result2 error) {
	fake.directorDiffMutex.Lock()
	defer fake.directorDiffMutex.Unlock()
	fake.DirectorDiffStub = nil
	if fake.directorDiffReturnsOnCall == nil {
		fake.directorDiffReturnsOnCall = make(map[int]struct {
			result1 api.DirectorDiff
			result2 error
		})
	}
	fake.directorDiffReturnsOnCall[i] = struct {
		result1 api.DirectorDiff
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) ListStagedProducts() (api.StagedProductsOutput, error) {
	fake.listStagedProductsMutex.Lock()
	ret, specificReturn := fake.listStagedProductsReturnsOnCall[len(fake.listStagedProductsArgsForCall)]
	fake.listStagedProductsArgsForCall = append(fake.listStagedProductsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListStagedProducts", []interface{}{})
	fake.listStagedProductsMutex.Unlock()
	if fake.ListStagedProductsStub != nil {
		return fake.ListStagedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedProductsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BoshDiffService) ListStagedProductsCallCount() int {
	fake.listStagedProductsMutex.RLock()
	defer fake.listStagedProductsMutex.RUnlock()
	return len(fake.listStagedProductsArgsForCall)
}

func (fake *BoshDiffService) ListStagedProductsCalls(stub func() (api.StagedProductsOutput, error)) {
	fake.listStagedProductsMutex.Lock()
	defer fake.listStagedProductsMutex.Unlock()
	fake.ListStagedProductsStub = stub
}

func (fake *BoshDiffService) ListStagedProductsReturns(result1 api.StagedProductsOutput, result2 error) {
	fake.listStagedProductsMutex.Lock()
	defer fake.listStagedProductsMutex.Unlock()
	fake.ListStagedProductsStub = nil
	fake.listStagedProductsReturns = struct {
		result1 api.StagedProductsOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) ListStagedProductsReturnsOnCall(i int, result1 api.StagedProductsOutput, result2 error) {
	fake.listStagedProductsMutex.Lock()
	defer fake.listStagedProductsMutex.Unlock()
	fake.ListStagedProductsStub = nil
	if fake.listStagedProductsReturnsOnCall == nil {
		fake.listStagedProductsReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsOutput
			result2 error
		})
	}
	fake.listStagedProductsReturnsOnCall[i] = struct {
		result1 api.StagedProductsOutput
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) ProductDiff(arg1 string) (api.ProductDiff, error) {
	fake.productDiffMutex.Lock()
	ret, specificReturn := fake.productDiffReturnsOnCall[len(fake.productDiffArgsForCall)]
	fake.productDiffArgsForCall = append(fake.productDiffArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ProductDiff", []interface{}{arg1})
	fake.productDiffMutex.Unlock()
	if fake.ProductDiffStub != nil {
		return fake.ProductDiffStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productDiffReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BoshDiffService) ProductDiffCallCount() int {
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	return len(fake.productDiffArgsForCall)
}

func (fake *BoshDiffService) ProductDiffCalls(stub func(string) (api.ProductDiff, error)) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = stub
}

func (fake *BoshDiffService) ProductDiffArgsForCall(i int) string {
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	argsForCall := fake.productDiffArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BoshDiffService) ProductDiffReturns(result1 api.ProductDiff, result2 error) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = nil
	fake.productDiffReturns = struct {
		result1 api.ProductDiff
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) ProductDiffReturnsOnCall(i int, result1 api.ProductDiff, result2 error) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = nil
	if fake.productDiffReturnsOnCall == nil {
		fake.productDiffReturnsOnCall = make(map[int]struct {
			result1 api.ProductDiff
			result2 error
		})
	}
	fake.productDiffReturnsOnCall[i] = struct {
		result1 api.ProductDiff
		result2 error
	}{result1, result2}
}

func (fake *BoshDiffService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.directorDiffMutex.RLock()
	defer fake.directorDiffMutex.RUnlock()
	fake.listStagedProductsMutex.RLock()
	defer fake.listStagedProductsMutex.RUnlock()
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BoshDiffService) recordInvocation(key string, args []interface{}) {
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
