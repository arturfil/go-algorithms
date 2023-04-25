package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
    file, err := os.Open("data.json")
    if err != nil { fmt.Println(err) }

    byteVal, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteVal, &problems)
    
    for i := 0; i < len(problems.Problems); i++ {
        if (problems.Problems[i].Done == false) {
            fmt.Println("Problem title ", problems.Problems[i])
        }
    }

    defer file.Close()
}
