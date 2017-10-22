package prop_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/data/prop"
	"github.com/rakin-ishmam/marksheet_server/testformat"
)

func TestCellValid(t *testing.T) {
	tt := []struct {
		name string
		res  bool
		exp  bool
	}{
		{"valid:function", prop.Function.Valid(), true},
		{"valid:value", prop.Value.Valid(), true},
		{"invlid ", prop.Cell('c').Valid(), false},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.res)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}

}

func TestResValid(t *testing.T) {
	tt := []struct {
		name string
		res  bool
		exp  bool
	}{
		{"valid:dir", prop.Dir.Valid(), true},
		{"valid:sheet", prop.Sheet.Valid(), true},
		{"invalid", prop.Resource("dfd").Valid(), false},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.res)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
