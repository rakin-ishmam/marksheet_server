package errs

import (
	"github.com/rakin-ishmam/marksheet_server/constant"
	"github.com/rakin-ishmam/marksheet_server/op"
	"github.com/rakin-ishmam/marksheet_server/services/str"
)

// Err represents system error structure
type Err struct {
	op   op.Operation
	kind Kind
}

func (e Err) Error() string {
	return str.ConcatBySpliter(constant.ErrSpliter, e.kind.String(), e.op.Op())
}

// Op gives Operation value
func (e Err) Op() string {
	return e.op.Op()
}

func genErr(op op.Operation, kind Kind) *Err {
	return &Err{op: op, kind: kind}
}

// InvalidErr returns error for invalid operation
func InvalidErr(op op.Operation) *Err {
	return genErr(op, Invalid)
}

// ExistErr returns error for putting something that is already exists
func ExistErr(op op.Operation) *Err {
	return genErr(op, Exist)
}

// NotExistErr returns error for operation on non exist item
func NotExistErr(op op.Operation) *Err {
	return genErr(op, NotExist)
}

// UnauthorisedErr returns error for operation on unauthorised action
func UnauthorisedErr(op op.Operation) *Err {
	return genErr(op, Unauthorised)
}
