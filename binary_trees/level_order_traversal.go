package binarytrees

func levelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{}}
    var res [][]int

    // deque
    q := []*TreeNode{root}
    for len(q) > 0 { 
        qLen := len(q)
        var level []int
        for i := 0; i < qLen; i++ {
            node := q[0] 
            level = append(level, node.Val)
            q = q[1:] // poping

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        res = append(res, level)
    }
    return res
}
