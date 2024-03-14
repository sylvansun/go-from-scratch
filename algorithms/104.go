package algorithms

func calculateDepth(root *TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}
	left := calculateDepth(root.Left) + 1
	right := calculateDepth(root.Right) + 1
	ret = maxInt(left, right)
	return ret
}
