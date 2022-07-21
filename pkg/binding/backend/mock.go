// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/binding/backend/interface.go

// Package backend is a generated GoMock package.
package backend

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	asker "github.com/redhat-developer/odo/pkg/binding/asker"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockAddBindingBackend is a mock of AddBindingBackend interface.
type MockAddBindingBackend struct {
	ctrl     *gomock.Controller
	recorder *MockAddBindingBackendMockRecorder
}

// MockAddBindingBackendMockRecorder is the mock recorder for MockAddBindingBackend.
type MockAddBindingBackendMockRecorder struct {
	mock *MockAddBindingBackend
}

// NewMockAddBindingBackend creates a new mock instance.
func NewMockAddBindingBackend(ctrl *gomock.Controller) *MockAddBindingBackend {
	mock := &MockAddBindingBackend{ctrl: ctrl}
	mock.recorder = &MockAddBindingBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddBindingBackend) EXPECT() *MockAddBindingBackendMockRecorder {
	return m.recorder
}

// AskBindAsFiles mocks base method.
func (m *MockAddBindingBackend) AskBindAsFiles(flags map[string]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskBindAsFiles", flags)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskBindAsFiles indicates an expected call of AskBindAsFiles.
func (mr *MockAddBindingBackendMockRecorder) AskBindAsFiles(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskBindAsFiles", reflect.TypeOf((*MockAddBindingBackend)(nil).AskBindAsFiles), flags)
}

// AskBindingName mocks base method.
func (m *MockAddBindingBackend) AskBindingName(defaultName string, flags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskBindingName", defaultName, flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskBindingName indicates an expected call of AskBindingName.
func (mr *MockAddBindingBackendMockRecorder) AskBindingName(defaultName, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskBindingName", reflect.TypeOf((*MockAddBindingBackend)(nil).AskBindingName), defaultName, flags)
}

// AskNamingStrategy mocks base method.
func (m *MockAddBindingBackend) AskNamingStrategy(flags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskNamingStrategy", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskNamingStrategy indicates an expected call of AskNamingStrategy.
func (mr *MockAddBindingBackendMockRecorder) AskNamingStrategy(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskNamingStrategy", reflect.TypeOf((*MockAddBindingBackend)(nil).AskNamingStrategy), flags)
}

// AskOutputFilePath mocks base method.
func (m *MockAddBindingBackend) AskOutputFilePath(flags map[string]string, defaultValue string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskOutputFilePath", flags, defaultValue)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskOutputFilePath indicates an expected call of AskOutputFilePath.
func (mr *MockAddBindingBackendMockRecorder) AskOutputFilePath(flags, defaultValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskOutputFilePath", reflect.TypeOf((*MockAddBindingBackend)(nil).AskOutputFilePath), flags, defaultValue)
}

// SelectCreationOptions mocks base method.
func (m *MockAddBindingBackend) SelectCreationOptions(flags map[string]string) ([]asker.CreationOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectCreationOptions", flags)
	ret0, _ := ret[0].([]asker.CreationOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectCreationOptions indicates an expected call of SelectCreationOptions.
func (mr *MockAddBindingBackendMockRecorder) SelectCreationOptions(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectCreationOptions", reflect.TypeOf((*MockAddBindingBackend)(nil).SelectCreationOptions), flags)
}

// SelectNamespace mocks base method.
func (m *MockAddBindingBackend) SelectNamespace(flags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectNamespace", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectNamespace indicates an expected call of SelectNamespace.
func (mr *MockAddBindingBackendMockRecorder) SelectNamespace(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectNamespace", reflect.TypeOf((*MockAddBindingBackend)(nil).SelectNamespace), flags)
}

// SelectServiceInstance mocks base method.
func (m *MockAddBindingBackend) SelectServiceInstance(serviceName string, serviceMap map[string]unstructured.Unstructured) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectServiceInstance", serviceName, serviceMap)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectServiceInstance indicates an expected call of SelectServiceInstance.
func (mr *MockAddBindingBackendMockRecorder) SelectServiceInstance(serviceName, serviceMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectServiceInstance", reflect.TypeOf((*MockAddBindingBackend)(nil).SelectServiceInstance), serviceName, serviceMap)
}

// SelectWorkloadInstance mocks base method.
func (m *MockAddBindingBackend) SelectWorkloadInstance(workloadName string) (string, schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectWorkloadInstance", workloadName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(schema.GroupVersionKind)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SelectWorkloadInstance indicates an expected call of SelectWorkloadInstance.
func (mr *MockAddBindingBackendMockRecorder) SelectWorkloadInstance(workloadName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectWorkloadInstance", reflect.TypeOf((*MockAddBindingBackend)(nil).SelectWorkloadInstance), workloadName)
}

// Validate mocks base method.
func (m *MockAddBindingBackend) Validate(flags map[string]string, withDevfile bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", flags, withDevfile)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockAddBindingBackendMockRecorder) Validate(flags, withDevfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAddBindingBackend)(nil).Validate), flags, withDevfile)
}
