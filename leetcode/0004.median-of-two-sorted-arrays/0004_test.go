package leetcode

import (
	"testing"
)

type test0003 struct {
	question question0003
	answer   float64
}

type question0003 struct {
	array1 []int
	array2 []int
}

func Test_findMedianSortedArrays(t *testing.T) {
	cases := []test0003{
		{question: question0003{[]int{}, []int{2, 3}}, answer: 2.5},
		{question: question0003{[]int{1, 3}, []int{2}}, answer: 2.0},
		{question: question0003{[]int{1, 2}, []int{3, 4}}, answer: 2.5},
		{question: question0003{[]int{0, 0}, []int{0, 0}}, answer: 0},
		{question: question0003{[]int{}, []int{1}}, answer: 1.0},
		{question: question0003{[]int{2}, []int{}}, answer: 2.0},
	}

	for _, data := range cases {
		result := findMedianSortedArrays(data.question.array1, data.question.array2)
		if result != data.answer {
			t.Errorf("input:%v\toutput:%v\texcept:%v\n", data.question, result, data.answer)
		}
	}
}
