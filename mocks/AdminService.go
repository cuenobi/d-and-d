// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	model "d-and-d/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// AdminService is an autogenerated mock type for the AdminService type
type AdminService struct {
	mock.Mock
}

type AdminService_Expecter struct {
	mock *mock.Mock
}

func (_m *AdminService) EXPECT() *AdminService_Expecter {
	return &AdminService_Expecter{mock: &_m.Mock}
}

// CreateClass provides a mock function with given fields: class
func (_m *AdminService) CreateClass(class *model.Class) error {
	ret := _m.Called(class)

	if len(ret) == 0 {
		panic("no return value specified for CreateClass")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Class) error); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_CreateClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateClass'
type AdminService_CreateClass_Call struct {
	*mock.Call
}

// CreateClass is a helper method to define mock.On call
//   - class *model.Class
func (_e *AdminService_Expecter) CreateClass(class interface{}) *AdminService_CreateClass_Call {
	return &AdminService_CreateClass_Call{Call: _e.mock.On("CreateClass", class)}
}

func (_c *AdminService_CreateClass_Call) Run(run func(class *model.Class)) *AdminService_CreateClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Class))
	})
	return _c
}

func (_c *AdminService_CreateClass_Call) Return(_a0 error) *AdminService_CreateClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_CreateClass_Call) RunAndReturn(run func(*model.Class) error) *AdminService_CreateClass_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDifficultyLevel provides a mock function with given fields: diff
func (_m *AdminService) CreateDifficultyLevel(diff *model.DifficultyLevels) error {
	ret := _m.Called(diff)

	if len(ret) == 0 {
		panic("no return value specified for CreateDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.DifficultyLevels) error); ok {
		r0 = rf(diff)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_CreateDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDifficultyLevel'
type AdminService_CreateDifficultyLevel_Call struct {
	*mock.Call
}

// CreateDifficultyLevel is a helper method to define mock.On call
//   - diff *model.DifficultyLevels
func (_e *AdminService_Expecter) CreateDifficultyLevel(diff interface{}) *AdminService_CreateDifficultyLevel_Call {
	return &AdminService_CreateDifficultyLevel_Call{Call: _e.mock.On("CreateDifficultyLevel", diff)}
}

func (_c *AdminService_CreateDifficultyLevel_Call) Run(run func(diff *model.DifficultyLevels)) *AdminService_CreateDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DifficultyLevels))
	})
	return _c
}

func (_c *AdminService_CreateDifficultyLevel_Call) Return(_a0 error) *AdminService_CreateDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_CreateDifficultyLevel_Call) RunAndReturn(run func(*model.DifficultyLevels) error) *AdminService_CreateDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRace provides a mock function with given fields: race
func (_m *AdminService) CreateRace(race *model.Race) error {
	ret := _m.Called(race)

	if len(ret) == 0 {
		panic("no return value specified for CreateRace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Race) error); ok {
		r0 = rf(race)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_CreateRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRace'
type AdminService_CreateRace_Call struct {
	*mock.Call
}

// CreateRace is a helper method to define mock.On call
//   - race *model.Race
func (_e *AdminService_Expecter) CreateRace(race interface{}) *AdminService_CreateRace_Call {
	return &AdminService_CreateRace_Call{Call: _e.mock.On("CreateRace", race)}
}

func (_c *AdminService_CreateRace_Call) Run(run func(race *model.Race)) *AdminService_CreateRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Race))
	})
	return _c
}

