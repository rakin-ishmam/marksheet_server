package user

// Name represents system's user nam
type Name string

// Valid validate name
func (n Name) Valid() bool {
	return ok(n.String())
}

func (n Name) String() string {
	return string(n)
}

func ok(name string) bool {
	return okLen(name) && okStr(name) && !(name == GlobalUser().String())
}

func okLen(str string) bool {
	if l := len(str); l < 5 || l > 255 {
		return false
	}

	return true
}

func okStr(str string) bool {
	if str == GlobalUser().String() {
		return false
	}

	for _, c := range str {
		if !okChar(c) {
			return false
		}
	}

	return true
}

func okChar(c int32) bool {
	if c < 'a' || c > 'z' {
		return false
	}

	return true
}

// GlobalUser returns all users
func GlobalUser() Name {
	return "global"
}
