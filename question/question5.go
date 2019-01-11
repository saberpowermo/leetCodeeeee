package question

import (
	"fmt"
	"strings"
)

/*
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	示例 1：

	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。
	示例 2：

	输入: "cbbd"
	输出: "bb"
 */
func Test5() {
	result := longestPalindrome("aaaaa")
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	runes := []rune(s)
	length := len(runes)

	switch length {
	case 1:
		return s
	case 2:
		if runes[0] == runes[1] {
			return s
		} else {
			return string(runes[0])
		}
	default:
		result := ""
		for index, r := range runes {
			if index == 0 || index == length-1 {
				continue
			}
			currentC := string(r)
			pre, back := index-1, index+1
			isCenter := true
			for {
				if isCenter {
					currentCenter := false
 					if pre >= 0 && strings.HasSuffix(currentC, string(runes[pre])) {
						currentC = string(runes[pre]) + currentC
						pre--
						currentCenter = true
					}
					if back <= length-1 && strings.HasPrefix(currentC, string(runes[back])) {
						currentC = currentC + string(runes[back])
						back++
						currentCenter = true
					}
					isCenter = currentCenter
					fmt.Println("center: " + currentC)
				} else {
					if pre >= 0 && back <= length-1 && runes[pre] == runes[back] {
						currentC = string(runes[pre]) + currentC + string(runes[back])
						pre--
						back++
					} else {
						if len(currentC) > len(result) {
							result = currentC
						}
						break
					}
				}

				if len(currentC) > len(result) {
					result = currentC
				}

				fmt.Println("result : " + result)
				if pre < 0 && back >= length {
					break
				}
			}
			fmt.Println()
		}
		return result
	}
}
