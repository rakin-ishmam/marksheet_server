package errs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rakin-ishmam/marksheet_server/constant"
	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

func TestErr(t *testing.T) {
	ts := []struct {
		name string
		kind errs.Kind
		op   op.Operation
	}{
		{"invalid", errs.Invalid, op.Add("user", "name", "test name")},
		{"exist", errs.Exist, op.Add("user", "name", "test name")},
		{"not exist", errs.NotExist, op.Add("user", "name", "test name")},
		{"unauthorised", errs.Unauthorised, op.Add("user", "name", "test name")},
	}

	for _, v := range ts {
		t.Run(v.name, func(t *testing.T) {
			expected := fmt.Sprintf("%v%v%v", v.kind.String(), constant.ErrSpliter, v.op.Op())
			res := genErr(v.op, v.kind).Error()
			test := testformat.NewWithValue(v.name, expected, res)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func genErr(op op.Operation, kind errs.Kind) error {
	switch kind {
	case errs.Invalid:
		return errs.InvalidErr(op)
	case errs.Exist:
		return errs.ExistErr(op)
	case errs.NotExist:
		return errs.NotExistErr(op)
	case errs.Unauthorised:
		return errs.UnauthorisedErr(op)
	}

	return errors.New("unknown")
}
