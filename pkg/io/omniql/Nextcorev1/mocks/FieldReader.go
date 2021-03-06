// Code generated by mockery v1.0.0
package mocks

import Nextcorev1 "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1"
import mock "github.com/stretchr/testify/mock"

// FieldReader is an autogenerated mock type for the FieldReader type
type FieldReader struct {
	mock.Mock
}

// Default provides a mock function with given fields:
func (_m *FieldReader) Default() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Documentation provides a mock function with given fields:
func (_m *FieldReader) Documentation() (Nextcorev1.DocumentationReader, error) {
	ret := _m.Called()

	var r0 Nextcorev1.DocumentationReader
	if rf, ok := ret.Get(0).(func() Nextcorev1.DocumentationReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Nextcorev1.DocumentationReader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *FieldReader) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *FieldReader) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
