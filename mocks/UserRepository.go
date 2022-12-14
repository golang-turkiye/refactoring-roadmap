// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: userID
func (_m *MockUserRepository) GetUserByID(userID uint) (*domain.User, error) {
	ret := _m.Called(userID)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(uint) *domain.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserRepository(t testing.TB) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
