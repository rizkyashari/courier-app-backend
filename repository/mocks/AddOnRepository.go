// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AddOnRepository is an autogenerated mock type for the AddOnRepository type
type AddOnRepository struct {
	mock.Mock
}

type mockConstructorTestingTNewAddOnRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAddOnRepository creates a new instance of AddOnRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAddOnRepository(t mockConstructorTestingTNewAddOnRepository) *AddOnRepository {
	mock := &AddOnRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
