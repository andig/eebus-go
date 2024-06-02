// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	mock "github.com/stretchr/testify/mock"

	spine_goapi "github.com/enbility/spine-go/api"
)

// CemEVCEMInterface is an autogenerated mock type for the CemEVCEMInterface type
type CemEVCEMInterface struct {
	mock.Mock
}

type CemEVCEMInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemEVCEMInterface) EXPECT() *CemEVCEMInterface_Expecter {
	return &CemEVCEMInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemEVCEMInterface) AddFeatures() {
	_m.Called()
}

// CemEVCEMInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemEVCEMInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemEVCEMInterface_Expecter) AddFeatures() *CemEVCEMInterface_AddFeatures_Call {
	return &CemEVCEMInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemEVCEMInterface_AddFeatures_Call) Run(run func()) *CemEVCEMInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCEMInterface_AddFeatures_Call) Return() *CemEVCEMInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCEMInterface_AddFeatures_Call) RunAndReturn(run func()) *CemEVCEMInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemEVCEMInterface) AddUseCase() {
	_m.Called()
}

// CemEVCEMInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemEVCEMInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemEVCEMInterface_Expecter) AddUseCase() *CemEVCEMInterface_AddUseCase_Call {
	return &CemEVCEMInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemEVCEMInterface_AddUseCase_Call) Run(run func()) *CemEVCEMInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCEMInterface_AddUseCase_Call) Return() *CemEVCEMInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCEMInterface_AddUseCase_Call) RunAndReturn(run func()) *CemEVCEMInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for AvailableScenariosForEntity")
	}

	var r0 []uint
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []uint); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	return r0
}

// CemEVCEMInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CemEVCEMInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CemEVCEMInterface_AvailableScenariosForEntity_Call {
	return &CemEVCEMInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CemEVCEMInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CemEVCEMInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCEMInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CemEVCEMInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// CurrentPerPhase provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) CurrentPerPhase(entity spine_goapi.EntityRemoteInterface) ([]float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for CurrentPerPhase")
	}

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) ([]float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []float64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCEMInterface_CurrentPerPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentPerPhase'
type CemEVCEMInterface_CurrentPerPhase_Call struct {
	*mock.Call
}

// CurrentPerPhase is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) CurrentPerPhase(entity interface{}) *CemEVCEMInterface_CurrentPerPhase_Call {
	return &CemEVCEMInterface_CurrentPerPhase_Call{Call: _e.mock.On("CurrentPerPhase", entity)}
}

func (_c *CemEVCEMInterface_CurrentPerPhase_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_CurrentPerPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_CurrentPerPhase_Call) Return(_a0 []float64, _a1 error) *CemEVCEMInterface_CurrentPerPhase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCEMInterface_CurrentPerPhase_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) ([]float64, error)) *CemEVCEMInterface_CurrentPerPhase_Call {
	_c.Call.Return(run)
	return _c
}

// EnergyCharged provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) EnergyCharged(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EnergyCharged")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCEMInterface_EnergyCharged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnergyCharged'
type CemEVCEMInterface_EnergyCharged_Call struct {
	*mock.Call
}

// EnergyCharged is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) EnergyCharged(entity interface{}) *CemEVCEMInterface_EnergyCharged_Call {
	return &CemEVCEMInterface_EnergyCharged_Call{Call: _e.mock.On("EnergyCharged", entity)}
}

