// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	fvm "github.com/dapperlabs/flow-go/fvm"
	mock "github.com/stretchr/testify/mock"
)

// TransactionProcessor is an autogenerated mock type for the TransactionProcessor type
type TransactionProcessor struct {
	mock.Mock
}

// Process provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *TransactionProcessor) Process(_a0 *fvm.VirtualMachine, _a1 fvm.Context, _a2 *fvm.TransactionProcedure, _a3 fvm.Ledger) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fvm.VirtualMachine, fvm.Context, *fvm.TransactionProcedure, fvm.Ledger) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}