// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocknetwork

import (
	network "github.com/onflow/flow-go/network"
	mock "github.com/stretchr/testify/mock"
)

// SubscriptionManager is an autogenerated mock type for the SubscriptionManager type
type SubscriptionManager struct {
	mock.Mock
}

// GetChannelIDs provides a mock function with given fields:
func (_m *SubscriptionManager) GetChannelIDs() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetEngine provides a mock function with given fields: channelID
func (_m *SubscriptionManager) GetEngine(channelID string) (network.Engine, error) {
	ret := _m.Called(channelID)

	var r0 network.Engine
	if rf, ok := ret.Get(0).(func(string) network.Engine); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(network.Engine)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: channelID, engine
func (_m *SubscriptionManager) Register(channelID string, engine network.Engine) error {
	ret := _m.Called(channelID, engine)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, network.Engine) error); ok {
		r0 = rf(channelID, engine)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: channelID
func (_m *SubscriptionManager) Unregister(channelID string) error {
	ret := _m.Called(channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
