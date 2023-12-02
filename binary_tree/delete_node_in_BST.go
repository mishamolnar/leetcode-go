package binary_tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil && root.Val == key && root.Left == nil && root.Right == nil {
		return nil
	}
	findNode(root, key)
	return root //doest work for case when root is to delete because old root is returned
}

func findNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	} else if root.Val == key {
		return deleteThisNode(root)
	} else if root.Val > key {
		root.Left = findNode(root.Left, key)
		return root
	} else {
		root.Right = findNode(root.Right, key)
		return root
	}
}

func deleteThisNode(toDelete *TreeNode) *TreeNode {
	if toDelete == nil {
		return nil
	} else if toDelete.Right != nil {
		addNode(toDelete.Right, toDelete.Left)
		toDelete = toDelete.Right
	} else {
		toDelete = toDelete.Left
	}
	return toDelete
}

func addNode(root *TreeNode, toAdd *TreeNode) *TreeNode {
	if root == nil {
		return toAdd
	} else if toAdd == nil {
		return root
	} else if root.Val > toAdd.Val {
		root.Left = addNode(root.Left, toAdd)
	} else {
		root.Right = addNode(root.Right, toAdd)
	}
	return root
}
