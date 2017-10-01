package user_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"
	"github.com/rakin-ishmam/marksheet_server/user"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

func TestErrInvalidUsr(t *testing.T) {
	tt := []struct {
		name  string
		value string
		exp   error
	}{
		{"test 1", "name 1", errs.InvalidErr(op.Parse("user", "name", "name 1"))},
		{"test 2", "name 2", errs.InvalidErr(op.Parse("user", "name", "name 2"))},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, user.ErrInvalidUsr(v.value))
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
