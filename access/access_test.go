package access_test

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/access"
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
			fmt.Sprintf("%v", access.NewUserRight(user.TestUser(), access.SuperRights())),
			nil,
		},

		{
			"valid2",
			user.GlobalUser(),
			fmt.Sprintf(
				"%v u%v",
				access.NewUserRight(user.TestUser(), access.SuperRights()),
				access.NewUserRight(user.TestUser(), access.SuperRights()),
			),
			nil,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			accessor, err := access.NewAccessor(v.owner, v.users)
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
	acc, err := access.NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v %v",
			access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
			access.NewUserRight(user.TestUser()+"t", access.NewRights(access.Read, access.Write)),
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
				return access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write))
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
				return access.NewUserRight(user.GlobalUser(), access.SuperRights())
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
	acc, err := access.NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v",
			access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
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
			"test1",
			func() {
				rt := acc.Add(user.TestUser() + "t")
				rt.Add(access.Read)
				rt.Add(access.Write)
			},
			func() interface{} {
				return acc
			},
			func() interface{} {
				return fmt.Sprintf(
					"%v %v",
					access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
					access.NewUserRight(user.TestUser()+"t", access.NewRights(access.Read, access.Write)),
				)
			},
		},
		{
			"test2",
			func() {
				rt := acc.Add(user.TestUser() + "t")
				rt.Add(access.Delete)
			},
			func() interface{} {
				return acc
			},
			func() interface{} {
				return fmt.Sprintf(
					"%v %v",
					access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
					access.NewUserRight(user.TestUser()+"t", access.NewRights(access.Read, access.Write, access.Delete)),
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

func TestAccessRemove(t *testing.T) {
	acc, err := access.NewAccessor(
		user.GlobalUser(),
		fmt.Sprintf(
			"%v %v",
			access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
			access.NewUserRight(user.TestUser()+"t", access.NewRights(access.Read, access.Write)),
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
					access.NewUserRight(user.TestUser()+"t", access.NewRights(access.Read, access.Write)),
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
