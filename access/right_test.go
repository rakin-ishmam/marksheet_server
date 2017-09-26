package access_test

import (
	"testing"

	"github.com/rakin-ishmam/marksheet_server/access"
)

func TestValid(t *testing.T) {
	tt := []struct {
		name string
		val  access.Right
		res  bool
	}{
		{"read right", access.Read, true},
		{"write right", access.Write, true},
		{"edit right", access.Edit, true},
		{"delete right", access.Delete, true},
		{"invalid right", 't', false},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if res := v.val.Valid(); res != v.res {
				t.Fatalf("%v expected %v but got %v", v.name, v.res, res)
			}
		})
	}
}

func TestRightString(t *testing.T) {
	tt := []struct {
		name string
		val  access.Righter
		res  string
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
			if res := v.val.String(); res != v.res {
				t.Fatalf("%v expected %v but got %v", v.name, v.res, res)
			}
		})
	}
}

func TestOp(t *testing.T) {
	var rts access.Righter
	opSeq := []struct {
		name string
		op   func()
		val  string
	}{
		{
			name: "init",
			op: func() {
				rts = access.NewRights()
			},
			val: "",
		},
		{
			name: "add read",
			op: func() {
				rts.Add(access.Read)
			},
			val: "r",
		},
		{
			name: "add write",
			op: func() {
				rts.Add(access.Write)
			},
			val: "rw",
		},
		{
			name: "add edit",
			op: func() {
				rts.Add(access.Edit)
			},
			val: "rwe",
		},
		{
			name: "add delete",
			op: func() {
				rts.Add(access.Delete)
			},
			val: "rwed",
		},
		{
			name: "remove write",
			op: func() {
				rts.Remove(access.Write)
			},
			val: "red",
		},
		{
			name: "add write",
			op: func() {
				rts.Add(access.Write)
			},
			val: "redw",
		},
		{
			name: "again add write",
			op: func() {
				rts.Add(access.Write)
			},
			val: "redw",
		},
	}

	for _, v := range opSeq {
		t.Run(v.name, func(t *testing.T) {
			v.op()
			if v.val != rts.String() {
				t.Fatalf("%v expected %v but got %v", v.name, v.val, rts.String())
			}
		})
	}
}
