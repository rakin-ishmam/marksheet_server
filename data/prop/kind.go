package prop

// Cell retpresents cell kind of sheet
type Cell byte

// Different kind of cell
const (
	Function Cell = 'f'
	Value    Cell = 'v'
)

// String convert Cell to string
func (c Cell) String() string {
	return string(c)
}

// Valid returns true of valid Cell
func (c Cell) Valid() bool {
	switch c {
	case Function:
		return true
	case Value:
		return true
	}

	return false
}

// Resource represents resource type, dir or sheet
type Resource string

// Different types of share
const (
	Dir   Resource = "dir"
	Sheet Resource = "sheet"
)

// Valid returns true for valid resource type, else returns false
func (r Resource) Valid() bool {
	switch r {
	case Dir:
		return true
	case Sheet:
		return true
	}

	return false
}
