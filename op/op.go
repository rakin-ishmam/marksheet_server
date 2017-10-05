package op

import (
	"fmt"

	"github.com/rakin-ishmam/marksheet_server/constant"
	"github.com/rakin-ishmam/marksheet_server/services/str"
)

// Operation represents operation of the system in specific format
type Operation interface {
	Op() string
}

type add struct {
	where string
	what  string
	value interface{}
}

func (a add) Op() string {
	return str.ConcatBySpliter(
		constant.OpSpliter,
		"add",
		a.where,
		a.what,
		fmt.Sprintf("%v", a.value),
	)

}

// Add returns add Operation
func Add(where, what string, value interface{}) Operation {
	return &add{
		where,
		what,
		value,
	}
}

type remove struct {
	where string
	what  string
	value interface{}
}

func (r remove) Op() string {
	return str.ConcatBySpliter(
		constant.OpSpliter,
		"remove",
		r.where,
		r.what,
		fmt.Sprintf("%v", r.value),
	)
}

// Remove returns remove operation
func Remove(where, what string, value interface{}) Operation {
	return &remove{
		where,
		what,
		value,
	}
}

type parse struct {
	where string
	what  string
	value interface{}
}

func (p parse) Op() string {
	return str.ConcatBySpliter(
		constant.OpSpliter,
		"parse",
		p.where,
		p.what,
		fmt.Sprintf("%v", p.value),
	)
}

// Parse returns parse operation
func Parse(where, what string, value interface{}) Operation {
	return &parse{
		where: where,
		what:  what,
		value: value,
	}
}

type validation struct {
	where string
	what  string
	value interface{}
}

func (v validation) Op() string {
	return str.ConcatBySpliter(
		constant.OpSpliter,
		"validation",
		v.where,
		v.what,
		fmt.Sprintf("%v", v.value),
	)
}

// Validation returns validation operation
func Validation(where, what string, value interface{}) Operation {
	return &validation{
		where: where,
		what:  what,
		value: value,
	}
}