func (_c *CemEVCEMInterface_EnergyCharged_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_EnergyCharged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_EnergyCharged_Call) Return(_a0 float64, _a1 error) *CemEVCEMInterface_EnergyCharged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCEMInterface_EnergyCharged_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemEVCEMInterface_EnergyCharged_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntityType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVCEMInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CemEVCEMInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CemEVCEMInterface_IsCompatibleEntityType_Call {
	return &CemEVCEMInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CemEVCEMInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CemEVCEMInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCEMInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemEVCEMInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CemEVCEMInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
	ret := _m.Called(entity, scenario)

	if len(ret) == 0 {
		panic("no return value specified for IsScenarioAvailableAtEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, uint) bool); ok {
		r0 = rf(entity, scenario)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemEVCEMInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CemEVCEMInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CemEVCEMInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call {
	return &CemEVCEMInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CemEVCEMInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// PhasesConnected provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) PhasesConnected(entity spine_goapi.EntityRemoteInterface) (uint, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PhasesConnected")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (uint, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) uint); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCEMInterface_PhasesConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PhasesConnected'
type CemEVCEMInterface_PhasesConnected_Call struct {
	*mock.Call
}

// PhasesConnected is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) PhasesConnected(entity interface{}) *CemEVCEMInterface_PhasesConnected_Call {
	return &CemEVCEMInterface_PhasesConnected_Call{Call: _e.mock.On("PhasesConnected", entity)}
}

func (_c *CemEVCEMInterface_PhasesConnected_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_PhasesConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_PhasesConnected_Call) Return(_a0 uint, _a1 error) *CemEVCEMInterface_PhasesConnected_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCEMInterface_PhasesConnected_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (uint, error)) *CemEVCEMInterface_PhasesConnected_Call {
	_c.Call.Return(run)
	return _c
}

// PowerPerPhase provides a mock function with given fields: entity
func (_m *CemEVCEMInterface) PowerPerPhase(entity spine_goapi.EntityRemoteInterface) ([]float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PowerPerPhase")
	}

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) ([]float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []float64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemEVCEMInterface_PowerPerPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PowerPerPhase'
type CemEVCEMInterface_PowerPerPhase_Call struct {
	*mock.Call
}

// PowerPerPhase is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemEVCEMInterface_Expecter) PowerPerPhase(entity interface{}) *CemEVCEMInterface_PowerPerPhase_Call {
	return &CemEVCEMInterface_PowerPerPhase_Call{Call: _e.mock.On("PowerPerPhase", entity)}
}

func (_c *CemEVCEMInterface_PowerPerPhase_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemEVCEMInterface_PowerPerPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemEVCEMInterface_PowerPerPhase_Call) Return(_a0 []float64, _a1 error) *CemEVCEMInterface_PowerPerPhase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemEVCEMInterface_PowerPerPhase_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) ([]float64, error)) *CemEVCEMInterface_PowerPerPhase_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CemEVCEMInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteEntitiesScenarios")
	}

	var r0 []eebus_goapi.RemoteEntityScenarios
	if rf, ok := ret.Get(0).(func() []eebus_goapi.RemoteEntityScenarios); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]eebus_goapi.RemoteEntityScenarios)
		}
	}

	return r0
}

// CemEVCEMInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CemEVCEMInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CemEVCEMInterface_Expecter) RemoteEntitiesScenarios() *CemEVCEMInterface_RemoteEntitiesScenarios_Call {
	return &CemEVCEMInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CemEVCEMInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CemEVCEMInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCEMInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CemEVCEMInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemEVCEMInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CemEVCEMInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CemEVCEMInterface) RemoveUseCase() {
	_m.Called()
}

// CemEVCEMInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CemEVCEMInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CemEVCEMInterface_Expecter) RemoveUseCase() *CemEVCEMInterface_RemoveUseCase_Call {
	return &CemEVCEMInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CemEVCEMInterface_RemoveUseCase_Call) Run(run func()) *CemEVCEMInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemEVCEMInterface_RemoveUseCase_Call) Return() *CemEVCEMInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCEMInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CemEVCEMInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemEVCEMInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemEVCEMInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemEVCEMInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemEVCEMInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemEVCEMInterface_UpdateUseCaseAvailability_Call {
	return &CemEVCEMInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemEVCEMInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemEVCEMInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemEVCEMInterface_UpdateUseCaseAvailability_Call) Return() *CemEVCEMInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemEVCEMInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemEVCEMInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemEVCEMInterface creates a new instance of CemEVCEMInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemEVCEMInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemEVCEMInterface {
	mock := &CemEVCEMInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
