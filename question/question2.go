package question
import "fmt"

/**
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
 */
func Test2() {
	//node13 := &ListNode{
	//	Val: 3,
	//}
	//node12 := &ListNode{
	//	Val:  4,
	//	Next: node13,
	//}
	node11 := &ListNode{
		Val:  5,
		Next: nil,
	}

	//node23 := &ListNode{
	//	Val: 4,
	//}
	//node22 := &ListNode{
	//	Val:  6,
	//	Next: node23,
	//}
	node21 := &ListNode{
		Val:  5,
		Next: nil,
	}

	result := addTwoNumbers(node11, node21)
	for {
		if result == nil {
			return
		}
		fmt.Print(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	currentStep := result
	for {
		if currentStep == nil {
			break
		}
		// 生成下一步
		var l1Int, l2Int int
		if l1 != nil {
			l1Int = l1.Val
		}
		if l2 != nil {
			l2Int = l2.Val
		}
		intResult := l1Int + l2Int + currentStep.Val

		resultNext := &ListNode{
			Val:  0,
			Next: nil,
		}
		if intResult >= 10 {
			currentStep.Val = intResult % 10
			resultNext.Val = 1
		} else {
			currentStep.Val = intResult
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil && resultNext.Val == 0 {
			currentStep.Next = nil
			currentStep = nil
		} else {
			currentStep.Next = resultNext
			currentStep = resultNext
		}
	}

	return result
}