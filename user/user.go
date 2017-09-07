package user

// Name represents system's user nam
type Name string

// Valid validate name
func (n Name) Valid() bool {
	return ok(string(n))
}

func ok(name string) bool {
	return okLen(name) && okChar(name)
}

func okLen(str string) bool {
	if l := len(str); l < 5 || l > 255 {
		return false
	}

	return true
}

func okChar(str string) bool {
	for _, c := range str {
		if c < 'a' || c > 'z' {
			return false
		}
	}

	return true
}
