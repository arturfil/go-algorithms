package graphs

// DFS solution
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// dependency table
    table := map[int][]int{}
	current := map[int]bool{} // hash_set
	seen := map[int]bool{} // valid_checked
	for _, prereq := range prerequisites {
		// you are adding an array to the arry in the k, v on the map
		table[prereq[0]] = append(table[prereq[0]], prereq[1])
	}

	var hashCycle func(int) bool
	hashCycle = func(course int) bool { // return true => has cycle
		if current[course] { return true }
		if seen[course] { return false }
		current[course] = true

		for _, dep := range table[course] {
			if hashCycle(dep) { return true }
		}

		delete(current, course)
		seen[course] = true
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hashCycle(i) { return false }
	}

	return true
}

// BFS Solution
func CanFinishBFS(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    degree := make([]int, numCourses)

    for _, prereq := range prerequisites {
        graph[prereq[0]] = append(graph[prereq[1]], prereq[0])
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

    return len(bfs) == numCourses
}
