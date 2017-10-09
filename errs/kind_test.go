package errs_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/errs"
)

func TestKind(t *testing.T) {
	ts := []struct {
		name  string
		value errs.Kind
		res   string
	}{
		{"invalid", errs.Invalid, "invalid"},
		{"exist", errs.Exist, "exist"},
		{"not exit", errs.NotExist, "not exit"},
		{"unauthorised", errs.Unauthorised, "unauthorised"},
		{"limit", errs.Limit, "limit"},
		{"unknown", 100, "unknown"},
	}

	for _, v := range ts {
		t.Run(v.name, func(t *testing.T) {
			if res := v.value.String(); res != v.res {
				t.Fatalf("%v test expected %v but got %v", v.name, v.res, res)
			}
		})
	}
}
