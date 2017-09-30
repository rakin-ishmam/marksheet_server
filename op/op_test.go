package op_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/op"
)

func TestRemove(t *testing.T) {
	at := []struct {
		name  string
		value op.Operation
		res   string
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

	for _, tt := range at {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.value.Op()
			if res != tt.res {
				t.Fatalf("%v expected %v but got %v", tt.name, tt.res, res)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	at := []struct {
		name  string
		value op.Operation
		res   string
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

	for _, tt := range at {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.value.Op()
			if res != tt.res {
				t.Fatalf("%v expected %v but got %v", tt.name, tt.res, res)
			}
		})
	}
}

func TestParse(t *testing.T) {
	at := []struct {
		name  string
		value op.Operation
		res   string
	}{
		{
			"right",
			op.Parse("access", "right", "r"),
			"parse:access:right:r",
		},
	}

	for _, tt := range at {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.value.Op()
			if res != tt.res {
				t.Fatalf("%v expected %v but got %v", tt.name, tt.res, res)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	at := []struct {
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

	for _, tt := range at {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.value.Op()
			if res != tt.exp {
				t.Fatalf("%v expected %v but got %v", tt.name, tt.exp, res)
			}
		})
	}
}
