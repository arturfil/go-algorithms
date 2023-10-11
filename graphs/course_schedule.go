package graphs

// DFS solution
func CanFinish(numCourses int, prerequisites [][]int) bool {
    table := map[int][]int{}
    for _, req := range prerequisites {
        table[req[0]] = append(table[req[0]], req[1])
    }
    visited := map[int]struct{}{}

    for i := 0; i < numCourses; i++ {
        if !checkPrereqs(i, visited, table) {
            return false
        }
    }
    return true
}

func checkPrereqs(course int, visited map[int]struct{}, table map[int][]int) bool {
    if len(table[course]) == 0 {
        return true
    }
    if _, yes := visited[course]; yes {
        return false
    }
    visited[course] = struct{}{}
    for _, p := range table[course] {
        if !checkPrereqs(p, visited, table) {
            return false
        }
    }
    delete(visited, course)
    table[course] = nil
    return true
}

// BFS Solution
func CanFinishBFS(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    degree := make([]int, numCourses)

    for _, prereq := range prerequisites {
        graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
        degree[prereq[0]] += 1
    }

    bfs := make([]int, 0)
    for course, d := range degree {
        if d == 0 {
            bfs = append(bfs, course)
        }
    }

    for i := 0; i < len(bfs); i++ {
        course := bfs[i]
        for _, j := range graph[course] {
            degree[j] -= 1
            if degree[j] == 0 {
                bfs = append(bfs, j)
            }
        }
    }

    if len(bfs) == numCourses { return true }
    return false
}
