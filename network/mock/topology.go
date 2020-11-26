// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Topology is an autogenerated mock type for the Topology type
type Topology struct {
	mock.Mock
}

// GenerateFanout provides a mock function with given fields: ids
func (_m *Topology) GenerateFanout(ids flow.IdentityList) (flow.IdentityList, error) {
	ret := _m.Called(ids)

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func(flow.IdentityList) flow.IdentityList); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.IdentityList) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
