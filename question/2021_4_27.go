package question

import "fmt"

func Test2021427() {
	r6 := &TreeNode{
		Val:   18,
		Left:  nil,
		Right: nil,
	}
	r5 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	r4 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	r3 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: r6,
	}
	r2 := &TreeNode{
		Val:   5,
		Left:  r4,
		Right: r5,
	}
	r1 := &TreeNode{
		Val:   10,
		Left:  r2,
		Right: r3,
	}
	result := rangeSumBST(r1, 7, 15)
	fmt.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
