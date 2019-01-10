package question

import (
	"fmt"
	"sort"
)

/*
	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
	你可以假设 nums1 和 nums2 不会同时为空。
	示例 1:

	nums1 = [1, 3]
	nums2 = [2]

	则中位数是 2.0
	示例 2:

	nums1 = [1, 2]
	nums2 = [3, 4]

	则中位数是 (2 + 3)/2 = 2.5
 */
func Test4() {
	num1 := []int{1, 3}
	num2 := []int{2}
	result := findMedianSortedArrays(num1, num2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	var resultArray []int

	resultArray = append(nums1, nums2...)
	sort.Ints(resultArray)
	length := len(resultArray)
	if length%2 == 0 {
		i1 := (length / 2) - 1
		i2 := length / 2
		result = float64(resultArray[i1]+resultArray[i2]) / 2
	} else {
		result = float64(resultArray[length/2])
	}

	return result
}
