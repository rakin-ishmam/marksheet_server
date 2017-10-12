package access

import (
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
	"github.com/rakin-ishmam/marksheet_server/user"
)

// errInvalidRight generate invalid error for Right
func errInvalidRight(rt Right) error {
	return errs.InvalidErr(op.Validation("access", "right", rt))
}

// errParseUsrRight generate for parse user right
func errParseUsrRight(val string) error {
	return errs.InvalidErr(op.Parse("access", "userright", val))
}

func errMaxAddUser(name user.Name) error {
	return errs.LimitErr(op.Add("access", "user", name.String()))
}

func errAddInvUser(name user.Name) error {
	return errs.InvalidErr(op.Add("access", "user", name.String()))
}
