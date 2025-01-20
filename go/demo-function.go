## function demo

package main

import (
	"fmt"
	"strconv"
	"strings"
)
	
type score struct {
  name  string
  score int
}

func main() {

	scores := []score{}
	shouldContinue := true

	for shouldContinue {

		fmt.Println("1) Enter a score")
		fmt.Println("2) Print Report")
		fmt.Println("q) Quit")
		fmt.Println("")
		fmt.Println("Please select an option")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			scores = append(scores, addStudent())
		case "2":
      			printReport(scores) 
		default:
			//exit loop
			shouldContinue = false
		}

	}
}
func addStudent() score {

  fmt.Println("Enter a student name and score")
  var name, rawScore string
  fmt.Scanln(&name, &rawScore)
  s, _ := strconv.Atoi(rawScore)
  return score{name: name, score: s}

}

func printReport(scores []score) {

  fmt.Println("Student Scores")
  fmt.Println(strings.Repeat("-", 14))
  //fmt.Println(scores)
  for _, s := range scores {
    fmt.Println(s.name, s.score)
  }
  fmt.Println("")
} 
