package binarytrees

func IsValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(root, left, right *TreeNode) bool {
	if  root == nil { return true }
	if left != nil && left.Val >= root.Val { return false }
	if right != nil && root.Val >= right.Val { return false }
	return validate(root.Left, left, root) && validate(root.Right, root, right)
}
