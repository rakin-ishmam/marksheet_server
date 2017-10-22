package prop_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/data/prop"
	"github.com/rakin-ishmam/marksheet_server/testformat"
)

func TestDirValid(t *testing.T) {
	strMax := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tt := []struct {
		name string
		res  bool
		exp  bool
	}{
		{"valid", prop.NewDirName("a").Valid(), true},
		{"invalid character", prop.NewDirName("9dfdf").Valid(), false},
		{"invalid length", prop.NewDirName("").Valid(), false},
		{"invalid length max", prop.NewDirName(strMax).Valid(), false},
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
