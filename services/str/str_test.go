package str_test

import "testing"
import "github.com/rakin-ishmam/marksheet_server/services/str"
import "github.com/rakin-ishmam/marksheet_server/testformat"

func TestConcatBySpliter(t *testing.T) {
	ts := []struct {
		name    string
		vals    []string
		spliter string
		exp     string
	}{
		{"test1", []string{"a", "b", "c"}, ":", "a:b:c"},
		{"test2", []string{"a", "b", "c"}, "#", "a#b#c"},
	}

	for _, v := range ts {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(
				v.name,
				v.exp,
				str.ConcatBySpliter(v.spliter, v.vals...),
			)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
