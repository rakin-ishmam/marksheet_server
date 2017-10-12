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
	name     []string
	expected []ValueFunc
	result   []ValueFunc
}

// Add append new test
func (t *Test) Add(name string, expected, result interface{}) {
	t.name = append(t.name, name)
	t.expected = append(t.expected, ConvVF(expected))
	t.result = append(t.result, ConvVF(result))
}

// Test compares that is expected value and result equal
func (t Test) Test() error {

	for i := range t.expected {
		if err := testOne(t.name[i], t.expected[i], t.result[i]); err != nil {
			return err
		}
	}

	return nil
}

func testOne(name string, exp, res ValueFunc) error {
	expv := fmt.Sprintf("value=%v", exp())
	resv := fmt.Sprintf("value=%v", res())

	if expv != resv {
		return fmt.Errorf(
			"test is %v, expected value is (%v) but got (%v)",
			name,
			exp(),
			res(),
		)
	}

	return nil
}

// New returns Test instance
func New(name string, expected, result ValueFunc) Test {
	return Test{
		[]string{name},
		[]ValueFunc{expected},
		[]ValueFunc{result},
	}
}

// NewWithValue returns Test instance
func NewWithValue(name string, expected, result interface{}) Test {
	return Test{
		name:     []string{name},
		expected: []ValueFunc{ConvVF(expected)},
		result:   []ValueFunc{ConvVF(result)},
	}
}

// NewEmpty return empty Test
func NewEmpty() Test {
	return Test{[]string{}, []ValueFunc{}, []ValueFunc{}}
}
