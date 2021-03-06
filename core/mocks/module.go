// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	module "github.com/odpf/entropy/core/module"
	mock "github.com/stretchr/testify/mock"

	resource "github.com/odpf/entropy/core/resource"
)

// Module is an autogenerated mock type for the Module type
type Module struct {
	mock.Mock
}

type Module_Expecter struct {
	mock *mock.Mock
}

func (_m *Module) EXPECT() *Module_Expecter {
	return &Module_Expecter{mock: &_m.Mock}
}

// Plan provides a mock function with given fields: ctx, spec, act
func (_m *Module) Plan(ctx context.Context, spec module.Spec, act module.ActionRequest) (*resource.Resource, error) {
	ret := _m.Called(ctx, spec, act)

	var r0 *resource.Resource
	if rf, ok := ret.Get(0).(func(context.Context, module.Spec, module.ActionRequest) *resource.Resource); ok {
		r0 = rf(ctx, spec, act)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, module.Spec, module.ActionRequest) error); ok {
		r1 = rf(ctx, spec, act)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Module_Plan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Plan'
type Module_Plan_Call struct {
	*mock.Call
}

// Plan is a helper method to define mock.On call
//  - ctx context.Context
//  - spec module.Spec
//  - act module.ActionRequest
func (_e *Module_Expecter) Plan(ctx interface{}, spec interface{}, act interface{}) *Module_Plan_Call {
	return &Module_Plan_Call{Call: _e.mock.On("Plan", ctx, spec, act)}
}

func (_c *Module_Plan_Call) Run(run func(ctx context.Context, spec module.Spec, act module.ActionRequest)) *Module_Plan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(module.Spec), args[2].(module.ActionRequest))
	})
	return _c
}

func (_c *Module_Plan_Call) Return(_a0 *resource.Resource, _a1 error) *Module_Plan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Sync provides a mock function with given fields: ctx, spec
func (_m *Module) Sync(ctx context.Context, spec module.Spec) (*resource.State, error) {
	ret := _m.Called(ctx, spec)

	var r0 *resource.State
	if rf, ok := ret.Get(0).(func(context.Context, module.Spec) *resource.State); ok {
		r0 = rf(ctx, spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.State)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, module.Spec) error); ok {
		r1 = rf(ctx, spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Module_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type Module_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//  - ctx context.Context
//  - spec module.Spec
func (_e *Module_Expecter) Sync(ctx interface{}, spec interface{}) *Module_Sync_Call {
	return &Module_Sync_Call{Call: _e.mock.On("Sync", ctx, spec)}
}

func (_c *Module_Sync_Call) Run(run func(ctx context.Context, spec module.Spec)) *Module_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(module.Spec))
	})
	return _c
}

func (_c *Module_Sync_Call) Return(_a0 *resource.State, _a1 error) *Module_Sync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
