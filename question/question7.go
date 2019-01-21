package question

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	示例 1:

	输入: 123
	输出: 321
	 示例 2:

	输入: -123
	输出: -321
	示例 3:

	输入: 120
	输出: 21
	注意:

	假设我们的环境只能存储得下 32 位的有符号整数，
	则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
 */

func Test7() {
 	result := reverse(1534236469)
	fmt.Println(result)
}

func reverse(x int) int {
	t := strconv.Itoa(x)
	var buffer bytes.Buffer
	minus := false
	tag1 := true
	for i := len(t) -1; i >= 0; i-- {
		if i == 0 && t[i] == 45 {
			minus = true
 			break
		}
		if t[i] == 48 {
			if !tag1 {
				buffer.WriteString(string(t[i]))
			}
		} else {
			tag1 = false
		}

		if t[i] != 48 {
			buffer.WriteString(string(t[i]))
		}

	}
	res := buffer.String()
	if minus {
		res  = "-" + res
	}
	result, _ := strconv.Atoi(res)
	finalResult := int32(result)
	if result != int(finalResult) {
		return 0
	} else {
		return result
	}
 }
