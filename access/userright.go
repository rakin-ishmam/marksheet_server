package access

import (
	"fmt"
	"strings"

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
	user, err := parseUserName(str)
	if err != nil {
		return nil, err
	}

	rights, err := parseUserRights(str)
	if err != nil {
		return nil, err
	}

	return NewUserRight(user, rights), nil
}

func parseUserName(str string) (user.Name, error) {
	strs := strings.SplitN(str, "*", 2)
	if len(strs) < 2 {
		return "", errParseUsrRight(str)
	}

	return user.NewName(strs[0])
}

func parseUserRights(str string) (Righter, error) {
	strs := strings.SplitN(str, "*", 2)
	if len(strs) < 2 {
		return nil, errParseUsrRight(str)
	}

	return ParseRights(strs[1])
}
