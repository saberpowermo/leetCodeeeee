package question

import "fmt"

/**
在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test2021427_1() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	result := checkStraightLine(coordinates)
	fmt.Println(result)
}

func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)
	xl := float64(0)
	for index, value := range coordinates {
		if index == l-1 {
			return true
		}
		currentX := value[0]
		currentY := value[1]
		nextX := coordinates[index+1][0]
		nextY := coordinates[index+1][1]
		cXl := float64(0)
		if nextY > currentY {
			cXl = (float64(nextY) - float64(currentY)) / (float64(nextX) - float64(currentX))
		} else {
			cXl = (float64(currentY) - float64(nextY)) / (float64(currentX) - float64(nextX))
		}
 		if xl == 0 {
			xl = cXl
		} else {
			if xl != cXl{
				return false
			}
		}

	}
	return true
}
