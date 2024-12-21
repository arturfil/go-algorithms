package binarytrees

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil { return false }

    if isSameTree(root, subRoot) { return true }

    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(a *TreeNode, b *TreeNode) bool {
    if a == nil && b == nil { return true }
    if a == nil || b == nil { return false }

    if a.Val != b.Val { return false }

    return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}
