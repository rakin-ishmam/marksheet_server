package user_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"

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
			name: "valid",
			val:  "usertest",
			res:  true,
		},
	}

	for _, tcase := range tt {
		if res := tcase.val.Valid(); res != tcase.res {
			t.Run(tcase.name, func(t *testing.T) {
				t.Fatalf("%s expected %v but got %v", tcase.name, tcase.res, res)
			})
		}

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
			if nm.String() != v.okRes {
				t.Fatalf("New name should be (%v) but %v", v.okRes, nm)
			}
			if err == nil && v.errRes != nil {
				t.Fatalf("%v expected err (%v) but got nil", v.name, v.errRes)
			}
			if err != nil && v.errRes != nil {
				if err.Error() != v.errRes.Error() {
					t.Fatalf("%v expected err (%v) but got %v", v.name, v.errRes, err)
				}
			}
		})
	}
}
