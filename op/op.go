package op

import "fmt"

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
	return fmt.Sprintf("add:%v:%v:%v", a.where, a.what, a.value)
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
	return fmt.Sprintf("remove:%v:%v:%v", r.where, r.what, r.value)
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
	return fmt.Sprintf("parse:%v:%v:%v", p.where, p.what, p.value)
}

// Parse returns parse operation
func Parse(where, what string, value interface{}) Operation {
	return &parse{
		where: where,
		what:  what,
		value: value,
	}
}
