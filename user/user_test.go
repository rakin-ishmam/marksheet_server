package user_test

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/user"
)

func TestValid(t *testing.T) {
	tt := []struct {
		name string
		val  user.Name
		res  bool
	}{
		{
			name: "min len",
			val:  "r",
			res:  false,
		},
		{
			name: "max len",
			val: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			res: false,
		},
		{
			name: "invalid char",
			val:  "_testtest",
			res:  false,
		},
		{
			name: "global user",
			val:  "global",
			res:  false,
		},
		{
			name: "test user",
			val:  user.TestUser(),
			res:  true,
		},
		{
			name: "valid",
			val:  "usertest",
			res:  true,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewTest(
				v.name,
				testformat.ConvVF(v.val.Valid()),
				testformat.ConvVF(v.res),
			)
			if err := test.Test(); err != nil {
				t.Fatalf(err.Error())
			}

		})
	}
}

func TestNewName(t *testing.T) {
	st := []struct {
		name   string
		value  string
		okRes  string
		errRes error
	}{
		{
			"valid",
			"testname",
			"testname",
			nil,
		},
		{
			"invalid",
			"tes$tname",
			"",
			errs.InvalidErr(op.Parse("user", "name", "tes$tname")),
		},
	}

	for _, v := range st {
		t.Run(v.name, func(t *testing.T) {
			nm, err := user.NewName(v.value)
			testStr := testformat.NewTest(
				fmt.Sprintf("err->%v", v.name),
				testformat.ConvVF(v.okRes),
				testformat.ConvVF(nm.String()),
			)
			if err := testStr.Test(); err != nil {
				t.Fatal(err.Error())
			}

			testErr := testformat.NewTest(
				fmt.Sprintf("value->%v", v.name),
				testformat.ConvVF(v.errRes),
				testformat.ConvVF(err),
			)
			if err := testErr.Test(); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
