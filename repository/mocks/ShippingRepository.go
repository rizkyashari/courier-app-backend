// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ShippingRepository is an autogenerated mock type for the ShippingRepository type
type ShippingRepository struct {
	mock.Mock
}

type mockConstructorTestingTNewShippingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewShippingRepository creates a new instance of ShippingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewShippingRepository(t mockConstructorTestingTNewShippingRepository) *ShippingRepository {
	mock := &ShippingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
