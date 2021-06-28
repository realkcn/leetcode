package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var index1, index2, curIdx int
	var odd, found bool
	var result float64
	midIdx := (len(nums1) + len(nums2)) / 2
	if (len(nums1) + len(nums2)) == midIdx*2 {
		odd = false
		midIdx--
	} else {
		odd = true
	}
	var remain []int
	for {
		if index1 >= len(nums1) {
			remain = nums2[index2:]
			break
		}
		if index2 >= len(nums2) {
			remain = nums1[index1:]
			break
		}
		var curData int
		if nums1[index1] < nums2[index2] {
			curData = nums1[index1]
			index1++
		} else {
			curData = nums2[index2]
			index2++
		}
		if curIdx == midIdx {
			if !odd && found {
				return (result + float64(curData)) / 2
			}
			if !odd {
				found = true
				result = float64(curData)
				midIdx++
			} else {
				return float64(curData)
			}
		}
		curIdx++
	}
	if !odd {
		if found {
			return (result + float64(remain[0])) / 2
		}
		return float64(remain[midIdx-curIdx]+remain[midIdx-curIdx+1]) / 2
	} else {
		return float64(remain[midIdx-curIdx])
	}
}
