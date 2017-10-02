package user

import (
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

// errInvalidUsr generate error for invalid user name
func errInvalidUsr(val string) error {
	return errs.InvalidErr(op.Parse("user", "name", val))
}
