package user

import (
	"github.com/rakin-ishmam/marksheet_server/config"
)

// Name represents system's user nam
type Name string

// Valid validate name
func (n Name) Valid() bool {
	return ok(n.String())
}

func (n Name) String() string {
	return string(n)
}

// NewName convert string to Name
func NewName(str string) (Name, error) {
	nm := Name(str)
	if !nm.Valid() {
		return Name(""), errInvalidUsr(str)
	}

	return nm, nil
}

func ok(name string) bool {
	return okLen(name) && okStr(name) && !(name == GlobalUser().String())
}

func okLen(str string) bool {
	if l := len(str); l < config.MinUserName || l > config.MaxtUserName {
		return false
	}

	return true
}

func okStr(str string) bool {
	if str == GlobalUser().String() {
		return false
	}

	for _, c := range str {
		if !okChar(c) {
			return false
		}
	}

	return true
}

func okChar(c int32) bool {
	if c < 'a' || c > 'z' {
		return false
	}

	return true
}

// GlobalUser returns all users
func GlobalUser() Name {
	return "global"
}

// TestUser returns user name for test purpose
func TestUser() Name {
	return "globaltest"
}
