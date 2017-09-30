package testformat

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	tt := []struct {
		name     string
		expValue interface{}
		resValue interface{}
		value    error
	}{
		{
			"expected nil result not nil",
			nil,
			"value",
			fmt.Errorf(
				"test is %v, expected value is (%v) but got (%v)",
				"expected nil result not nil",
				nil,
				"value",
			),
		},
		{
			"expected not nil result nil",
			"value",
			nil,
			fmt.Errorf(
				"test is %v, expected value is (%v) but got (%v)",
				"expected not nil result nil",
				"value",
				nil,
			),
		},
		{
			"different value",
			"value",
			"value 2",
			fmt.Errorf(
				"test is %v, expected value is (%v) but got (%v)",
				"different value",
				"value",
				"value 2",
			),
		},
		{
			"all ok",
			"value",
			"value",
			nil,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := New(
				v.name,
				ConvVF(v.expValue),
				ConvVF(v.resValue),
			)
			res := test.Test()

			resv := fmt.Sprintf("%v", res)
			expv := fmt.Sprintf("%v", v.value)

			if resv != expv {
				t.Fatalf("name->%v, expected->%v, but result->%v", v.name, expv, resv)
			}

		})
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			test := NewWithValue(
				v.name,
				v.expValue,
				v.resValue,
			)
			res := test.Test()

			resv := fmt.Sprintf("%v", res)
			expv := fmt.Sprintf("%v", v.value)

			if resv != expv {
				t.Fatalf("name->%v, expected->%v, but result->%v", v.name, expv, resv)
			}

		})
	}
}
