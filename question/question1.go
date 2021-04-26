package question

import (
	"fmt"
)

/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
	示例:
	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
*/
func Test1() {
	nums := []int{11, 3, 5, 6, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result[0])
	fmt.Println(result[1])
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	m := map[int]int{}
	var interpolation int
	for index, value := range nums {
		interpolation = target - value
		i, ok := m[interpolation]
		if ok {
			result[0] = i
			result[1] = index
		} else {
			m[value] = index
		}
	}
	return result
}
