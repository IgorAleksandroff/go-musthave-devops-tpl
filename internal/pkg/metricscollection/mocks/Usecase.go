// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// SaveCounterMetric provides a mock function with given fields: name, value
func (_m *Usecase) SaveCounterMetric(name string, value int64) {
	_m.Called(name, value)
}

// SaveGaugeMetric provides a mock function with given fields: name, value
func (_m *Usecase) SaveGaugeMetric(name string, value float64) {
	_m.Called(name, value)
}

type NewUsecaseT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t NewUsecaseT) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
