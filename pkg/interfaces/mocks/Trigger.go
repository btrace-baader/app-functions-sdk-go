// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bootstrap "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap"

	interfaces "github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"

	mock "github.com/stretchr/testify/mock"

	sync "sync"
)

// Trigger is an autogenerated mock type for the Trigger type
type Trigger struct {
	mock.Mock
}

// Initialize provides a mock function with given fields: wg, ctx, background
func (_m *Trigger) Initialize(wg *sync.WaitGroup, ctx context.Context, background <-chan interfaces.BackgroundMessage) (bootstrap.Deferred, error) {
	ret := _m.Called(wg, ctx, background)

	var r0 bootstrap.Deferred
	if rf, ok := ret.Get(0).(func(*sync.WaitGroup, context.Context, <-chan interfaces.BackgroundMessage) bootstrap.Deferred); ok {
		r0 = rf(wg, ctx, background)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrap.Deferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sync.WaitGroup, context.Context, <-chan interfaces.BackgroundMessage) error); ok {
		r1 = rf(wg, ctx, background)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTrigger interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrigger creates a new instance of Trigger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrigger(t mockConstructorTestingTNewTrigger) *Trigger {
	mock := &Trigger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
