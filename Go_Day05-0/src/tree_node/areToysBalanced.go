package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

var l, r int

func subcheck(tree *TreeNode, sum *int) {
	if tree.HasToy {
		*sum++
	}
	if tree.Left != nil {
		//fmt.Println(tree.Left.HasToy)
		if tree.Left.HasToy {
			*sum++
			//fmt.Println(sum)
		}
		subcheck(tree.Left, sum)
	} else if tree.Right != nil {
		if tree.Right.HasToy {
			*sum++
			//fmt.Println(sum)
		}
		subcheck(tree.Right, sum)
	}
}
func check(tree *TreeNode) {
	fmt.Println(tree.Left.HasToy)
	subcheck(tree.Left, &l)
	subcheck(tree.Right, &r)
}
func main() {
	tree := TreeNode{
		false, &TreeNode{true, nil, nil}, &TreeNode{true, nil, nil}}
	check(&tree)
	fmt.Println(l, r)
}
