// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	model "d-and-d/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// OptionsRepository is an autogenerated mock type for the OptionsRepository type
type OptionsRepository struct {
	mock.Mock
}

type OptionsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *OptionsRepository) EXPECT() *OptionsRepository_Expecter {
	return &OptionsRepository_Expecter{mock: &_m.Mock}
}

// CreateClass provides a mock function with given fields: class
func (_m *OptionsRepository) CreateClass(class *model.Class) error {
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

// OptionsRepository_CreateClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateClass'
type OptionsRepository_CreateClass_Call struct {
	*mock.Call
}

// CreateClass is a helper method to define mock.On call
//   - class *model.Class
func (_e *OptionsRepository_Expecter) CreateClass(class interface{}) *OptionsRepository_CreateClass_Call {
	return &OptionsRepository_CreateClass_Call{Call: _e.mock.On("CreateClass", class)}
}

func (_c *OptionsRepository_CreateClass_Call) Run(run func(class *model.Class)) *OptionsRepository_CreateClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Class))
	})
	return _c
}

func (_c *OptionsRepository_CreateClass_Call) Return(_a0 error) *OptionsRepository_CreateClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_CreateClass_Call) RunAndReturn(run func(*model.Class) error) *OptionsRepository_CreateClass_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDifficultyLevel provides a mock function with given fields: diffLevel
func (_m *OptionsRepository) CreateDifficultyLevel(diffLevel *model.DifficultyLevels) error {
	ret := _m.Called(diffLevel)

	if len(ret) == 0 {
		panic("no return value specified for CreateDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.DifficultyLevels) error); ok {
		r0 = rf(diffLevel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OptionsRepository_CreateDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDifficultyLevel'
type OptionsRepository_CreateDifficultyLevel_Call struct {
	*mock.Call
}

// CreateDifficultyLevel is a helper method to define mock.On call
//   - diffLevel *model.DifficultyLevels
func (_e *OptionsRepository_Expecter) CreateDifficultyLevel(diffLevel interface{}) *OptionsRepository_CreateDifficultyLevel_Call {
	return &OptionsRepository_CreateDifficultyLevel_Call{Call: _e.mock.On("CreateDifficultyLevel", diffLevel)}
}

func (_c *OptionsRepository_CreateDifficultyLevel_Call) Run(run func(diffLevel *model.DifficultyLevels)) *OptionsRepository_CreateDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DifficultyLevels))
	})
	return _c
}

func (_c *OptionsRepository_CreateDifficultyLevel_Call) Return(_a0 error) *OptionsRepository_CreateDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_CreateDifficultyLevel_Call) RunAndReturn(run func(*model.DifficultyLevels) error) *OptionsRepository_CreateDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRace provides a mock function with given fields: race
func (_m *OptionsRepository) CreateRace(race *model.Race) error {
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

// OptionsRepository_CreateRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRace'
type OptionsRepository_CreateRace_Call struct {
	*mock.Call
}

// CreateRace is a helper method to define mock.On call
//   - race *model.Race
func (_e *OptionsRepository_Expecter) CreateRace(race interface{}) *OptionsRepository_CreateRace_Call {
	return &OptionsRepository_CreateRace_Call{Call: _e.mock.On("CreateRace", race)}
}

func (_c *OptionsRepository_CreateRace_Call) Run(run func(race *model.Race)) *OptionsRepository_CreateRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Race))
	})
	return _c
}

