package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// ARRAYS
	/*
		Arrays are fixed size and cannot grow.
		You must specify the size when declaring.
	*/

	// Fixed size array
	var scores [10]int///the following line says that the array can hold 10 integers starting from index 0 to 9
	scores[0] = 339

	// Initialize array with values
	scores = [10]int{9001, 9333, 212, 33}

	// Using len and range to iterate
	for index, value := range scores {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
//in Go we rarely use arrays directly; instead, we use slices, which are more flexible and powerful.
	
	// SLICES - Lightweight wrappers over arrays
	// 
	/*
		Unlike arrays, slices don't require a size in declaration.
		Slices are lightweight and can grow dynamically.
	*/

	// Create slice with values
	scoresSlice := []int{1, 4, 293, 4, 9}
	fmt.Println("Slice:", scoresSlice)

	// Create slice with make - length 10, capacity 10
	scoresWithLength := make([]int, 10)
	fmt.Println("Slice with length 10:", scoresWithLength)

	// Create slice with length 0, capacity 10
	scoresZeroLength := make([]int, 0, 10)
	fmt.Println("Slice length 0, capacity 10:", scoresZeroLength)

	// Reslice to access higher indexes
	scoresZeroLength = scoresZeroLength[0:8]
	scoresZeroLength[7] = 9033
	fmt.Println("After reslicing and setting index 7:", scoresZeroLength)

	// Append - expands slice dynamically
	emptySlice := make([]int, 0, 10)
	emptySlice = append(emptySlice, 5)
	fmt.Println("After append:", emptySlice)

	// Demonstrating slice capacity growth with 2x algorithm 
	// (not guaranteed but common implementation)
	scoresGrow := make([]int, 0, 5)//initial capacity is 5
	c := cap(scoresGrow)//get current capacity
	fmt.Println("Initial capacity:", c)//prints 5
	for i := 0; i < 25; i++ {//append 25 elements
		scoresGrow = append(scoresGrow, i)//append i to slice
		if cap(scoresGrow) != c {//check if capacity has changed
			c = cap(scoresGrow)	//update capacity
			fmt.Println("Capacity expanded to:", c)
		}
	}

	// Append to slice with existing values
	scoresWithValues := make([]int, 5)
	scoresWithValues = append(scoresWithValues, 9332)
	fmt.Println("Append to slice with 5 values:", scoresWithValues) // [0 0 0 0 0 9332]

	// Four common ways to initialize slices
	names := []string{"leto", "jessica", "paul"}//hefirstoneshouldn’t need much of an explanation.You use thiswhenyouknowthevalues thatyouwant inthearrayaheadoftime
	checks := make([]bool, 10) //The second one is useful when you’ll be writing into specific indexes of a slice.
	var namesNil []string
	scoresCapacity := make([]int, 0, 20)

	fmt.Println("Names:", names)
	fmt.Println("Checks:", checks)
	fmt.Println("Names nil:", namesNil)
	fmt.Println("Scores capacity:", scoresCapacity)

	// Slice as a window (not a copy like in Python/Ruby)
	scoresList := []int{1, 2, 3, 4, 5}
	slice := scoresList[2:4] // from index 2 to 4 (not including 4)
	slice[0] = 999
	fmt.Println("Original after slice modification:", scoresList) // [1 2 999 4 5]

	// Slice syntax variations
	haystack := "the spice must flow"
	fmt.Println("Index of space after 5 chars:", strings.Index(haystack[5:], " "))

	// Remove all but last element
	scoreRemove := []int{1, 2, 3, 4, 5}
	scoreRemove = scoreRemove[:len(scoreRemove)-1]
	fmt.Println("All but last:", scoreRemove)

	// Remove at index function
	scoreToRemove := []int{1, 2, 3, 4, 5}
	scoreToRemove = removeAtIndex(scoreToRemove, 2)
	fmt.Println("After removing index 2:", scoreToRemove)

	// Copy function
	scores100 := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores100[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores100)
	worst := make([]int, 5)
	copy(worst, scores100[:5])
	fmt.Println("Worst 5 scores:", worst)

	
	// MAPS - Hash tables / Dictionaries
	/*
		Maps are key-value pairs that grow dynamically.
		Keys and values can be any type.
	*/

	// Create map
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println("Power of vegeta:", power, "exists:", exists) // 0, false

	// Get length and delete
	total := len(lookup)
	fmt.Println("Number of keys:", total)
	delete(lookup, "goku")

	// Map with initial size

	// Map as composite literal
	powerLevels := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
	fmt.Println("Power levels:", powerLevels)

	// Iterate over map (random order each time)
	for key, value := range powerLevels {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

/*
Helper function to remove element at index from slice.
Won't preserve order - swaps last element with target.
*/
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// Swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
