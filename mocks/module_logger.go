// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/entropy/domain"
	mock "github.com/stretchr/testify/mock"
)

// ModuleLogger is an autogenerated mock type for the ModuleLogger type
type ModuleLogger struct {
	mock.Mock
}

type ModuleLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *ModuleLogger) EXPECT() *ModuleLogger_Expecter {
	return &ModuleLogger_Expecter{mock: &_m.Mock}
}

// Act provides a mock function with given fields: r, action, params
func (_m *ModuleLogger) Act(r *domain.Resource, action string, params map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(r, action, params)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(*domain.Resource, string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(r, action, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Resource, string, map[string]interface{}) error); ok {
		r1 = rf(r, action, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleLogger_Act_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Act'
type ModuleLogger_Act_Call struct {
	*mock.Call
}

// Act is a helper method to define mock.On call
//  - r *domain.Resource
//  - action string
//  - params map[string]interface{}
func (_e *ModuleLogger_Expecter) Act(r interface{}, action interface{}, params interface{}) *ModuleLogger_Act_Call {
	return &ModuleLogger_Act_Call{Call: _e.mock.On("Act", r, action, params)}
}

func (_c *ModuleLogger_Act_Call) Run(run func(r *domain.Resource, action string, params map[string]interface{})) *ModuleLogger_Act_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Resource), args[1].(string), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *ModuleLogger_Act_Call) Return(_a0 map[string]interface{}, _a1 error) *ModuleLogger_Act_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Apply provides a mock function with given fields: r
func (_m *ModuleLogger) Apply(r *domain.Resource) (domain.ResourceStatus, error) {
	ret := _m.Called(r)

	var r0 domain.ResourceStatus
	if rf, ok := ret.Get(0).(func(*domain.Resource) domain.ResourceStatus); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(domain.ResourceStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Resource) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleLogger_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ModuleLogger_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//  - r *domain.Resource
func (_e *ModuleLogger_Expecter) Apply(r interface{}) *ModuleLogger_Apply_Call {
	return &ModuleLogger_Apply_Call{Call: _e.mock.On("Apply", r)}
}

func (_c *ModuleLogger_Apply_Call) Run(run func(r *domain.Resource)) *ModuleLogger_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Resource))
	})
	return _c
}

func (_c *ModuleLogger_Apply_Call) Return(_a0 domain.ResourceStatus, _a1 error) *ModuleLogger_Apply_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ID provides a mock function with given fields:
func (_m *ModuleLogger) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ModuleLogger_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type ModuleLogger_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *ModuleLogger_Expecter) ID() *ModuleLogger_ID_Call {
	return &ModuleLogger_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *ModuleLogger_ID_Call) Run(run func()) *ModuleLogger_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ModuleLogger_ID_Call) Return(_a0 string) *ModuleLogger_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

// Log provides a mock function with given fields: ctx, r, filter
func (_m *ModuleLogger) Log(ctx context.Context, r *domain.Resource, filter map[string]string) (<-chan domain.LogChunk, error) {
	ret := _m.Called(ctx, r, filter)

	var r0 <-chan domain.LogChunk
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Resource, map[string]string) <-chan domain.LogChunk); ok {
		r0 = rf(ctx, r, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan domain.LogChunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Resource, map[string]string) error); ok {
		r1 = rf(ctx, r, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleLogger_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type ModuleLogger_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
//  - ctx context.Context
//  - r *domain.Resource
//  - filter map[string]string
func (_e *ModuleLogger_Expecter) Log(ctx interface{}, r interface{}, filter interface{}) *ModuleLogger_Log_Call {
	return &ModuleLogger_Log_Call{Call: _e.mock.On("Log", ctx, r, filter)}
}

func (_c *ModuleLogger_Log_Call) Run(run func(ctx context.Context, r *domain.Resource, filter map[string]string)) *ModuleLogger_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Resource), args[2].(map[string]string))
	})
	return _c
}

func (_c *ModuleLogger_Log_Call) Return(_a0 <-chan domain.LogChunk, _a1 error) *ModuleLogger_Log_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Validate provides a mock function with given fields: r
func (_m *ModuleLogger) Validate(r *domain.Resource) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Resource) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModuleLogger_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type ModuleLogger_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//  - r *domain.Resource
func (_e *ModuleLogger_Expecter) Validate(r interface{}) *ModuleLogger_Validate_Call {
	return &ModuleLogger_Validate_Call{Call: _e.mock.On("Validate", r)}
}

func (_c *ModuleLogger_Validate_Call) Run(run func(r *domain.Resource)) *ModuleLogger_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Resource))
	})
	return _c
}

func (_c *ModuleLogger_Validate_Call) Return(_a0 error) *ModuleLogger_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}