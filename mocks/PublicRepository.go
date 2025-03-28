// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	model "d-and-d/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// PublicRepository is an autogenerated mock type for the PublicRepository type
type PublicRepository struct {
	mock.Mock
}

type PublicRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *PublicRepository) EXPECT() *PublicRepository_Expecter {
	return &PublicRepository_Expecter{mock: &_m.Mock}
}

// GetPublicCharacter provides a mock function with no fields
func (_m *PublicRepository) GetPublicCharacter() ([]*model.Character, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPublicCharacter")
	}

	var r0 []*model.Character
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Character, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Character); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Character)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicRepository_GetPublicCharacter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPublicCharacter'
type PublicRepository_GetPublicCharacter_Call struct {
	*mock.Call
}

// GetPublicCharacter is a helper method to define mock.On call
func (_e *PublicRepository_Expecter) GetPublicCharacter() *PublicRepository_GetPublicCharacter_Call {
	return &PublicRepository_GetPublicCharacter_Call{Call: _e.mock.On("GetPublicCharacter")}
}

func (_c *PublicRepository_GetPublicCharacter_Call) Run(run func()) *PublicRepository_GetPublicCharacter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PublicRepository_GetPublicCharacter_Call) Return(_a0 []*model.Character, _a1 error) *PublicRepository_GetPublicCharacter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicRepository_GetPublicCharacter_Call) RunAndReturn(run func() ([]*model.Character, error)) *PublicRepository_GetPublicCharacter_Call {
	_c.Call.Return(run)
	return _c
}

// GetPublicQuest provides a mock function with no fields
func (_m *PublicRepository) GetPublicQuest() ([]*model.Quest, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPublicQuest")
	}

	var r0 []*model.Quest
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Quest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Quest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Quest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicRepository_GetPublicQuest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPublicQuest'
type PublicRepository_GetPublicQuest_Call struct {
	*mock.Call
}

// GetPublicQuest is a helper method to define mock.On call
func (_e *PublicRepository_Expecter) GetPublicQuest() *PublicRepository_GetPublicQuest_Call {
	return &PublicRepository_GetPublicQuest_Call{Call: _e.mock.On("GetPublicQuest")}
}

func (_c *PublicRepository_GetPublicQuest_Call) Run(run func()) *PublicRepository_GetPublicQuest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PublicRepository_GetPublicQuest_Call) Return(_a0 []*model.Quest, _a1 error) *PublicRepository_GetPublicQuest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicRepository_GetPublicQuest_Call) RunAndReturn(run func() ([]*model.Quest, error)) *PublicRepository_GetPublicQuest_Call {
	_c.Call.Return(run)
	return _c
}

// NewPublicRepository creates a new instance of PublicRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPublicRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PublicRepository {
	mock := &PublicRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
