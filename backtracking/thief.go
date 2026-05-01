package backtracking

// Definition of a binary tree node
type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func rob(root *TreeNode[int]) int {
	if root == nil {
		return -1
	}

	a, b := sumTree(root)

	return maxInt(a, b)
}

func sumTree(root *TreeNode[int]) (int, int) {
	if root == nil {
		return 0, 0
	}

	includeL, excludeL := sumTree(root.Left)
	includeR, excludeR := sumTree(root.Right)

	includeN := excludeL + excludeR + root.Data

	excludeN := maxInt(excludeL, includeL) + maxInt(excludeR, includeR)

	return includeN, excludeN
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
