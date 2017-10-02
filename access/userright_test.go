package access

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestNewUserRight(t *testing.T) {
	uName, _ := user.NewName("test")
	rts := NewRights(Read, Write)
	ur := NewUserRight(uName, rts)

	exp := fmt.Sprintf("%v*%v", uName.String(), rts.String())
	res := ur.String()

	if exp != res {
		t.Fatalf("expectd %v but got %v", exp, res)
	}
}

func TestParseUserRight(t *testing.T) {
	tt := []struct {
		name  string
		value string
		exp   func() (Righter, error)
	}{
		{
			"valid",
			fmt.Sprintf("%v*%v", user.TestUser().String(), NewRights(Read, Write)),
			func() (Righter, error) {
				return NewUserRight(user.TestUser(), NewRights(Read, Write)), nil
			},
		},
		{
			"without devider",
			fmt.Sprintf("%v%v", user.TestUser().String(), NewRights(Read, Write)),
			func() (Righter, error) {
				return nil, errParseUsrRight(fmt.Sprintf("%v%v", user.TestUser().String(), NewRights(Read, Write)))
			},
		},
		{
			"invalid user name",
			fmt.Sprintf("%v*%v", "a", NewRights(Read, Write)),
			func() (Righter, error) {
				_, err := user.NewName("a")
				return nil, err
			},
		},
		{
			"invalid rights",
			fmt.Sprintf("%v*%v", user.TestUser().String(), "avb"),
			func() (Righter, error) {
				return nil, errInvalidRight(Right('a'))
			},
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			res, err := ParseUserRight(v.value)

			expRes, expErr := v.exp()

			testErr := testformat.NewWithValue(
				fmt.Sprintf("err->%v", v.name),
				expErr,
				err,
			)
			if err = testErr.Test(); err != nil {
				t.Fatalf(err.Error())
			}

			testV := testformat.New(
				fmt.Sprintf("value->%v", v.name),
				convRighterVF(expRes),
				convRighterVF(res),
			)

			if err = testV.Test(); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func convRighterVF(v Righter) testformat.ValueFunc {
	return func() interface{} {
		if v == nil {
			return nil
		}

		return v.String()
	}
}

func TestUserRightAdd(t *testing.T) {
	rts := NewRights(Read)
	r := NewUserRight(
		user.TestUser(),
		NewRights(Read),
	)

	r.Add(Write)
	rts.Add(Write)

	exp := fmt.Sprintf("%v*%v", user.TestUser(), rts)
	test := testformat.NewWithValue("Add", exp, r)

	if err := test.Test(); err != nil {
		t.Fatal(err.Error())
		return
	}
}

func TestUserRightHas(t *testing.T) {
	r := NewUserRight(
		user.TestUser(),
		NewRights(Read, Write),
	)

	tt := []struct {
		name string
		exp  bool
		res  bool
	}{
		{"read", true, r.Has(Read)},
		{"write", true, r.Has(Write)},
		{"delete", false, r.Has(Delete)},
		{"edit", false, r.Has(Edit)},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(
				v.name,
				v.exp,
				v.res,
			)

			if err := test.Test(); err != nil {
				t.Fatal(err.Error())
				return
			}
		})
	}
}

func TestUserRightRemove(t *testing.T) {
	rts := SuperRights()
	r := NewUserRight(
		user.TestUser(),
		SuperRights(),
	)

	tt := []struct {
		name string
		op   func()
	}{
		{"read", func() {
			r.Remove(Read)
			rts.Remove(Read)
		}},
		{"write", func() {
			r.Remove(Write)
			rts.Remove(Write)
		}},
		{"delete", func() {
			r.Remove(Delete)
			rts.Remove(Delete)
		}},
		{"edit", func() {
			r.Remove(Edit)
			rts.Remove(Edit)
		}},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			v.op()
			test := testformat.NewWithValue(
				v.name,
				fmt.Sprintf("%v*%v", user.TestUser(), rts),
				r,
			)

			if err := test.Test(); err != nil {
				t.Fatal(err.Error())
				return
			}
		})
	}
}
