package lengthOfLongestSubstring

import "testing"

type lengthOfLongestSubstringTestCase struct {
	inp string
	exp int
}

var lengthOfLongestSubstringTestCases []lengthOfLongestSubstringTestCase = []lengthOfLongestSubstringTestCase{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range lengthOfLongestSubstringTestCases {
		if got := lengthOfLongestSubstring(tc.inp); got != tc.exp {
			t.Errorf("lengthOfLongestSubtring(%v) returned %v, expected %v", tc.inp, got, tc.exp)
		}
	}
}
