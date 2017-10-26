// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/trussle/courier/pkg/audit (interfaces: Log)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/trussle/courier/pkg/models"
	reflect "reflect"
)

// MockLog is a mock of Log interface
type MockLog struct {
	ctrl     *gomock.Controller
	recorder *MockLogMockRecorder
}

// MockLogMockRecorder is the mock recorder for MockLog
type MockLogMockRecorder struct {
	mock *MockLog
}

// NewMockLog creates a new mock instance
func NewMockLog(ctrl *gomock.Controller) *MockLog {
	mock := &MockLog{ctrl: ctrl}
	mock.recorder = &MockLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLog) EXPECT() *MockLogMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockLog) Append(arg0 models.Transaction) error {
	ret := m.ctrl.Call(m, "Append", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Append indicates an expected call of Append
func (mr *MockLogMockRecorder) Append(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockLog)(nil).Append), arg0)
}