package binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := math.MinInt32
		newQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Val > curr {
				curr = node.Val
			}

			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
		}
		queue = newQueue
		res = append(res, curr)
	}
	return res
}
