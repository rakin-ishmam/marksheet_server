package access

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
func (r *Rights) Add(rt Right) {
	if r.Has(rt) {
		return
	}

	*r = append(*r, rt)
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
func NewRights(rt ...Right) *Rights {
	rts := Rights{}
	for _, v := range rt {
		rts.Add(v)
	}

	return &rts
}
