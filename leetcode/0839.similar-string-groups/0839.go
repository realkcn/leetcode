package leetcode

type strIdList []int
type charMap map[uint8]strIdList

func numSimilarGroups(strs []string) int {
	count := len(strs)
	strLen := len(strs[0])
	allStrMap := [300]charMap{}
	connectMatrix := [300][300]bool{}

	for i := 0; i < strLen; i++ {
		allStrMap[i] = make(charMap)
	}
	for strIdx, curStr := range strs {
		var equalMatrix [300]int

		for i := 0; i < strIdx-1; i++ {
			equalMatrix[i] = 0
		}
		for i := 0; i < strLen; i++ {
			cmap, exist := allStrMap[i][curStr[i]]
			if exist {
				for _, id := range cmap {
					equalMatrix[id]++
				}
				allStrMap[i][curStr[i]] = append(allStrMap[i][curStr[i]], strIdx)
			} else {
				allStrMap[i][curStr[i]] = []int{strIdx}
			}
		}
		for i := 0; i < strIdx; i++ {
			if equalMatrix[i] >= strLen-2 {
				connectMatrix[i][strIdx] = true
				connectMatrix[strIdx][i] = true
			}
		}
	}
	added := 1
	result := 1
	doingList := []int{0}
	var doneFlag [300]bool
	doneFlag[0] = true
	for true {
		for len(doingList) > 0 {
			curIdx := doingList[0]
			doingList = doingList[1:]
			for i := 0; i < count; i++ {
				if i != curIdx && !doneFlag[i] && connectMatrix[i][curIdx] {
					doneFlag[i] = true
					added++
					doingList = append(doingList, i)
				}
			}
		}
		if added != count {
			for i := 0; i < count; i++ {
				if !doneFlag[i] {
					doingList = []int{i}
					doneFlag[i] = true
					added++
					result++
					break
				}
			}
		} else {
			break
		}
	}
	return result
}
