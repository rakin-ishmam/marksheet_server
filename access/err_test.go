package access

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
	"github.com/rakin-ishmam/marksheet_server/testformat"
)

func TestErrInvalidRight(t *testing.T) {
	tt := []struct {
		name  string
		value Right
		exp   error
	}{
		{"test 1", Right('a'), errs.InvalidErr(op.Validation("access", "right", Right('a')))},
		{"test 2", Right('b'), errs.InvalidErr(op.Validation("access", "right", Right('b')))},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, errInvalidRight(v.value))
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestErrParseUsrRight(t *testing.T) {
	tt := []struct {
		name  string
		value string
		exp   error
	}{
		{"test 1", "usrrighterr1", errs.InvalidErr(op.Parse("access", "userright", "usrrighterr1"))},
		{"test 2", "usrrighterr2", errs.InvalidErr(op.Parse("access", "userright", "usrrighterr2"))},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, errParseUsrRight(v.value))
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
