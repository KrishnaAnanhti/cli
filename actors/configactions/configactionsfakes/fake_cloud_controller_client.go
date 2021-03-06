// This file was generated by counterfeiter
package configactionsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actors/configactions"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
)

type FakeCloudControllerClient struct {
	TargetCFStub        func(APIURL string, skipSSLValidation bool) (ccv2.Warnings, error)
	targetCFMutex       sync.RWMutex
	targetCFArgsForCall []struct {
		APIURL            string
		skipSSLValidation bool
	}
	targetCFReturns struct {
		result1 ccv2.Warnings
		result2 error
	}
	APIStub        func() string
	aPIMutex       sync.RWMutex
	aPIArgsForCall []struct{}
	aPIReturns     struct {
		result1 string
	}
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct{}
	aPIVersionReturns     struct {
		result1 string
	}
	AuthorizationEndpointStub        func() string
	authorizationEndpointMutex       sync.RWMutex
	authorizationEndpointArgsForCall []struct{}
	authorizationEndpointReturns     struct {
		result1 string
	}
	DopplerEndpointStub        func() string
	dopplerEndpointMutex       sync.RWMutex
	dopplerEndpointArgsForCall []struct{}
	dopplerEndpointReturns     struct {
		result1 string
	}
	LoggregatorEndpointStub        func() string
	loggregatorEndpointMutex       sync.RWMutex
	loggregatorEndpointArgsForCall []struct{}
	loggregatorEndpointReturns     struct {
		result1 string
	}
	RoutingEndpointStub        func() string
	routingEndpointMutex       sync.RWMutex
	routingEndpointArgsForCall []struct{}
	routingEndpointReturns     struct {
		result1 string
	}
	TokenEndpointStub        func() string
	tokenEndpointMutex       sync.RWMutex
	tokenEndpointArgsForCall []struct{}
	tokenEndpointReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudControllerClient) TargetCF(APIURL string, skipSSLValidation bool) (ccv2.Warnings, error) {
	fake.targetCFMutex.Lock()
	fake.targetCFArgsForCall = append(fake.targetCFArgsForCall, struct {
		APIURL            string
		skipSSLValidation bool
	}{APIURL, skipSSLValidation})
	fake.recordInvocation("TargetCF", []interface{}{APIURL, skipSSLValidation})
	fake.targetCFMutex.Unlock()
	if fake.TargetCFStub != nil {
		return fake.TargetCFStub(APIURL, skipSSLValidation)
	} else {
		return fake.targetCFReturns.result1, fake.targetCFReturns.result2
	}
}

func (fake *FakeCloudControllerClient) TargetCFCallCount() int {
	fake.targetCFMutex.RLock()
	defer fake.targetCFMutex.RUnlock()
	return len(fake.targetCFArgsForCall)
}

func (fake *FakeCloudControllerClient) TargetCFArgsForCall(i int) (string, bool) {
	fake.targetCFMutex.RLock()
	defer fake.targetCFMutex.RUnlock()
	return fake.targetCFArgsForCall[i].APIURL, fake.targetCFArgsForCall[i].skipSSLValidation
}

func (fake *FakeCloudControllerClient) TargetCFReturns(result1 ccv2.Warnings, result2 error) {
	fake.TargetCFStub = nil
	fake.targetCFReturns = struct {
		result1 ccv2.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudControllerClient) API() string {
	fake.aPIMutex.Lock()
	fake.aPIArgsForCall = append(fake.aPIArgsForCall, struct{}{})
	fake.recordInvocation("API", []interface{}{})
	fake.aPIMutex.Unlock()
	if fake.APIStub != nil {
		return fake.APIStub()
	} else {
		return fake.aPIReturns.result1
	}
}

func (fake *FakeCloudControllerClient) APICallCount() int {
	fake.aPIMutex.RLock()
	defer fake.aPIMutex.RUnlock()
	return len(fake.aPIArgsForCall)
}

func (fake *FakeCloudControllerClient) APIReturns(result1 string) {
	fake.APIStub = nil
	fake.aPIReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	} else {
		return fake.aPIVersionReturns.result1
	}
}

func (fake *FakeCloudControllerClient) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeCloudControllerClient) APIVersionReturns(result1 string) {
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) AuthorizationEndpoint() string {
	fake.authorizationEndpointMutex.Lock()
	fake.authorizationEndpointArgsForCall = append(fake.authorizationEndpointArgsForCall, struct{}{})
	fake.recordInvocation("AuthorizationEndpoint", []interface{}{})
	fake.authorizationEndpointMutex.Unlock()
	if fake.AuthorizationEndpointStub != nil {
		return fake.AuthorizationEndpointStub()
	} else {
		return fake.authorizationEndpointReturns.result1
	}
}

