// Code generated by MockGen. DO NOT EDIT.
// Source: internal/calculator/calculator.go

// Package calculator is a generated GoMock package.
package calculator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCalculator is a mock of Calculator interface.
type MockCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorMockRecorder
}

// MockCalculatorMockRecorder is the mock recorder for MockCalculator.
type MockCalculatorMockRecorder struct {
	mock *MockCalculator
}

// NewMockCalculator creates a new mock instance.
func NewMockCalculator(ctrl *gomock.Controller) *MockCalculator {
	mock := &MockCalculator{ctrl: ctrl}
	mock.recorder = &MockCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculator) EXPECT() *MockCalculatorMockRecorder {
	return m.recorder
}

// Calc mocks base method.
func (m *MockCalculator) Calc(trxs []Transaction) []TransactionReturn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Calc", trxs)
	ret0, _ := ret[0].([]TransactionReturn)
	return ret0
}

// Calc indicates an expected call of Calc.
func (mr *MockCalculatorMockRecorder) Calc(trxs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Calc", reflect.TypeOf((*MockCalculator)(nil).Calc), trxs)
}
