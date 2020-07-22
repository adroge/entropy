// Code generated by MockGen. DO NOT EDIT.
// Source: entropy.go

// Package mock_entropy is a generated GoMock package.
package mock_entropy

import (
	entropy "github.com/adroge/entropy"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMethods is a mock of Methods interface.
type MockMethods struct {
	ctrl     *gomock.Controller
	recorder *MockMethodsMockRecorder
}

// MockMethodsMockRecorder is the mock recorder for MockMethods.
type MockMethodsMockRecorder struct {
	mock *MockMethods
}

// NewMockMethods creates a new mock instance.
func NewMockMethods(ctrl *gomock.Controller) *MockMethods {
	mock := &MockMethods{ctrl: ctrl}
	mock.recorder = &MockMethodsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMethods) EXPECT() *MockMethodsMockRecorder {
	return m.recorder
}

// Alphabets mocks base method.
func (m *MockMethods) Alphabets(newAlphabets []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alphabets", newAlphabets)
	ret0, _ := ret[0].(error)
	return ret0
}

// Alphabets indicates an expected call of Alphabets.
func (mr *MockMethodsMockRecorder) Alphabets(newAlphabets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alphabets", reflect.TypeOf((*MockMethods)(nil).Alphabets), newAlphabets)
}

// Bounds mocks base method.
func (m *MockMethods) Bounds(veryWeak, weak, reasonable, strong float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bounds", veryWeak, weak, reasonable, strong)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bounds indicates an expected call of Bounds.
func (mr *MockMethodsMockRecorder) Bounds(veryWeak, weak, reasonable, strong interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bounds", reflect.TypeOf((*MockMethods)(nil).Bounds), veryWeak, weak, reasonable, strong)
}

// Calculate mocks base method.
func (m *MockMethods) Calculate(input string) (entropy.EntropyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Calculate", input)
	ret0, _ := ret[0].(entropy.EntropyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Calculate indicates an expected call of Calculate.
func (mr *MockMethodsMockRecorder) Calculate(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Calculate", reflect.TypeOf((*MockMethods)(nil).Calculate), input)
}

// Descriptions mocks base method.
func (m *MockMethods) Descriptions(tags entropy.DescriptionTags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptions", tags)
	ret0, _ := ret[0].(error)
	return ret0
}

// Descriptions indicates an expected call of Descriptions.
func (mr *MockMethodsMockRecorder) Descriptions(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptions", reflect.TypeOf((*MockMethods)(nil).Descriptions), tags)
}

// EntropyBounds mocks base method.
func (m *MockMethods) EntropyBounds() []float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntropyBounds")
	ret0, _ := ret[0].([]float64)
	return ret0
}

// EntropyBounds indicates an expected call of EntropyBounds.
func (mr *MockMethodsMockRecorder) EntropyBounds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntropyBounds", reflect.TypeOf((*MockMethods)(nil).EntropyBounds))
}