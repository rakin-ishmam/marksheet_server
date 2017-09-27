package access

import (
	"github.com/rakin-ishmam/marksheet_server/errs"
	"github.com/rakin-ishmam/marksheet_server/op"
)

// Right represents permission of the path or file
type Right byte

// Posible rights
const (
	Read   Right = 'r'
	Write        = 'w'
	Edit         = 'e'
	Delete       = 'd'
)

// Valid validates the right
func (r Right) Valid() bool {
	switch r {
	case Read:
		return true
	case Write:
		return true
	case Edit:
		return true
	case Delete:
		return true
	}

	return false
}

// Rights is list of Right
type Rights []Right

// Has returns true if Right exist else false
func (r Rights) Has(rt Right) bool {
	for _, v := range r {
		if v == rt {
			return true
		}
	}

	return false
}

// Add new Right to Rights
func (r *Rights) Add(rt Right) error {
	if r.Has(rt) {
		return errs.ExistErr(op.Add("access", "right", rt))
	}

	*r = append(*r, rt)
	return nil
}

// Remove Right from Rights
func (r *Rights) Remove(rt Right) {
	ls := []Right(*r)

	for i, v := range ls {
		if rt == v {
			ls = append(ls[:i], ls[i+1:]...)
			break
		}
	}

	*r = ls
}

func (r Rights) String() string {
	bts := []byte{}

	for _, v := range []Right(r) {
		bts = append(bts, byte(v))
	}

	return string(bts)
}

// NewRights returns Rights
func NewRights(rt ...Right) Righter {
	rts := Rights{}
	for _, v := range rt {
		rts.Add(v)
	}

	return &rts
}

// ParseRights convert string to Rights
func ParseRights(str string) (Righter, error) {
	rts := Rights{}
	for _, v := range str {
		rt := Right(v)
		if !rt.Valid() {
			return nil, errs.InvalidErr(op.Parse("access", "right", str))
		}

		if err := rts.Add(rt); err != nil {
			return nil, err
		}
	}

	return &rts, nil
}

// SuperRights returns Righter with all Rights
func SuperRights() Righter {
	return NewRights(Read, Write, Edit, Delete)
}
