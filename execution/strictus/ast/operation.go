package ast

type Operation int

const (
	OperationOr Operation = iota
	OperationAnd
	OperationEqual
	OperationUnequal
	OperationLess
	OperationGreater
	OperationLessEqual
	OperationGreaterEqual
	OperationPlus
	OperationMinus
	OperationMul
	OperationDiv
	OperationMod
)