func (_c *OptionsRepository_CreateRace_Call) Return(_a0 error) *OptionsRepository_CreateRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_CreateRace_Call) RunAndReturn(run func(*model.Race) error) *OptionsRepository_CreateRace_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteClass provides a mock function with given fields: classID
func (_m *OptionsRepository) DeleteClass(classID string) error {
	ret := _m.Called(classID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteClass")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(classID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OptionsRepository_DeleteClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteClass'
type OptionsRepository_DeleteClass_Call struct {
	*mock.Call
}

// DeleteClass is a helper method to define mock.On call
//   - classID string
func (_e *OptionsRepository_Expecter) DeleteClass(classID interface{}) *OptionsRepository_DeleteClass_Call {
	return &OptionsRepository_DeleteClass_Call{Call: _e.mock.On("DeleteClass", classID)}
}

func (_c *OptionsRepository_DeleteClass_Call) Run(run func(classID string)) *OptionsRepository_DeleteClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_DeleteClass_Call) Return(_a0 error) *OptionsRepository_DeleteClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_DeleteClass_Call) RunAndReturn(run func(string) error) *OptionsRepository_DeleteClass_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDifficultyLevel provides a mock function with given fields: diffLvID
func (_m *OptionsRepository) DeleteDifficultyLevel(diffLvID string) error {
	ret := _m.Called(diffLvID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(diffLvID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OptionsRepository_DeleteDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDifficultyLevel'
type OptionsRepository_DeleteDifficultyLevel_Call struct {
	*mock.Call
}

// DeleteDifficultyLevel is a helper method to define mock.On call
//   - diffLvID string
func (_e *OptionsRepository_Expecter) DeleteDifficultyLevel(diffLvID interface{}) *OptionsRepository_DeleteDifficultyLevel_Call {
	return &OptionsRepository_DeleteDifficultyLevel_Call{Call: _e.mock.On("DeleteDifficultyLevel", diffLvID)}
}

func (_c *OptionsRepository_DeleteDifficultyLevel_Call) Run(run func(diffLvID string)) *OptionsRepository_DeleteDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_DeleteDifficultyLevel_Call) Return(_a0 error) *OptionsRepository_DeleteDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_DeleteDifficultyLevel_Call) RunAndReturn(run func(string) error) *OptionsRepository_DeleteDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRace provides a mock function with given fields: raceID
func (_m *OptionsRepository) DeleteRace(raceID string) error {
	ret := _m.Called(raceID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(raceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OptionsRepository_DeleteRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRace'
type OptionsRepository_DeleteRace_Call struct {
	*mock.Call
}

// DeleteRace is a helper method to define mock.On call
//   - raceID string
func (_e *OptionsRepository_Expecter) DeleteRace(raceID interface{}) *OptionsRepository_DeleteRace_Call {
	return &OptionsRepository_DeleteRace_Call{Call: _e.mock.On("DeleteRace", raceID)}
}

func (_c *OptionsRepository_DeleteRace_Call) Run(run func(raceID string)) *OptionsRepository_DeleteRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_DeleteRace_Call) Return(_a0 error) *OptionsRepository_DeleteRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_DeleteRace_Call) RunAndReturn(run func(string) error) *OptionsRepository_DeleteRace_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllClass provides a mock function with no fields
func (_m *OptionsRepository) GetAllClass() ([]*model.Class, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllClass")
	}

	var r0 []*model.Class
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Class, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Class); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Class)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetAllClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllClass'
type OptionsRepository_GetAllClass_Call struct {
	*mock.Call
}

// GetAllClass is a helper method to define mock.On call
func (_e *OptionsRepository_Expecter) GetAllClass() *OptionsRepository_GetAllClass_Call {
	return &OptionsRepository_GetAllClass_Call{Call: _e.mock.On("GetAllClass")}
}

func (_c *OptionsRepository_GetAllClass_Call) Run(run func()) *OptionsRepository_GetAllClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OptionsRepository_GetAllClass_Call) Return(_a0 []*model.Class, _a1 error) *OptionsRepository_GetAllClass_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetAllClass_Call) RunAndReturn(run func() ([]*model.Class, error)) *OptionsRepository_GetAllClass_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllDifficultyLevel provides a mock function with no fields
func (_m *OptionsRepository) GetAllDifficultyLevel() ([]*model.DifficultyLevels, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllDifficultyLevel")
	}

	var r0 []*model.DifficultyLevels
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.DifficultyLevels, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.DifficultyLevels); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DifficultyLevels)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetAllDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllDifficultyLevel'
type OptionsRepository_GetAllDifficultyLevel_Call struct {
	*mock.Call
}

// GetAllDifficultyLevel is a helper method to define mock.On call
func (_e *OptionsRepository_Expecter) GetAllDifficultyLevel() *OptionsRepository_GetAllDifficultyLevel_Call {
	return &OptionsRepository_GetAllDifficultyLevel_Call{Call: _e.mock.On("GetAllDifficultyLevel")}
}

