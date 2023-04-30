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

type Problems struct {
    Problems []Problem `json:"problems"`
}

func ReadJson() {
    var problems Problems
    var notDone []Problem

    file, err := os.Open("data.json")
    if err != nil { fmt.Println(err) }

    byteVal, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteVal, &problems)
    
    for i := 0; i < len(problems.Problems); i++ {
        if (problems.Problems[i].Done == false) {
            notDone = append(notDone, problems.Problems[i])
        }
    }
    rand.Seed(time.Now().UnixNano()) 
    randIdx := rand.Intn(len(notDone) - 0 + 1) + 0
    
    fmt.Println("Problem chosen --->", notDone[randIdx].Title)
}
