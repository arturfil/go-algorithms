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

// RandomProblem - gets the data from GetProblemsFromJsonData & from there,
// it choosees a random problem to solve in leetcode
func ChooseRandomProblem() {
	listPtr := flag.Bool("exclude", false, "flag to filter problems by")
	ctgyPtr := flag.String("category", "all", "flag to filter by category")

	// parse values from flags
	flag.Parse()

	var notDone []Problem
	problems := GetProblemsFromJsonData()

	for i := 0; i < len(problems); i++ {
		if problems[i].Done != false {
			continue
		}

		if *ctgyPtr != "all" || *listPtr {

			if *ctgyPtr != "all" && problems[i].Category == *ctgyPtr && !*listPtr {
				notDone = append(notDone, problems[i])

			} else if *listPtr && problems[i].IsFrom75 == *listPtr && *ctgyPtr == "all" {
				notDone = append(notDone, problems[i])

			} else if problems[i].Category == *ctgyPtr && problems[i].IsFrom75 == *listPtr {
				notDone = append(notDone, problems[i])
			}

		} else {
			notDone = append(notDone, problems[i])
		}
	}

	fmt.Println("notDone slice length:", len(notDone))

	// PrintDoneProblems()

	rand.NewSource(time.Now().UnixNano())
	randIdx := rand.Intn(len(notDone))

	fmt.Println("Problem chosen --->", notDone[randIdx].Title)
}

// GetProblemsFromJsonData - gets the data from a json file and returns it into a slice
func GetProblemsFromJsonData() []Problem {

	problems := struct{ Problems []Problem }{}

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	byteVal, _ := io.ReadAll(file)
	json.Unmarshal(byteVal, &problems)

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
