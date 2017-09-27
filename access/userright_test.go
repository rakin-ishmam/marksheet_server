package access_test

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/access"
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"

	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestNewUserRight(t *testing.T) {
	uName, _ := user.NewName("test")
	rts := access.NewRights(access.Read, access.Write)
	ur := access.NewUserRight(uName, rts)

	exp := fmt.Sprintf("%v*%v", uName.String(), rts.String())
	res := ur.String()

	if exp != res {
		t.Fatalf("expectd %v but got %v", exp, res)
	}
}

func TestParseUserRight(t *testing.T) {
	tt := []struct {
		name       string
		value      string
		resRighter access.Righter
		resErr     error
	}{
		{
			"valid",
			fmt.Sprintf("%v*%v", user.TestUser().String(), access.NewRights(access.Read, access.Write)),
			access.NewUserRight(user.TestUser(), access.NewRights(access.Read, access.Write)),
			nil,
		},
		{
			"without devider",
			fmt.Sprintf("%v%v", user.TestUser().String(), access.NewRights(access.Read, access.Write)),
			nil,
			errs.InvalidErr(op.Parse("access", "userright", fmt.Sprintf("%v%v", user.TestUser().String(), access.NewRights(access.Read, access.Write)))),
		},
		{
			"invalid user name",
			fmt.Sprintf("%v*%v", "a", access.NewRights(access.Read, access.Write)),
			nil,
			errs.InvalidErr(op.Parse("user", "name", "a")),
		},
		{
			"invalid rights",
			fmt.Sprintf("%v*%v", user.TestUser().String(), "avb"),
			nil,
			errs.InvalidErr(op.Parse("access", "right", "avb")),
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			res, err := access.ParseUserRight(v.value)

			testErr := testformat.NewTest(
				fmt.Sprintf("err->%v", v.name),
				testformat.ConvVF(v.resErr),
				testformat.ConvVF(err),
			)
			if err = testErr.Test(); err != nil {
				t.Fatalf(err.Error())
			}

			testV := testformat.NewTest(
				fmt.Sprintf("value->%v", v.name),
				convRighterVF(v.resRighter),
				convRighterVF(res),
			)

			if err = testV.Test(); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func convRighterVF(v access.Righter) testformat.ValueFunc {
	return func() interface{} {
		if v == nil {
			return nil
		}

		return v.String()
	}
}

func TestUserRightAdd(t *testing.T) {
	rts := access.NewRights(access.Read)
	r := access.NewUserRight(
		user.TestUser(),
		access.NewRights(access.Read),
	)

	r.Add(access.Write)
	rts.Add(access.Write)

	exp := fmt.Sprintf("%v*%v", user.TestUser(), rts)
	test := testformat.NewTest("Add", testformat.ConvVF(exp), testformat.ConvVF(r))

	if err := test.Test(); err != nil {
		t.Fatal(err.Error())
		return
	}
}

func TestUserRightHas(t *testing.T) {
	r := access.NewUserRight(
		user.TestUser(),
		access.NewRights(access.Read, access.Write),
	)

	tt := []struct {
		name string
		exp  bool
		res  bool
	}{
		{"read", true, r.Has(access.Read)},
		{"write", true, r.Has(access.Write)},
		{"delete", false, r.Has(access.Delete)},
		{"edit", false, r.Has(access.Edit)},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewTest(
				v.name,
				testformat.ConvVF(v.exp),
				testformat.ConvVF(v.res),
			)

			if err := test.Test(); err != nil {
				t.Fatal(err.Error())
				return
			}
		})
	}
}

func TestUserRightRemove(t *testing.T) {
	rts := access.SuperRights()
	r := access.NewUserRight(
		user.TestUser(),
		access.SuperRights(),
	)

	tt := []struct {
		name string
		op   func()
	}{
		{"read", func() {
			r.Remove(access.Read)
			rts.Remove(access.Read)
		}},
		{"write", func() {
			r.Remove(access.Write)
			rts.Remove(access.Write)
		}},
		{"delete", func() {
			r.Remove(access.Delete)
			rts.Remove(access.Delete)
		}},
		{"edit", func() {
			r.Remove(access.Edit)
			rts.Remove(access.Edit)
		}},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			v.op()
			test := testformat.NewTest(
				v.name,
				testformat.ConvVF(fmt.Sprintf("%v*%v", user.TestUser(), rts)),
				testformat.ConvVF(r),
			)

			if err := test.Test(); err != nil {
				t.Fatal(err.Error())
				return
			}
		})
	}
}
