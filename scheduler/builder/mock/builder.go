// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler/builder/builder.go

// Package mock_builder is a generated GoMock package.
package mock_builder

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/aliykh/reter/scheduler/models"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunner) Run(ctx context.Context, task models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), ctx, task)
}
