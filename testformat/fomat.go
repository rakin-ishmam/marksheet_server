package testformat

import (
	"fmt"
)

// ValueFunc use for to get value
type ValueFunc func() interface{}

// ConvVF takes a value and returns ValueFunc
func ConvVF(v interface{}) ValueFunc {
	return func() interface{} {
		return v
	}
}

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

// New returns Test instance
func New(name string, expected, result ValueFunc) Test {
	return Test{name, expected, result}
}

// NewWithValue returns Test instance
func NewWithValue(name string, expected, result interface{}) Test {
	return Test{
		name:     name,
		expected: ConvVF(expected),
		result:   ConvVF(result),
	}
}
