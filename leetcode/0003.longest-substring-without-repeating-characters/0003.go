package leetcode

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	charSet := make(map[int32]int)

	curLen := 0
	for idx, ch := range s {
		lastindex, exist := charSet[ch]
		if !exist || lastindex < 0 {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			for k, v := range charSet {
				if v <= lastindex {
					charSet[k] = -1
				}
			}
			curLen = idx - lastindex
		}
		charSet[ch] = idx
	}
	if curLen > maxLen {
		return curLen
	}
	return maxLen
}
