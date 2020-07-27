package visit

import (
	"fmt"

	"github.com/9999-dollor/algorithm/leetcode/tree"
)

func Inorder(root *tree.TreeNode) {
	if root == nil {
		return
	}
	var stack []*tree.TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			/* place pointer to a tree node on
			   the stack before traversing
			  the node's left subtree */
			stack = append(stack, cur)
			cur = cur.Left
		}
		/* Current must be NULL at this point */
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d  ", cur.Val)
		/* we have visited the node and its
		   left subtree.  Now, it's right
		   subtree's turn */
		cur = cur.Right
	}
}

func InOrderRecursive(root *tree.TreeNode) {
	if root == nil {
		return
	}
	InOrderRecursive(root.Left)
	fmt.Printf("%d  ", root.Val)
	InOrderRecursive(root.Right)
}
