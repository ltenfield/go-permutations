package main

import (
	"fmt"
)

// original code to permutate letters was taken from stack overflow with playground examples
// https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
// http://play.golang.org/p/0bWDCibSUJ
// will adapt for integer array in sorted order
func GenerateCombinations(anArray []int, length int) <-chan []int {
	c := make(chan []int)
	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan []int) {
		defer close(c)                         // Once the iteration function is finished, we close the channel
		AddNumber(c, []int{}, anArray, length) // We start by feeding it an empty combination
	}(c)
	return c // Return the channel to the calling function
}

// AddNumber adds an integer to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddNumber once again
// to add yet another integer to the new combination in case length allows it
func AddNumber(c chan []int, combo []int, anArray []int, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length < 1 {
		return
	}
	var newCombo []int
	for _, anInteger := range anArray {
		newCombo = append(combo, anInteger)
		c <- newCombo // send a combination result via channel
		AddNumber(c, newCombo, anArray, length-1)
	}
}

func Sum(anArray []int) int {
	result := 0
	for _, v := range anArray {
		result = result + v
	}
	return result
}

// goal is group combinations that results in the same sum
func main() {
	result := make(map[int][][]int)
	// Cool feature of golang: here we range a channel output and when channel closes, range ends and for loop is completed
	for combination := range GenerateCombinations([]int{2, 3, 5, 7, 9, 10}, 2) {
		if len(combination) < 2 {
			continue
		}
		sumOfArrayElements := Sum(combination)
		listOfCombinationsForSum := result[sumOfArrayElements]                 // sumOfArrayElements becomes a map key
		newListOfCombinations := append(listOfCombinationsForSum, combination) // add a new combination to list of combinations
		result[sumOfArrayElements] = newListOfCombinations                     // replace old list of combinations with new list of combinations within map
		fmt.Println(combination)                                               // list a processed combination
	} // END range loop
	// we just print the map with nested arrays
	// Cool golang feature: golang provides string debug output for native types like maps and slices
	// Actually dealing with slices is much easier then using list.List for example as a result container
	// for example you can not range over a list.List
	// Why are lists used infrequently in Go? -> https://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go
	fmt.Println("Done!", result)
}