func (_c *OptionsRepository_GetAllDifficultyLevel_Call) Run(run func()) *OptionsRepository_GetAllDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OptionsRepository_GetAllDifficultyLevel_Call) Return(_a0 []*model.DifficultyLevels, _a1 error) *OptionsRepository_GetAllDifficultyLevel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetAllDifficultyLevel_Call) RunAndReturn(run func() ([]*model.DifficultyLevels, error)) *OptionsRepository_GetAllDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllRace provides a mock function with no fields
func (_m *OptionsRepository) GetAllRace() ([]*model.Race, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllRace")
	}

	var r0 []*model.Race
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Race, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Race); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Race)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetAllRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllRace'
type OptionsRepository_GetAllRace_Call struct {
	*mock.Call
}

// GetAllRace is a helper method to define mock.On call
func (_e *OptionsRepository_Expecter) GetAllRace() *OptionsRepository_GetAllRace_Call {
	return &OptionsRepository_GetAllRace_Call{Call: _e.mock.On("GetAllRace")}
}

func (_c *OptionsRepository_GetAllRace_Call) Run(run func()) *OptionsRepository_GetAllRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OptionsRepository_GetAllRace_Call) Return(_a0 []*model.Race, _a1 error) *OptionsRepository_GetAllRace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetAllRace_Call) RunAndReturn(run func() ([]*model.Race, error)) *OptionsRepository_GetAllRace_Call {
	_c.Call.Return(run)
	return _c
}

// GetClassByID provides a mock function with given fields: id
func (_m *OptionsRepository) GetClassByID(id string) (*model.Class, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetClassByID")
	}

	var r0 *model.Class
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Class, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Class); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Class)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetClassByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClassByID'
type OptionsRepository_GetClassByID_Call struct {
	*mock.Call
}

// GetClassByID is a helper method to define mock.On call
//   - id string
func (_e *OptionsRepository_Expecter) GetClassByID(id interface{}) *OptionsRepository_GetClassByID_Call {
	return &OptionsRepository_GetClassByID_Call{Call: _e.mock.On("GetClassByID", id)}
}

func (_c *OptionsRepository_GetClassByID_Call) Run(run func(id string)) *OptionsRepository_GetClassByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_GetClassByID_Call) Return(_a0 *model.Class, _a1 error) *OptionsRepository_GetClassByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetClassByID_Call) RunAndReturn(run func(string) (*model.Class, error)) *OptionsRepository_GetClassByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetDifficultyLevelById provides a mock function with given fields: id
func (_m *OptionsRepository) GetDifficultyLevelById(id string) (*model.DifficultyLevels, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetDifficultyLevelById")
	}

	var r0 *model.DifficultyLevels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.DifficultyLevels, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.DifficultyLevels); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DifficultyLevels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetDifficultyLevelById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDifficultyLevelById'
type OptionsRepository_GetDifficultyLevelById_Call struct {
	*mock.Call
}

// GetDifficultyLevelById is a helper method to define mock.On call
//   - id string
func (_e *OptionsRepository_Expecter) GetDifficultyLevelById(id interface{}) *OptionsRepository_GetDifficultyLevelById_Call {
	return &OptionsRepository_GetDifficultyLevelById_Call{Call: _e.mock.On("GetDifficultyLevelById", id)}
}

func (_c *OptionsRepository_GetDifficultyLevelById_Call) Run(run func(id string)) *OptionsRepository_GetDifficultyLevelById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_GetDifficultyLevelById_Call) Return(_a0 *model.DifficultyLevels, _a1 error) *OptionsRepository_GetDifficultyLevelById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetDifficultyLevelById_Call) RunAndReturn(run func(string) (*model.DifficultyLevels, error)) *OptionsRepository_GetDifficultyLevelById_Call {
	_c.Call.Return(run)
	return _c
}

// GetRaceByID provides a mock function with given fields: id
func (_m *OptionsRepository) GetRaceByID(id string) (*model.Race, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetRaceByID")
	}

	var r0 *model.Race
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Race, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Race); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Race)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptionsRepository_GetRaceByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRaceByID'
type OptionsRepository_GetRaceByID_Call struct {
	*mock.Call
}

