// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "tech-test/internal/author/entity"

	mock "github.com/stretchr/testify/mock"
)

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *AuthorRepository) Get(_a0 string) (entity.Author, error) {
	ret := _m.Called(_a0)

	var r0 entity.Author
	if rf, ok := ret.Get(0).(func(string) entity.Author); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Author)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthorRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthorRepository creates a new instance of AuthorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorRepository(t mockConstructorTestingTNewAuthorRepository) *AuthorRepository {
	mock := &AuthorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
