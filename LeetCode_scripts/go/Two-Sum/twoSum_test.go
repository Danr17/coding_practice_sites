package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		name   string
		slice  []int
		target int
		result []int
	}{
		{"primul", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"urmatorul", []int{5, 7, 3, 15}, 8, []int{0, 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.slice, tc.target)
			if !reflect.DeepEqual(tc.result, result) {
				t.Errorf("expected result to be %v; got %v", tc.result, result)
			}
		})
	}
}
