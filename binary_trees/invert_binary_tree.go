package binarytrees

func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        left := &root.Left
        right := &root.Right

        root.Left = invertTree(*right)
        root.Right = invertTree(*left)
    }
    return root
}
