package access

// Righter maneges the right of the user
type Righter interface {
	Has(Right) bool
	Add(Right) error
	Remove(Right)
	String() string
}
