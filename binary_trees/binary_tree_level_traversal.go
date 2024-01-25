package binarytrees

// levelOrder will return the nodes of the trees
// level by level using a bfs algorithm
func levelOrder(root *TreeNode) [][]int {
    // if the root is nil return empty 2d array
    if root == nil { return [][]int{} } 
    res := [][]int{}

    // create your queue with the first value
    q := []*TreeNode { root } 

    for len(q) > 0 {
        qLen := len(q)
        level := []int{}

        for i := 0; i < qLen; i++ {
            node := q[0]
            level = append(level, node.Val)
            q = q[1:] // popping 

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
