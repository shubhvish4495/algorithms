package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

func returnBst(h, l int, nums []int) *helper.TreeNode {
	if h < l {
		return nil
	}
	if l == h {
		return &helper.TreeNode{
			Val: nums[l],
		}
	}

	mid := (h + l) / 2
	leftTree := returnBst(mid-1, l, nums)
	rightTree := returnBst(h, mid+1, nums)
	return &helper.TreeNode{
		Val:   nums[mid],
		Left:  leftTree,
		Right: rightTree,
	}
}

func sortedArrayToBST(nums []int) *helper.TreeNode {
	return returnBst(len(nums)-1, 0, nums)
}

func sortedArrayToBST2(nums []int) *helper.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &helper.TreeNode{Val: nums[i]}
	if i != 0 {
		root.Left = sortedArrayToBST(nums[:i])
	}
	if i != len(nums)-1 {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	t := sortedArrayToBST2(nums)
	t.Print()

	fmt.Print("\n")

	t2 := sortedArrayToBST(nums)
	t2.Print()
}
