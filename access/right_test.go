package access_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/testformat"

	"github.com/rakin-ishmam/marksheet_server/access"
)

func TestValid(t *testing.T) {
	tt := []struct {
		name string
		val  access.Right
		exp  bool
	}{
		{"read right", access.Read, true},
		{"write right", access.Write, true},
		{"edit right", access.Edit, true},
		{"delete right", access.Delete, true},
		{"invalid right", 't', false},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(
				v.name,
				v.exp,
				v.val.IsRight(),
			)
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRightString(t *testing.T) {
	tt := []struct {
		name string
		val  access.Righter
		exp  string
	}{
		{"empty test", access.NewRights(), ""},
		{"all value test",
			access.NewRights(
				access.Read,
				access.Write,
				access.Edit,
				access.Delete),
			"rwed"},
		{"extrea read test",
			access.NewRights(
				access.Read,
				access.Read,
				access.Write,
				access.Edit,
				access.Delete),
			"rwed"},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := testformat.NewWithValue(v.name, v.exp, v.val.String())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestOp(t *testing.T) {
	var rts access.Righter
	opSeq := []struct {
		name string
		op   func()
		exp  string
	}{
		{
			name: "init",
			op: func() {
				rts = access.NewRights()
			},
			exp: "",
		},
		{
			name: "add read",
			op: func() {
				rts.Add(access.Read)
			},
			exp: "r",
		},
		{
			name: "add write",
			op: func() {
				rts.Add(access.Write)
			},
			exp: "rw",
		},
		{
			name: "add edit",
			op: func() {
				rts.Add(access.Edit)
			},
			exp: "rwe",
		},
		{
			name: "add delete",
			op: func() {
				rts.Add(access.Delete)
			},
			exp: "rwed",
		},
		{
			name: "remove write",
			op: func() {
				rts.Remove(access.Write)
			},
			exp: "red",
		},
		{
			name: "add write",
			op: func() {
				rts.Add(access.Write)
			},
			exp: "redw",
		},
		{
			name: "again add write",
			op: func() {
				rts.Add(access.Write)
			},
			exp: "redw",
		},
	}

	for _, v := range opSeq {
		t.Run(v.name, func(t *testing.T) {
			v.op()
			test := testformat.NewWithValue(v.name, v.exp, rts.String())
			if err := test.Test(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