// GetRaceByID is a helper method to define mock.On call
//   - id string
func (_e *OptionsRepository_Expecter) GetRaceByID(id interface{}) *OptionsRepository_GetRaceByID_Call {
	return &OptionsRepository_GetRaceByID_Call{Call: _e.mock.On("GetRaceByID", id)}
}

func (_c *OptionsRepository_GetRaceByID_Call) Run(run func(id string)) *OptionsRepository_GetRaceByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OptionsRepository_GetRaceByID_Call) Return(_a0 *model.Race, _a1 error) *OptionsRepository_GetRaceByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OptionsRepository_GetRaceByID_Call) RunAndReturn(run func(string) (*model.Race, error)) *OptionsRepository_GetRaceByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateClass provides a mock function with given fields: class
func (_m *OptionsRepository) UpdateClass(class *model.Class) error {
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

// OptionsRepository_UpdateClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateClass'
type OptionsRepository_UpdateClass_Call struct {
	*mock.Call
}

// UpdateClass is a helper method to define mock.On call
//   - class *model.Class
func (_e *OptionsRepository_Expecter) UpdateClass(class interface{}) *OptionsRepository_UpdateClass_Call {
	return &OptionsRepository_UpdateClass_Call{Call: _e.mock.On("UpdateClass", class)}
}

func (_c *OptionsRepository_UpdateClass_Call) Run(run func(class *model.Class)) *OptionsRepository_UpdateClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Class))
	})
	return _c
}

func (_c *OptionsRepository_UpdateClass_Call) Return(_a0 error) *OptionsRepository_UpdateClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_UpdateClass_Call) RunAndReturn(run func(*model.Class) error) *OptionsRepository_UpdateClass_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDifficultyLevel provides a mock function with given fields: diffLevel
func (_m *OptionsRepository) UpdateDifficultyLevel(diffLevel *model.DifficultyLevels) error {
	ret := _m.Called(diffLevel)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDifficultyLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.DifficultyLevels) error); ok {
		r0 = rf(diffLevel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OptionsRepository_UpdateDifficultyLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDifficultyLevel'
type OptionsRepository_UpdateDifficultyLevel_Call struct {
	*mock.Call
}

// UpdateDifficultyLevel is a helper method to define mock.On call
//   - diffLevel *model.DifficultyLevels
func (_e *OptionsRepository_Expecter) UpdateDifficultyLevel(diffLevel interface{}) *OptionsRepository_UpdateDifficultyLevel_Call {
	return &OptionsRepository_UpdateDifficultyLevel_Call{Call: _e.mock.On("UpdateDifficultyLevel", diffLevel)}
}

func (_c *OptionsRepository_UpdateDifficultyLevel_Call) Run(run func(diffLevel *model.DifficultyLevels)) *OptionsRepository_UpdateDifficultyLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DifficultyLevels))
	})
	return _c
}

func (_c *OptionsRepository_UpdateDifficultyLevel_Call) Return(_a0 error) *OptionsRepository_UpdateDifficultyLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_UpdateDifficultyLevel_Call) RunAndReturn(run func(*model.DifficultyLevels) error) *OptionsRepository_UpdateDifficultyLevel_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRace provides a mock function with given fields: race
func (_m *OptionsRepository) UpdateRace(race *model.Race) error {
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

// OptionsRepository_UpdateRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRace'
type OptionsRepository_UpdateRace_Call struct {
	*mock.Call
}

// UpdateRace is a helper method to define mock.On call
//   - race *model.Race
func (_e *OptionsRepository_Expecter) UpdateRace(race interface{}) *OptionsRepository_UpdateRace_Call {
	return &OptionsRepository_UpdateRace_Call{Call: _e.mock.On("UpdateRace", race)}
}

func (_c *OptionsRepository_UpdateRace_Call) Run(run func(race *model.Race)) *OptionsRepository_UpdateRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Race))
	})
	return _c
}

func (_c *OptionsRepository_UpdateRace_Call) Return(_a0 error) *OptionsRepository_UpdateRace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OptionsRepository_UpdateRace_Call) RunAndReturn(run func(*model.Race) error) *OptionsRepository_UpdateRace_Call {
	_c.Call.Return(run)
	return _c
}

// NewOptionsRepository creates a new instance of OptionsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOptionsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OptionsRepository {
	mock := &OptionsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