func (_c *AdminService_CreateRace_Call) Return(_a0 error) *AdminService_CreateRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_CreateRace_Call) RunAndReturn(run func(*model.Race) error) *AdminService_CreateRace_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteClass provides a mock function with given fields: id
func (_m *AdminService) DeleteClass(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteClass")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_DeleteClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteClass'
type AdminService_DeleteClass_Call struct {
	*mock.Call
}

// DeleteClass is a helper method to define mock.On call
//   - id string
func (_e *AdminService_Expecter) DeleteClass(id interface{}) *AdminService_DeleteClass_Call {
	return &AdminService_DeleteClass_Call{Call: _e.mock.On("DeleteClass", id)}
}

func (_c *AdminService_DeleteClass_Call) Run(run func(id string)) *AdminService_DeleteClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AdminService_DeleteClass_Call) Return(_a0 error) *AdminService_DeleteClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_DeleteClass_Call) RunAndReturn(run func(string) error) *AdminService_DeleteClass_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDifficultyLevel provides a mock function with given fields: id
func (_m *AdminService) DeleteDifficultyLevel(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_DeleteDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDifficultyLevel'
type AdminService_DeleteDifficultyLevel_Call struct {
	*mock.Call
}

// DeleteDifficultyLevel is a helper method to define mock.On call
//   - id string
func (_e *AdminService_Expecter) DeleteDifficultyLevel(id interface{}) *AdminService_DeleteDifficultyLevel_Call {
	return &AdminService_DeleteDifficultyLevel_Call{Call: _e.mock.On("DeleteDifficultyLevel", id)}
}

func (_c *AdminService_DeleteDifficultyLevel_Call) Run(run func(id string)) *AdminService_DeleteDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AdminService_DeleteDifficultyLevel_Call) Return(_a0 error) *AdminService_DeleteDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_DeleteDifficultyLevel_Call) RunAndReturn(run func(string) error) *AdminService_DeleteDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRace provides a mock function with given fields: id
func (_m *AdminService) DeleteRace(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_DeleteRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRace'
type AdminService_DeleteRace_Call struct {
	*mock.Call
}

// DeleteRace is a helper method to define mock.On call
//   - id string
func (_e *AdminService_Expecter) DeleteRace(id interface{}) *AdminService_DeleteRace_Call {
	return &AdminService_DeleteRace_Call{Call: _e.mock.On("DeleteRace", id)}
}

func (_c *AdminService_DeleteRace_Call) Run(run func(id string)) *AdminService_DeleteRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AdminService_DeleteRace_Call) Return(_a0 error) *AdminService_DeleteRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_DeleteRace_Call) RunAndReturn(run func(string) error) *AdminService_DeleteRace_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateClass provides a mock function with given fields: class
func (_m *AdminService) UpdateClass(class *model.Class) error {
	ret := _m.Called(class)

	if len(ret) == 0 {
		panic("no return value specified for UpdateClass")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Class) error); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_UpdateClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateClass'
type AdminService_UpdateClass_Call struct {
	*mock.Call
}

// UpdateClass is a helper method to define mock.On call
//   - class *model.Class
func (_e *AdminService_Expecter) UpdateClass(class interface{}) *AdminService_UpdateClass_Call {
	return &AdminService_UpdateClass_Call{Call: _e.mock.On("UpdateClass", class)}
}

func (_c *AdminService_UpdateClass_Call) Run(run func(class *model.Class)) *AdminService_UpdateClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Class))
	})
	return _c
}

func (_c *AdminService_UpdateClass_Call) Return(_a0 error) *AdminService_UpdateClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_UpdateClass_Call) RunAndReturn(run func(*model.Class) error) *AdminService_UpdateClass_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDifficultyLevel provides a mock function with given fields: diff
func (_m *AdminService) UpdateDifficultyLevel(diff *model.DifficultyLevels) error {
	ret := _m.Called(diff)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.DifficultyLevels) error); ok {
		r0 = rf(diff)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_UpdateDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDifficultyLevel'
type AdminService_UpdateDifficultyLevel_Call struct {
	*mock.Call
}

// UpdateDifficultyLevel is a helper method to define mock.On call
//   - diff *model.DifficultyLevels
func (_e *AdminService_Expecter) UpdateDifficultyLevel(diff interface{}) *AdminService_UpdateDifficultyLevel_Call {
	return &AdminService_UpdateDifficultyLevel_Call{Call: _e.mock.On("UpdateDifficultyLevel", diff)}
}

func (_c *AdminService_UpdateDifficultyLevel_Call) Run(run func(diff *model.DifficultyLevels)) *AdminService_UpdateDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DifficultyLevels))
	})
	return _c
}

func (_c *AdminService_UpdateDifficultyLevel_Call) Return(_a0 error) *AdminService_UpdateDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_UpdateDifficultyLevel_Call) RunAndReturn(run func(*model.DifficultyLevels) error) *AdminService_UpdateDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRace provides a mock function with given fields: race
func (_m *AdminService) UpdateRace(race *model.Race) error {
	ret := _m.Called(race)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Race) error); ok {
		r0 = rf(race)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdminService_UpdateRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRace'
type AdminService_UpdateRace_Call struct {
	*mock.Call
}

// UpdateRace is a helper method to define mock.On call
//   - race *model.Race
func (_e *AdminService_Expecter) UpdateRace(race interface{}) *AdminService_UpdateRace_Call {
	return &AdminService_UpdateRace_Call{Call: _e.mock.On("UpdateRace", race)}
}

func (_c *AdminService_UpdateRace_Call) Run(run func(race *model.Race)) *AdminService_UpdateRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Race))
	})
	return _c
}

func (_c *AdminService_UpdateRace_Call) Return(_a0 error) *AdminService_UpdateRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AdminService_UpdateRace_Call) RunAndReturn(run func(*model.Race) error) *AdminService_UpdateRace_Call {
	_c.Call.Return(run)
	return _c
}

// NewAdminService creates a new instance of AdminService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminService {
	mock := &AdminService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
