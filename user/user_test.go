package user

import (
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"
)

func TestValid(t *testing.T) {
	tt := []struct {
		name string
		val  Name
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
			val:  TestUser(),
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
			test := testformat.NewWithValue(
				v.name,
				v.val.Valid(),
				v.res,
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
			errInvalidUsr("tes$tname"),
		},
	}

	for _, v := range st {
		t.Run(v.name, func(t *testing.T) {
			nm, err := NewName(v.value)
			testStr := testformat.NewWithValue(
				fmt.Sprintf("err->%v", v.name),
				v.okRes,
				nm.String(),
			)
			if err := testStr.Test(); err != nil {
				t.Fatal(err.Error())
			}

			testErr := testformat.NewWithValue(
				fmt.Sprintf("value->%v", v.name),
				v.errRes,
				err,
			)
			if err := testErr.Test(); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
