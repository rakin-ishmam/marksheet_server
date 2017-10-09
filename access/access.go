package access

import (
	"fmt"
	"strings"

	"github.com/rakin-ishmam/marksheet_server/user"
)

// Accessor wraps methods of access
type Accessor interface {
	Has(user.Name) Righter
	Add(user.Name) Righter
	Remove(user.Name)
	String() string
}

type access struct {
	usrRts map[string]Righter
	owner  user.Name
}

func (a *access) Remove(name user.Name) {
	delete(a.usrRts, name.String())
}

// Has takes User name and returns Righter of that user
func (a access) Has(name user.Name) Righter {

	if name.String() == a.owner.String() {
		return NewUserRight(a.owner, SuperRights())
	}

	r, has := a.usrRts[name.String()]
	if !has {
		return nil
	}

	return r
}

// Add user to access
func (a *access) Add(name user.Name) Righter {
	if r := a.Has(name); r != nil {
		return r
	}

	r := NewUserRight(name, NewRights())
	a.usrRts[name.String()] = r

	return r
}

// String convert users access to string
func (a access) String() string {
	added := false
	str := ""
	for _, v := range a.usrRts {
		if added {
			str += " "
		}
		added = true

		str += fmt.Sprintf("%v", v)
	}

	return str
}

// NewAccessor takes owner user and string, then parse the users and returns Accessor
func NewAccessor(owner user.Name, users string) (Accessor, error) {
	acs := &access{owner: owner, usrRts: make(map[string]Righter)}

	for _, v := range strings.Split(users, " ") {
		rts, err := ParseUserRight(v)

		if err != nil {
			return nil, err
		}

		name, err := parseUserName(v)
		if err != nil {
			return nil, err
		}

		acs.usrRts[name.String()] = rts
	}

	return acs, nil
}
