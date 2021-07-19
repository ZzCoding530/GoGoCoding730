package finalShoot

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil { //这个三个if可以写在一起，不过,为了思路更清晰明显，分开写一下
		return root
	}
	if root == p {
		return root
	}
	if root == q {
		return root
	}

	//lowestCommonAncestor 这个函数的意义是，给你root 和 p , q 两个节点，返回其最近公共祖先节点，要么找到了，要么在root左边子树里，要么在root右边子树里
	left := lowestCommonAncestor(root.Left, p, q)   //左子树扔进去，返回一个值	要么返回nil，说明【左边子树里】没找到，要么返回找到的值
	right := lowestCommonAncestor(root.Right, p, q) //右子树扔进去，返回一个值	要么返回nil，说明【右边子树里】没找到，要么返回找到的值

	if left != nil && right != nil { //左右都不是nil，说明root就是要找的
		return root
	}

	if left == nil { //如果左边是nil，说明qp两个节点肯定俩都在右边，那他俩的公共祖先也在右边啊
		return right
	}

	return left //反之在左边
}
