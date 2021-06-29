package leetcode

import (
	"testing"
)

type test0839 struct {
	question []string
	answer   int
}

func Test_numSimilarGroups(t *testing.T) {
	cases := []test0839{
		{[]string{"jvhpg", "jhvpg", "hpvgj", "hvpgj", "vhgjp"}, 3},
		{[]string{"nmiwx", "mniwx", "wminx", "mnixw", "xnmwi"}, 2},
		{[]string{"blw", "bwl", "wlb"}, 1},
		{[]string{"tars", "rats", "arts", "star"}, 2},
		{[]string{"omv", "ovm"}, 1},
	}

	for _, data := range cases {
		result := numSimilarGroups(data.question)
		if result != data.answer {
			t.Errorf("input:%v\toutput:%v\texcept:%v\n", data.question, result, data.answer)
		}
	}
}
