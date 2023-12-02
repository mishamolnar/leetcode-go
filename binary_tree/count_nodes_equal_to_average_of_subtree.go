package binary_tree

func averageOfSubtree(root *TreeNode) int {
	_, _, res := recursiveIter(root)
	return res
}

func recursiveIter(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	} else {
		cL, sL, rL := recursiveIter(node.Left)
		cR, sR, rR := recursiveIter(node.Right)
		count := 1 + cL + cR
		sum := node.Val + sL + sR
		res := 0
		if sum/count == node.Val {
			res++
		}
		return count, sum, res + rR + rL
	}
}
