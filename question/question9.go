package question

import (
	"fmt"
	"math"
)

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

*/

func Test9() {
	is := isPalindrome(-121)
	fmt.Println("is: ", is)
}

func isPalindrome(x int) bool {
	// 不转成字符串来解决的话 可以 先判断数字是几位
	// 如果是奇数 那就以中间为准 两边扩散去 == 判断
	if x < 0 {
		return false
	}
	l := checkIntLen(x)
	fmt.Println("这个int位数:", l)
	isOdd := l%2 == 1
	if isOdd {
		// 奇数
		mid := l / 2
		start := 0
		end := l
		fmt.Println("mid: ", mid)
		for start < mid && end > mid {

			startInt := x / int(math.Pow10(start)) % 10
			endInt := x / int(math.Pow10(end-1)) % 10
			if startInt != endInt {
				return false
			}
			start++
			end--
		}
	} else {
		end := l / 2
		start := end + 1
		for start > 0 && end < l {
			startInt := x / int(math.Pow10(start-2)) % 10
			endInt := x / int(math.Pow10(end)) % 10
			fmt.Println("start: ", startInt, ", end: ", endInt)
			if startInt != endInt {
				return false
			}
			start--
			end++
		}
	}

	return true
}

func checkIntLen(x int) int {
	i := 0
	for x > 0 {
		x /= 10
		i++
	}
	return i
}
