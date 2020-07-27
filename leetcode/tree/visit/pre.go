package visit

import (
	"fmt"

	"github.com/9999-dollor/algorithm/define"
)

/***
tree 的遍历都是 left -right ,根据node

栈S;
p= root;
while(p || S不空){
    while(p){
        访问p节点；
        p的右子树入S;
        p = p的左子树;
    }
    p = S栈顶弹出;
}

作者：jason-2
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/die-dai-fa-by-jason-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


***/
func PreOrder(root *define.TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var stack []*define.TreeNode
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			fmt.Printf("%d ", cur.Val)
			stack = append(stack, cur.Right)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func PreorderRecursive(root *define.TreeNode) {
	if root == nil {
		return
	}

	PreorderRecursive(root.Left)
	PreorderRecursive(root.Right)
}
