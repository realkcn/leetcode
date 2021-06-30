package leetcode

import (
	"testing"
)

type test0010 struct {
	question question0010
	answer   bool
}
type question0010 struct {
	str     string
	pattern string
}

func Test_isMatch(t *testing.T) {
	cases := []test0010{
		{question0010{"aa", "a"}, false},
		{question0010{"aa", "a*"}, true},
		{question0010{"aa", "a*a*"}, true},
		{question0010{"ab", ".*"}, true},
		{question0010{"ab", ".*.*"}, true},
		{question0010{"aab", "c*a*b"}, true},
		{question0010{"mississippi", "mis*is*p*."}, false},
	}

	for _, data := range cases {
		result := isMatch(data.question.str, data.question.pattern)
		if result != data.answer {
			t.Errorf("input:%v\toutput:%v\texcept:%v\n", data.question, result, data.answer)
		}
	}
}
