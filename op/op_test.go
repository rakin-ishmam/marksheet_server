package op_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/op"
)

func TestRemove(t *testing.T) {
	tt := []struct {
		name  string
		value op.Operation
		exp   string
	}{
		{
			"string",
			op.Remove("user", "permission", "r"),
			"remove:user:permission:r",
		},
		{
			"int",
			op.Remove("user", "salary", 12),
			"remove:user:salary:12",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.value.Op())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	tt := []struct {
		name  string
		value op.Operation
		exp   string
	}{
		{
			"string",
			op.Add("user", "permission", "r"),
			"add:user:permission:r",
		},
		{
			"int",
			op.Add("user", "salary", 12),
			"add:user:salary:12",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.value.Op())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}

		})
	}
}

func TestParse(t *testing.T) {
	tt := []struct {
		name  string
		value op.Operation
		exp   string
	}{
		{
			"right",
			op.Parse("access", "right", "r"),
			"parse:access:right:r",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.value.Op())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}

		})
	}
}

func TestValidation(t *testing.T) {
	tt := []struct {
		name  string
		value op.Operation
		exp   string
	}{
		{
			"valiation",
			op.Validation("access", "right", "r"),
			"validation:access:right:r",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.value.Op())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}

		})
	}
}
