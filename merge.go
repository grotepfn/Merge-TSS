package main

import (
	"encoding/json"
	"flag"
	"log"
	"sort"
	"strconv"
)

var input []string

func main() {

	//input sets
	input := [][2]int{
		{-3, -2},
		{-2, 1},
		{0, 1},
		{4, 5},
		{2, 3},
		{4, 20}}

	flag.Parse()

	var args = flag.Args()

	if len(args) > 0 {

		input = [][2]int{}
	}
	for i := 0; i <= len(args)-1; i++ {
		var parsedInput []int
		err := json.Unmarshal([]byte(args[i]), &parsedInput)
		if err != nil {
			log.Fatalf("bad input %v", err)
		}
		if len(parsedInput) != 2 || parsedInput[1] < parsedInput[0] {
			log.Fatalf("pass array with length of two with first lower bound and second upper bound")
		} else {
			parsedInput := [2]int{parsedInput[0], parsedInput[1]}

			input = append(input, parsedInput)
		}

	}

	Merge(input)
}

//Merges integer arrays with a length of 2 if their ranges overleap
//first input is the lower bound, second input the upper bound
func Merge(list [][2]int) [][2]int {

	sort.SliceStable(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})

	mergedList := make([][2]int, 1)

	mergedList[0] = list[0]

	for i := 0; i < len(list); i++ {
		//new set attached, does not overleap
		if list[i][0] > mergedList[len(mergedList)-1][1] {
			mergedList = append(mergedList, list[i])
			//new set merched, overleaps
		} else if list[i][0] <= mergedList[len(mergedList)-1][1] && list[i][1] > mergedList[len(mergedList)-1][1] {

			mergedList[len(mergedList)-1][1] = list[i][1]
		}
	}

	for i := 0; i < len(mergedList); i++ {
		log.Printf(" [" + strconv.Itoa(mergedList[i][0]) + "," + strconv.Itoa(mergedList[i][1]) + "]")

	}

	return mergedList
}
