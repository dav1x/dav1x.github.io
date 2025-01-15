package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//students := []string{"Dent, Arthur", "MacMillan, Tricia", "Prefect, Ford"}
	//sScores := []int{87, 96, 64}

	type score struct {
		name  string
		score int
	}
	/*
		scores := map[string]int{
			students[0]: 87,
			students[1]: 96,
			students[2]: 64,
		}
	*/
	scores := []score{
		score{name: "Dent, Arthur", score: 87},
		{name: "MacMillan, Tricia", score: 96},
		{name: "Prefect, Ford", score: 64},
	}
	fmt.Println("Select score to print (1-3):")
	var option string
	fmt.Scanln(&option)

	var index int
	//strconv package in prod

	if option == "1" {
		index = 0
	} else if option == "2" {
		index = 1
	} else if option == "3" {
		index = 2
	} else {
		//err
		fmt.Println("You selected an unknown option: " + option + ", defaulting to 0")
		index = 0
	}

	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("You selected an unknown option: " + option + ", defaulting to 0")
		index = 0
	}

	fmt.Println("You selected option " + strconv.Itoa(index))
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	/*
		fmt.Println(students[0], scores[students[0]])
		fmt.Println(students[1], scores[students[1]])
		fmt.Println(students[2], scores[students[2]])
	*/
	fmt.Println(scores[index].name, scores[index].score)
	//fmt.Println(scores[1].name, scores[1].score)
	//fmt.Println(scores[2].name, scores[2].score)

	fmt.Println(strings.Repeat("*", 14))
	fmt.Println("Printing all scores")
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i].name, scores[i].score)
		//fmt.Println(students[i], sScores[i])
	}
}

