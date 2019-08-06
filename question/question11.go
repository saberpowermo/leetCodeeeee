package question

import (
	"fmt"
)

/**
11. 盛最多水的容器

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
*/

func Test11() {
	test := []int{1,8,6,2,5,4,8,3,7}
	max := maxArea(test)
	fmt.Println(max)
}

// 如果用On2的时间复杂度去做 就没什么意思了... 如何On的时间复杂度去完成呢?

func maxArea(height []int) int {
	max := 0
	minIndex := 0
	maxIndex := len(height) - 1

	for minIndex < maxIndex {
		// 水桶能承多少水 要看最低的那个板子， 所以两头height 比较 谁低谁收缩.
		// 用两头height 最小的 乘以两头的index差 就是水的面积.
		y := 0
		if height[minIndex] > height[maxIndex] {
			y = height[maxIndex]
		} else {
			y = height[minIndex]
		}
		x := maxIndex - minIndex
		newMax := x * y
		if newMax > max {
			max = newMax
		}

		if height[minIndex] > height[maxIndex] {
			maxIndex--
		} else {
			minIndex++
		}
	}
	return max
}
