// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package htmlselector

import mock "github.com/stretchr/testify/mock"

// MockBuilder is an autogenerated mock type for the Builder type
type MockBuilder struct {
	mock.Mock
}

// AddAttribute provides a mock function with given fields: name, value
func (_m *MockBuilder) AddAttribute(name []byte, value []byte) {
	_m.Called(name, value)
}

// AddTag provides a mock function with given fields: name
func (_m *MockBuilder) AddTag(name []byte) {
	_m.Called(name)
}
