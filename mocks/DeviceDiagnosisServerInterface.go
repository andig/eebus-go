// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// DeviceDiagnosisServerInterface is an autogenerated mock type for the DeviceDiagnosisServerInterface type
type DeviceDiagnosisServerInterface struct {
	mock.Mock
}

type DeviceDiagnosisServerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceDiagnosisServerInterface) EXPECT() *DeviceDiagnosisServerInterface_Expecter {
	return &DeviceDiagnosisServerInterface_Expecter{mock: &_m.Mock}
}

// SetLocalState provides a mock function with given fields: operatingState
func (_m *DeviceDiagnosisServerInterface) SetLocalState(operatingState *model.DeviceDiagnosisStateDataType) {
	_m.Called(operatingState)
}

// DeviceDiagnosisServerInterface_SetLocalState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLocalState'
type DeviceDiagnosisServerInterface_SetLocalState_Call struct {
	*mock.Call
}

// SetLocalState is a helper method to define mock.On call
//   - operatingState *model.DeviceDiagnosisStateDataType
func (_e *DeviceDiagnosisServerInterface_Expecter) SetLocalState(operatingState interface{}) *DeviceDiagnosisServerInterface_SetLocalState_Call {
	return &DeviceDiagnosisServerInterface_SetLocalState_Call{Call: _e.mock.On("SetLocalState", operatingState)}
}

func (_c *DeviceDiagnosisServerInterface_SetLocalState_Call) Run(run func(operatingState *model.DeviceDiagnosisStateDataType)) *DeviceDiagnosisServerInterface_SetLocalState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DeviceDiagnosisStateDataType))
	})
	return _c
}

func (_c *DeviceDiagnosisServerInterface_SetLocalState_Call) Return() *DeviceDiagnosisServerInterface_SetLocalState_Call {
	_c.Call.Return()
	return _c
}

func (_c *DeviceDiagnosisServerInterface_SetLocalState_Call) RunAndReturn(run func(*model.DeviceDiagnosisStateDataType)) *DeviceDiagnosisServerInterface_SetLocalState_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeviceDiagnosisServerInterface creates a new instance of DeviceDiagnosisServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceDiagnosisServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceDiagnosisServerInterface {
	mock := &DeviceDiagnosisServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
