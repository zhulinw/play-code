package tree

import "github.com/play-code/src/basedata/ele"

type TreeNode struct {
	Val   ele.TreeEle
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val ele.TreeEle) *TreeNode {
	return &TreeNode{Val: val}
}