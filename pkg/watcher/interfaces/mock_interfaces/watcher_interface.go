// Code generated by MockGen. DO NOT EDIT.
// Source: huangjihui511/event-mgr/pkg/watcher/interfaces (interfaces: Interface)

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
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

// Call mocks base method.
func (m *MockInterface) Call(arg0 context.Context) interfaces.ResultInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0)
	ret0, _ := ret[0].(interfaces.ResultInterface)
	return ret0
}

// Call indicates an expected call of Call.
func (mr *MockInterfaceMockRecorder) Call(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockInterface)(nil).Call), arg0)
}

// Name mocks base method.
func (m *MockInterface) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockInterfaceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInterface)(nil).Name))
}