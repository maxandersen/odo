// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/exec/interface.go

// Package exec is a generated GoMock package.
package exec

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ExecuteCommand mocks base method.
func (m *MockClient) ExecuteCommand(command []string, podName, containerName string, show bool, stdoutWriter, stderrWriter *io.PipeWriter) ([]string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteCommand", command, podName, containerName, show, stdoutWriter, stderrWriter)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecuteCommand indicates an expected call of ExecuteCommand.
func (mr *MockClientMockRecorder) ExecuteCommand(command, podName, containerName, show, stdoutWriter, stderrWriter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCommand", reflect.TypeOf((*MockClient)(nil).ExecuteCommand), command, podName, containerName, show, stdoutWriter, stderrWriter)
}
