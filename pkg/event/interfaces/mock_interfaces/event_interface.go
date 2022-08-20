// Code generated by MockGen. DO NOT EDIT.
// Source: huangjihui511/event-mgr/pkg/event/interfaces (interfaces: Interface)

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	interfaces "huangjihui511/event-mgr/pkg/watcher/interfaces"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Chan mocks base method.
func (m *MockInterface) Chan() <-chan interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chan")
	ret0, _ := ret[0].(<-chan interface{})
	return ret0
}

// Chan indicates an expected call of Chan.
func (mr *MockInterfaceMockRecorder) Chan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chan", reflect.TypeOf((*MockInterface)(nil).Chan))
}

// Watcher mocks base method.
func (m *MockInterface) Watcher() interfaces.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watcher")
	ret0, _ := ret[0].(interfaces.Interface)
	return ret0
}

// Watcher indicates an expected call of Watcher.
func (mr *MockInterfaceMockRecorder) Watcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watcher", reflect.TypeOf((*MockInterface)(nil).Watcher))
}