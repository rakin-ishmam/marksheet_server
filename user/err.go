package user

import (
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

// ErrInvalidUsr generate error for invalid user name
func ErrInvalidUsr(val string) error {
	return errs.InvalidErr(op.Parse("user", "name", val))
}
