package leetcode

func longestPalindrome(s string) string {
	maxLen := int16(1)
	maxStart := int16(0)

	curPalindromeLens := [1024]int16{1}
	palindromeCount := 1

	sameLen := int16(1)
	checkMax := func(len int16, start int16) {
		if len > maxLen {
			maxLen = len
			maxStart = start
		}
	}
	if len(s) > 1 {
		for i := int16(1); i < int16(len(s)); i++ {
			var newCount int
			newCount = 0
			for j := 0; j < palindromeCount; j++ {
				preIdx := i - curPalindromeLens[j] - 1
				if preIdx >= 0 && s[preIdx] == s[i] {
					curPalindromeLens[newCount] = curPalindromeLens[j] + 2
					checkMax(curPalindromeLens[newCount], i-curPalindromeLens[newCount]+1)
					newCount++
				}
			}
			if s[i] == s[i-1] {
				curPalindromeLens[newCount] = sameLen + 1
				sameLen++
				checkMax(curPalindromeLens[newCount], i-curPalindromeLens[newCount]+1)
				newCount++
			} else {
				sameLen = 1
				if i >= 2 && s[i] == s[i-2] {
					curPalindromeLens[newCount] = 3
					checkMax(curPalindromeLens[newCount], i-curPalindromeLens[newCount]+1)
					newCount++
				}
			}
			palindromeCount = newCount
		}
	}
	return s[maxStart : maxStart+maxLen]
}
