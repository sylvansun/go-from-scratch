package algorithms

func DistributeCoins(root *TreeNode) int {
	ret := 0
	dfs979(root, &ret)
	return ret
}

func dfs979(root *TreeNode, ret *int) int {
	moveleft, moveright := 0, 0
	if root == nil {
		return 0
	}
	if root.Left != nil {
		moveleft = dfs979(root.Left, ret)
	}
	if root.Right != nil {
		moveright = dfs979(root.Right, ret)
	}
	*ret += absInt(moveleft) + absInt(moveright)
	return moveleft + moveright + root.Val - 1
}
