package leetcode

func twoSum(nums []int, target int) []int {
	numidx := make(map[int]int)
	for idx2, num := range nums {
		idx1, exist := numidx[target-num]
		if exist {
			return []int{idx1, idx2}
		}
		numidx[num] = idx2
	}
	return []int{}
}
