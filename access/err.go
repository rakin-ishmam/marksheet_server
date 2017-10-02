package access

import (
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

// errInvalidRight generate invalid error for Right
func errInvalidRight(rt Right) error {
	return errs.InvalidErr(op.Validation("access", "right", rt))
}

// errParseUsrRight generate for parse user right
func errParseUsrRight(val string) error {
	return errs.InvalidErr(op.Parse("access", "userright", val))
}
