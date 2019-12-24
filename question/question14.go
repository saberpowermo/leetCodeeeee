package question

import (
	"fmt"
	"strings"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

*/

func Test14() {
	test := []string{"dog","racecar","car"}
	result := longestCommonPrefix(test)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	// 既然是最长公共前缀 可以考虑以最短的那个单词开始 为基准，
	// 从第一个字母去找 看别的数组有没有， 有就继续加一个字母 继续找
	// 没有的就换成下一个字母 继续找 直到找到最后1个字母
	if len(strs) == 0 {
		return ""
	}
	shorterString := strs[0]
	index := 0
	// 找出最小的字母
	for index < len(strs) {
		if len(strs[index]) < len(shorterString) {
			shorterString = strs[index]
		}
		index++
	}

	index = 0
	current := ""
	result := ""
 	for index < len(shorterString) {
		fmt.Println(index)
		current += string(shorterString[index])
		fmt.Println(current)
		add := false
		for _, s := range strs {
			if strings.HasPrefix(s, current) {
				add = true
			} else {
				add = false
				break
			}
		}
		if add {
			result = current
			index++
		} else {
			break
		}
	}

	return result
}
