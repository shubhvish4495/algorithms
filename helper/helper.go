package helper

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Print() {
	if t == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Print(fmt.Sprintf("%d ", t.Val))
	t.Left.Print()
	t.Right.Print()
}
