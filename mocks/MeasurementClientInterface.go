// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// MeasurementClientInterface is an autogenerated mock type for the MeasurementClientInterface type
type MeasurementClientInterface struct {
	mock.Mock
}

type MeasurementClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MeasurementClientInterface) EXPECT() *MeasurementClientInterface_Expecter {
	return &MeasurementClientInterface_Expecter{mock: &_m.Mock}
}

// RequestConstraints provides a mock function with given fields:
func (_m *MeasurementClientInterface) RequestConstraints() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestConstraints")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementClientInterface_RequestConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestConstraints'
type MeasurementClientInterface_RequestConstraints_Call struct {
	*mock.Call
}

// RequestConstraints is a helper method to define mock.On call
func (_e *MeasurementClientInterface_Expecter) RequestConstraints() *MeasurementClientInterface_RequestConstraints_Call {
	return &MeasurementClientInterface_RequestConstraints_Call{Call: _e.mock.On("RequestConstraints")}
}

func (_c *MeasurementClientInterface_RequestConstraints_Call) Run(run func()) *MeasurementClientInterface_RequestConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MeasurementClientInterface_RequestConstraints_Call) Return(_a0 *model.MsgCounterType, _a1 error) *MeasurementClientInterface_RequestConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementClientInterface_RequestConstraints_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *MeasurementClientInterface_RequestConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// RequestData provides a mock function with given fields:
func (_m *MeasurementClientInterface) RequestData() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestData")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementClientInterface_RequestData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestData'
type MeasurementClientInterface_RequestData_Call struct {
	*mock.Call
}

// RequestData is a helper method to define mock.On call
func (_e *MeasurementClientInterface_Expecter) RequestData() *MeasurementClientInterface_RequestData_Call {
	return &MeasurementClientInterface_RequestData_Call{Call: _e.mock.On("RequestData")}
}

func (_c *MeasurementClientInterface_RequestData_Call) Run(run func()) *MeasurementClientInterface_RequestData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MeasurementClientInterface_RequestData_Call) Return(_a0 *model.MsgCounterType, _a1 error) *MeasurementClientInterface_RequestData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementClientInterface_RequestData_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *MeasurementClientInterface_RequestData_Call {
	_c.Call.Return(run)
	return _c
}

// RequestDescriptions provides a mock function with given fields:
func (_m *MeasurementClientInterface) RequestDescriptions() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestDescriptions")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementClientInterface_RequestDescriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestDescriptions'
type MeasurementClientInterface_RequestDescriptions_Call struct {
	*mock.Call
}

// RequestDescriptions is a helper method to define mock.On call
func (_e *MeasurementClientInterface_Expecter) RequestDescriptions() *MeasurementClientInterface_RequestDescriptions_Call {
	return &MeasurementClientInterface_RequestDescriptions_Call{Call: _e.mock.On("RequestDescriptions")}
}

func (_c *MeasurementClientInterface_RequestDescriptions_Call) Run(run func()) *MeasurementClientInterface_RequestDescriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MeasurementClientInterface_RequestDescriptions_Call) Return(_a0 *model.MsgCounterType, _a1 error) *MeasurementClientInterface_RequestDescriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementClientInterface_RequestDescriptions_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *MeasurementClientInterface_RequestDescriptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMeasurementClientInterface creates a new instance of MeasurementClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMeasurementClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MeasurementClientInterface {
	mock := &MeasurementClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