func (fake *FakeCloudControllerClient) AuthorizationEndpointCallCount() int {
	fake.authorizationEndpointMutex.RLock()
	defer fake.authorizationEndpointMutex.RUnlock()
	return len(fake.authorizationEndpointArgsForCall)
}

func (fake *FakeCloudControllerClient) AuthorizationEndpointReturns(result1 string) {
	fake.AuthorizationEndpointStub = nil
	fake.authorizationEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) DopplerEndpoint() string {
	fake.dopplerEndpointMutex.Lock()
	fake.dopplerEndpointArgsForCall = append(fake.dopplerEndpointArgsForCall, struct{}{})
	fake.recordInvocation("DopplerEndpoint", []interface{}{})
	fake.dopplerEndpointMutex.Unlock()
	if fake.DopplerEndpointStub != nil {
		return fake.DopplerEndpointStub()
	} else {
		return fake.dopplerEndpointReturns.result1
	}
}

func (fake *FakeCloudControllerClient) DopplerEndpointCallCount() int {
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	return len(fake.dopplerEndpointArgsForCall)
}

func (fake *FakeCloudControllerClient) DopplerEndpointReturns(result1 string) {
	fake.DopplerEndpointStub = nil
	fake.dopplerEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) LoggregatorEndpoint() string {
	fake.loggregatorEndpointMutex.Lock()
	fake.loggregatorEndpointArgsForCall = append(fake.loggregatorEndpointArgsForCall, struct{}{})
	fake.recordInvocation("LoggregatorEndpoint", []interface{}{})
	fake.loggregatorEndpointMutex.Unlock()
	if fake.LoggregatorEndpointStub != nil {
		return fake.LoggregatorEndpointStub()
	} else {
		return fake.loggregatorEndpointReturns.result1
	}
}

func (fake *FakeCloudControllerClient) LoggregatorEndpointCallCount() int {
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	return len(fake.loggregatorEndpointArgsForCall)
}

func (fake *FakeCloudControllerClient) LoggregatorEndpointReturns(result1 string) {
	fake.LoggregatorEndpointStub = nil
	fake.loggregatorEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) RoutingEndpoint() string {
	fake.routingEndpointMutex.Lock()
	fake.routingEndpointArgsForCall = append(fake.routingEndpointArgsForCall, struct{}{})
	fake.recordInvocation("RoutingEndpoint", []interface{}{})
	fake.routingEndpointMutex.Unlock()
	if fake.RoutingEndpointStub != nil {
		return fake.RoutingEndpointStub()
	} else {
		return fake.routingEndpointReturns.result1
	}
}

func (fake *FakeCloudControllerClient) RoutingEndpointCallCount() int {
	fake.routingEndpointMutex.RLock()
	defer fake.routingEndpointMutex.RUnlock()
	return len(fake.routingEndpointArgsForCall)
}

func (fake *FakeCloudControllerClient) RoutingEndpointReturns(result1 string) {
	fake.RoutingEndpointStub = nil
	fake.routingEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) TokenEndpoint() string {
	fake.tokenEndpointMutex.Lock()
	fake.tokenEndpointArgsForCall = append(fake.tokenEndpointArgsForCall, struct{}{})
	fake.recordInvocation("TokenEndpoint", []interface{}{})
	fake.tokenEndpointMutex.Unlock()
	if fake.TokenEndpointStub != nil {
		return fake.TokenEndpointStub()
	} else {
		return fake.tokenEndpointReturns.result1
	}
}

func (fake *FakeCloudControllerClient) TokenEndpointCallCount() int {
	fake.tokenEndpointMutex.RLock()
	defer fake.tokenEndpointMutex.RUnlock()
	return len(fake.tokenEndpointArgsForCall)
}

func (fake *FakeCloudControllerClient) TokenEndpointReturns(result1 string) {
	fake.TokenEndpointStub = nil
	fake.tokenEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudControllerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.targetCFMutex.RLock()
	defer fake.targetCFMutex.RUnlock()
	fake.aPIMutex.RLock()
	defer fake.aPIMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.authorizationEndpointMutex.RLock()
	defer fake.authorizationEndpointMutex.RUnlock()
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	fake.routingEndpointMutex.RLock()
	defer fake.routingEndpointMutex.RUnlock()
	fake.tokenEndpointMutex.RLock()
	defer fake.tokenEndpointMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCloudControllerClient) recordInvocation(key string, args []interface{}) {
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

var _ configactions.CloudControllerClient = new(FakeCloudControllerClient)
