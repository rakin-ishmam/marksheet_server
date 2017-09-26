package access

import (
	"fmt"
	"strings"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
	"github.com/rakin-ishmam/marksheet_server/user"
)

// UserRight provide user rights
type UserRight struct {
	User   user.Name
	Rights Righter
}

// Has returns true if user has particular right, else returns false
func (u UserRight) Has(rt Right) bool {
	return u.Rights.Has(rt)
}

// Add particular Right to the user
func (u *UserRight) Add(rt Right) error {
	return u.Rights.Add(rt)
}

// Remove particular Right from the user
func (u *UserRight) Remove(rt Right) {
	u.Rights.Remove(rt)
}

// String convert user Rights to string
func (u *UserRight) String() string {
	return fmt.Sprintf("%v*%v", u.User.String(), u.Rights.String())
}

// NewUserRight return Righter of a User
func NewUserRight(name user.Name, rts Righter) Righter {
	return &UserRight{name, rts}
}

// ParseUserRight convert string to Righter of a user
func ParseUserRight(str string) (Righter, error) {
	divInd := strings.IndexAny(str, "*")
	if divInd < 0 {
		return nil, errs.InvalidErr(op.Parse("access", "userright", str))
	}

	user, err := user.NewName(str[:divInd])
	if err != nil {
		return nil, err
	}

	rights, err := ParseRights(str[divInd+1:])
	if err != nil {
		return nil, err
	}

	return NewUserRight(user, rights), nil
}
