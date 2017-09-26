package testformat

import (
	"fmt"
)

// ValueFunc use for to get value
type ValueFunc func() interface{}

// Test comapares the test result
type Test struct {
	name     string
	expected ValueFunc
	result   ValueFunc
}

// Test compares that is expected value and result equal
func (t Test) Test() error {
	expv := fmt.Sprintf("value=%v", t.expected())
	resv := fmt.Sprintf("value=%v", t.result())

	if expv != resv {
		return fmt.Errorf(
			"test is %v, expected value is (%v) but got (%v)",
			t.name,
			t.expected(),
			t.result(),
		)
	}

	return nil
}

// NewTest returns Test instance
func NewTest(name string, expected, result ValueFunc) Test {
	return Test{name, expected, result}
}
