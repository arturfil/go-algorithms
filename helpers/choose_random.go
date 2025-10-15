package helpers

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type Problem struct {
	Done      bool   `json:"done"`
	Title     string `json:"title"`
	IsFrom75  bool   `json:"isFrom75"`
	Algorithm string `json:"algorithm"`
	Category  string `json:"category"`
}

// shouldIncludeProblem checks if a problem matches the filter criteria
func shouldIncludeProblem(p Problem, category string, from75Only bool) bool {
	// Filter by category if specified
	if category != "all" && p.Category != category {
		return false
	}

	// Filter by isFrom75 if flag is set
	if from75Only && !p.IsFrom75 {
		return false
	}

	return true
}

// ChooseRandomProblem - gets the data from GetProblemsFromJsonData & from there,
// it chooses a random problem to solve in leetcode
func ChooseRandomProblem() {
	from75Ptr := flag.Bool("from75", false, "filter only problems from the Blind 75 list")
	ctgyPtr := flag.String("category", "all", "filter by category (e.g., arrays, strings, graphs)")

	// parse values from flags
	flag.Parse()

	var notDone []Problem
	problems := GetProblemsFromJsonData()

	// Apply filters
	for _, problem := range problems {
		// Skip completed problems
		if problem.Done {
			continue
		}

		// Apply category and from75 filters
		if shouldIncludeProblem(problem, *ctgyPtr, *from75Ptr) {
			notDone = append(notDone, problem)
		}
	}

	fmt.Println("notDone slice length:", len(notDone))

	// Handle case when no problems match the filters
	if len(notDone) == 0 {
		fmt.Println("No problems found matching the specified filters")
		return
	}

	// PrintDoneProblems()

	// Properly seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIdx := r.Intn(len(notDone))

	fmt.Println("Problem chosen --->", notDone[randIdx].Title)
}

// GetProblemsFromJsonData - gets the data from a json file and returns it into a slice
func GetProblemsFromJsonData() []Problem {
	problems := struct{ Problems []Problem }{}

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening data.json:", err)
		return []Problem{}
	}
	defer file.Close()

	byteVal, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading data.json:", err)
		return []Problem{}
	}

	if err := json.Unmarshal(byteVal, &problems); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return []Problem{}
	}

	return problems.Problems
}

// PrintDoneProblems - displyas problems done
func PrintDoneProblems() {
	problems := GetProblemsFromJsonData()
	doneProblems := []Problem{}

	for _, p := range problems {
		if p.Done == true {
			doneProblems = append(doneProblems, p)
			fmt.Println(p)
		}
	}

	fmt.Println("Done problems", len(doneProblems))
}
