package access

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"
	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestNewAccessor(t *testing.T) {
	tt := []struct {
		name  string
		owner user.Name
		users string

		expErr error
	}{
		{
			"valid1",
			user.GlobalUser(),
			fmt.Sprintf("%v", NewUserRight(user.TestUser(), SuperRights())),
			nil,
		},

		{
			"valid2",
			user.GlobalUser(),
			fmt.Sprintf(
				"%v u%v",
				NewUserRight(user.TestUser(), SuperRights()),
				NewUserRight(user.TestUser(), SuperRights()),
			),
			nil,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			accessor, err := NewAccessor(v.owner, v.users)
			test := testformat.NewWithValue(
				fmt.Sprintf("err %v", v.name),
				v.expErr,
				err,
			)
			if terr := test.Test(); terr != nil {
				t.Fatal(terr.Error())
				return
			}

			test = testformat.NewWithValue(
				fmt.Sprintf("value %v", v.name),
				v.users,
				accessor.String(),
			)

			if terr := test.Test(); terr != nil {
				t.Fatal(terr.Error())
			}

		})
	}
}

func TestAccessHas(t *testing.T) {
	acc, err := NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v %v",
			NewUserRight(user.TestUser(), NewRights(Read, Write)),
			NewUserRight(user.TestUser()+"t", NewRights(Read, Write)),
		),
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	tt := []struct {
		name string
		res  testformat.ValueFunc
		exp  testformat.ValueFunc
	}{
		{
			"test1",
			func() interface{} {
				rt := acc.Has(user.TestUser())
				return rt
			},
			func() interface{} {
				return NewUserRight(user.TestUser(), NewRights(Read, Write))
			},
		},
		{
			"test2 nil",
			func() interface{} {
				rt := acc.Has(user.TestUser() + "m")
				return rt
			},
			func() interface{} {
				return nil
			},
		},
		{
			"test3 owner",
			func() interface{} {
				rt := acc.Has(user.GlobalUser())
				return rt
			},
			func() interface{} {
				return NewUserRight(user.GlobalUser(), SuperRights())
			},
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.New(v.name, v.exp, v.res)
			if err = test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestAccessAdd(t *testing.T) {
	acc, err := NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v",
			NewUserRight(user.TestUser(), NewRights(Read, Write)),
		),
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	tt := []struct {
		name   string
		user   user.Name
		expRes Righter
		expErr error
	}{
		{
			"invalid",
			user.Name("t"),
			nil,
			errAddInvUser(user.Name("t")),
		},
		{
			"already exist",
			user.TestUser(),
			NewUserRight(user.TestUser(), NewRights(Read, Write)),
			nil,
		},
		{
			"new user",
			user.TestUser() + "v",
			NewUserRight(user.TestUser()+"v", NewRights()),
			nil,
		},
	}

	for _, v := range tt {
		res, err := acc.Add(v.user)

		test := testformat.NewEmpty()
		test.Add("value test-"+v.name, v.expRes, res)
		test.Add("err test-"+v.name, v.expErr, err)

		if err = test.Test(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAccessRemove(t *testing.T) {
	acc, err := NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v %v",
			NewUserRight(user.TestUser(), NewRights(Read, Write)),
			NewUserRight(user.TestUser()+"t", NewRights(Read, Write)),
		),
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	tt := []struct {
		name string
		op   func()
		res  testformat.ValueFunc
		exp  testformat.ValueFunc
	}{
		{
			"remove test1",
			func() {
				acc.Remove(user.TestUser())
			},
			func() interface{} {
				return acc
			},
			func() interface{} {
				return fmt.Sprintf(
					"%v",
					NewUserRight(user.TestUser()+"t", NewRights(Read, Write)),
				)
			},
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			v.op()
			test := testformat.New(v.name, v.exp, v.res)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
