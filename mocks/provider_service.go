// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/entropy/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProviderService is an autogenerated mock type for the ServiceInterface type
type ProviderService struct {
	mock.Mock
}

type ProviderService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderService) EXPECT() *ProviderService_Expecter {
	return &ProviderService_Expecter{mock: &_m.Mock}
}

// CreateProvider provides a mock function with given fields: ctx, res
func (_m *ProviderService) CreateProvider(ctx context.Context, res *domain.Provider) (*domain.Provider, error) {
	ret := _m.Called(ctx, res)

	var r0 *domain.Provider
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Provider) *domain.Provider); ok {
		r0 = rf(ctx, res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Provider) error); ok {
		r1 = rf(ctx, res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderService_CreateProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProvider'
type ProviderService_CreateProvider_Call struct {
	*mock.Call
}

// CreateProvider is a helper method to define mock.On call
//  - ctx context.Context
//  - res *domain.Provider
func (_e *ProviderService_Expecter) CreateProvider(ctx interface{}, res interface{}) *ProviderService_CreateProvider_Call {
	return &ProviderService_CreateProvider_Call{Call: _e.mock.On("CreateProvider", ctx, res)}
}

func (_c *ProviderService_CreateProvider_Call) Run(run func(ctx context.Context, res *domain.Provider)) *ProviderService_CreateProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Provider))
	})
	return _c
}

func (_c *ProviderService_CreateProvider_Call) Return(_a0 *domain.Provider, _a1 error) *ProviderService_CreateProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListProviders provides a mock function with given fields: ctx, parent, kind
func (_m *ProviderService) ListProviders(ctx context.Context, parent string, kind string) ([]*domain.Provider, error) {
	ret := _m.Called(ctx, parent, kind)

	var r0 []*domain.Provider
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*domain.Provider); ok {
		r0 = rf(ctx, parent, kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, parent, kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderService_ListProviders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProviders'
type ProviderService_ListProviders_Call struct {
	*mock.Call
}

// ListProviders is a helper method to define mock.On call
//  - ctx context.Context
//  - parent string
//  - kind string
func (_e *ProviderService_Expecter) ListProviders(ctx interface{}, parent interface{}, kind interface{}) *ProviderService_ListProviders_Call {
	return &ProviderService_ListProviders_Call{Call: _e.mock.On("ListProviders", ctx, parent, kind)}
}

func (_c *ProviderService_ListProviders_Call) Run(run func(ctx context.Context, parent string, kind string)) *ProviderService_ListProviders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProviderService_ListProviders_Call) Return(_a0 []*domain.Provider, _a1 error) *ProviderService_ListProviders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
