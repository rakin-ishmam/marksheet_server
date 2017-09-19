package errs

// Kind defines error kind
type Kind uint8

// Define list of error kind
const (
	Invalid Kind = iota
	Exist
	NotExist
	Unauthorised
)

// String converts Kind to string
func (k Kind) String() string {
	switch k {
	case Invalid:
		return "invalid"
	case Exist:
		return "exist"
	case NotExist:
		return "not exit"
	case Unauthorised:
		return "unauthorised"
	}

	return "unknown"
}
