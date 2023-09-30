package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Problem struct {
    Done bool `json:"done"`
    Title string `json:"title"`
}

func GetProblemsFromJsonData() []Problem {
    problems := struct {Problems []Problem}{}

    file, err := os.Open("data.json")
    if err != nil { fmt.Println(err) }

    byteVal, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteVal, &problems)

    return problems.Problems
}

func RandomProblem() {
    var notDone []Problem
    problems := GetProblemsFromJsonData()
        
    for i := 0; i < len(problems); i++ {
        if (problems[i].Done == false) {
            notDone = append(notDone, problems[i])
        }
    }

    rand.Seed(time.Now().UnixNano()) 
    randIdx := rand.Intn(len(notDone) - 0 + 1) + 0
    
    fmt.Println("Problem chosen --->", notDone[randIdx].Title)
}

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
