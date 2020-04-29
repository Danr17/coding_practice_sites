package lengthstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output int
	}{
		{"abc", "abcabcbb", 3},
		{"b", "bbbbb", 1},
		{"wke", "pwwkew", 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if tc.output != result {
				t.Errorf("expected result to be %v; got %v", tc.output, result)

			}

		})
	}
}
