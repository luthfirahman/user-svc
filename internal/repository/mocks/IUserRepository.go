// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/luthfirahman/user-svc/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: _a0, _a1
func (_m *IUserRepository) FindByID(_a0 context.Context, _a1 int) (*repository.User, bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *repository.User
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*repository.User, bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByPhoneNumber provides a mock function with given fields: _a0, _a1
func (_m *IUserRepository) FindByPhoneNumber(_a0 context.Context, _a1 string) (*repository.User, bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *repository.User
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*repository.User, bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *IUserRepository) Save(_a0 context.Context, _a1 *repository.User) (*repository.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *repository.User) (*repository.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *repository.User) *repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *repository.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *IUserRepository) Update(_a0 context.Context, _a1 *repository.User) (*repository.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *repository.User) (*repository.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *repository.User) *repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *repository.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepository(t mockConstructorTestingTNewIUserRepository) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
