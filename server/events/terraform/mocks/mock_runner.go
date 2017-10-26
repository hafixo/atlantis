// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server/events/terraform (interfaces: Runner)

package mocks

import (
	go_version "github.com/hashicorp/go-version"
	logging "github.com/hootsuite/atlantis/server/logging"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockRunner() *MockRunner {
	return &MockRunner{fail: pegomock.GlobalFailHandler}
}

func (mock *MockRunner) Version() *go_version.Version {
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Version", params, []reflect.Type{reflect.TypeOf((**go_version.Version)(nil)).Elem()})
	var ret0 *go_version.Version
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*go_version.Version)
		}
	}
	return ret0
}

func (mock *MockRunner) RunCommandWithVersion(log *logging.SimpleLogger, path string, args []string, v *go_version.Version, env string) (string, error) {
	params := []pegomock.Param{log, path, args, v, env}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunCommandWithVersion", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockRunner) RunInitAndEnv(log *logging.SimpleLogger, path string, env string, extraInitArgs []string, version *go_version.Version) ([]string, error) {
	params := []pegomock.Param{log, path, env, extraInitArgs, version}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunInitAndEnv", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockRunner) VerifyWasCalledOnce() *VerifierRunner {
	return &VerifierRunner{mock, pegomock.Times(1), nil}
}

func (mock *MockRunner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierRunner {
	return &VerifierRunner{mock, invocationCountMatcher, nil}
}

func (mock *MockRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierRunner {
	return &VerifierRunner{mock, invocationCountMatcher, inOrderContext}
}

type VerifierRunner struct {
	mock                   *MockRunner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierRunner) Version() *Runner_Version_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", params)
	return &Runner_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Runner_Version_OngoingVerification struct {
	mock              *MockRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *Runner_Version_OngoingVerification) GetCapturedArguments() {
}

func (c *Runner_Version_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierRunner) RunCommandWithVersion(log *logging.SimpleLogger, path string, args []string, v *go_version.Version, env string) *Runner_RunCommandWithVersion_OngoingVerification {
	params := []pegomock.Param{log, path, args, v, env}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunCommandWithVersion", params)
	return &Runner_RunCommandWithVersion_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Runner_RunCommandWithVersion_OngoingVerification struct {
	mock              *MockRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *Runner_RunCommandWithVersion_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, string, []string, *go_version.Version, string) {
	log, path, args, v, env := c.GetAllCapturedArguments()
	return log[len(log)-1], path[len(path)-1], args[len(args)-1], v[len(v)-1], env[len(env)-1]
}

func (c *Runner_RunCommandWithVersion_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 []string, _param2 [][]string, _param3 []*go_version.Version, _param4 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.([]string)
		}
		_param3 = make([]*go_version.Version, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(*go_version.Version)
		}
		_param4 = make([]string, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierRunner) RunInitAndEnv(log *logging.SimpleLogger, path string, env string, extraInitArgs []string, version *go_version.Version) *Runner_RunInitAndEnv_OngoingVerification {
	params := []pegomock.Param{log, path, env, extraInitArgs, version}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunInitAndEnv", params)
	return &Runner_RunInitAndEnv_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Runner_RunInitAndEnv_OngoingVerification struct {
	mock              *MockRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *Runner_RunInitAndEnv_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, string, string, []string, *go_version.Version) {
	log, path, env, extraInitArgs, version := c.GetAllCapturedArguments()
	return log[len(log)-1], path[len(path)-1], env[len(env)-1], extraInitArgs[len(extraInitArgs)-1], version[len(version)-1]
}

func (c *Runner_RunInitAndEnv_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 []string, _param2 []string, _param3 [][]string, _param4 []*go_version.Version) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([][]string, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.([]string)
		}
		_param4 = make([]*go_version.Version, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(*go_version.Version)
		}
	}
	return
}
