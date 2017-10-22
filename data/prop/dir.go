package prop

import (
	"github.com/rakin-ishmam/marksheet_server/config"
)

// DirName represents director name
type DirName string

// Valid returns true for valid dir name, else returns false
func (d DirName) Valid() bool {
	return d.okLen() && d.okChar()
}

func (d DirName) okLen() bool {
	if len := len(d); len < config.MinDirName || len > config.MaxDirName {
		return false
	}

	return true
}

func (d DirName) okChar() bool {
	for _, c := range d {
		if !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
			return false
		}
	}

	return true
}

// NewDirName converts string to DirName
func NewDirName(name string) DirName {
	return DirName(name)
}
