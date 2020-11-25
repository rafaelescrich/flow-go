// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	context "context"

	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Connector is an autogenerated mock type for the Connector type
type Connector struct {
	mock.Mock
}

// ConnectPeers provides a mock function with given fields: ctx, ids
func (_m *Connector) ConnectPeers(ctx context.Context, ids flow.IdentityList) map[flow.Identifier]error {
	ret := _m.Called(ctx, ids)

	var r0 map[flow.Identifier]error
	if rf, ok := ret.Get(0).(func(context.Context, flow.IdentityList) map[flow.Identifier]error); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[flow.Identifier]error)
		}
	}

	return r0
}

// DisconnectPeers provides a mock function with given fields: ctx, ids
func (_m *Connector) DisconnectPeers(ctx context.Context, ids flow.IdentityList) map[flow.Identifier]error {
	ret := _m.Called(ctx, ids)

	var r0 map[flow.Identifier]error
	if rf, ok := ret.Get(0).(func(context.Context, flow.IdentityList) map[flow.Identifier]error); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[flow.Identifier]error)
		}
	}

	return r0
}
