package leetcode

import (
	"testing"
)

type test0003 struct {
	question string
	answer   int
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	cases := []test0003{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, data := range cases {
		result := lengthOfLongestSubstring(data.question)
		if result != data.answer {
			t.Errorf("input:%v\toutput:%v\texcept:%v\n", data.question, result, data.answer)
		}
	}
}
