package binary_tree

import (
	"fmt"
)

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func makeBinaryTreeNode(val int) *BinaryTreeNode {
	var root BinaryTreeNode
	root.val = val
	return &root
}

func displayBinaryTree(root *BinaryTreeNode, dep int) {
	if root == nil {
		return
	}
	displayBinaryTree(root.right, dep+1)
	for i := 0; i < dep; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.val)
	displayBinaryTree(root.left, dep+1)
}

func (root *BinaryTreeNode) insert(val int) {
	if val < root.val {
		if root.left == nil {
			root.left = makeBinaryTreeNode(val)
		} else {
			root.left.insert(val)
		}
	} else {
		if root.right == nil {
			root.right = makeBinaryTreeNode(val)
		} else {
			root.right.insert(val)
		}
	}

}

func main() {
	data := []int{
		1, 3, 2, 8, 5, 4, 6, 6,
	}
	root := makeBinaryTreeNode(data[0])
	for _, x := range data[1:] {
		root.insert(x)
	}
	displayBinaryTree(root, 0)
}
