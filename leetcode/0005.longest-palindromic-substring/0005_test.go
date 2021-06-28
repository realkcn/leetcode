package leetcode

import "testing"

type test0005 struct {
	question string
	answer   string
}

func Test_longestPalindrome(t *testing.T) {
	cases := []test0005{
		{"aacabdkacaa", "aca"},
		{"aaaaa", "aaaaa"},
		{"aaaa", "aaaa"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}

	for _, data := range cases {
		result := longestPalindrome(data.question)
		if result != data.answer {
			t.Errorf("input:%v\toutput:%v\texcept:%v\n", data.question, result, data.answer)
		}
	}
}
