// In Go, the entry point to a program has to be a function called main within a package main.
package main

import (
	"fmt"
	"os"
)

// we prefix the function name with the package, e.g., fmt.Println
// Go is strict about importing packages. It will not compile if you import a package but don’t use it
func main() {
	var a int = 5 // we can declare variables with var keyword
	//or we can use shorthand declaration
	b := "its not over" //declared varaibles have to be used otherwise compile error
	//It’s important that you remember that := is used to declare the variable as well as assign a value to it. Why? Becausea variable can’t be declared twice (not in the same scope anyway)
	println("LESGOO we learning GO!!")
	if len(os.Args) != 2 { //checking command line arguments
		os.Exit(1) //exit if no argument is passed
	}
	fmt.Println("its over", os.Args[1])  //printing command line argument
	fmt.Printf("its over %d times\n", a) //formatted print println with format specifier
	fmt.Printf("well %s\n", b)           //formatted print with string format specifier
	//CALLING FUNCTIONS

	// 1. Calling function with no return value
	fmt.Println("\n--- Calling log() ---")
	log("This message is printed by the log function!")
	log("Functions make code reusable and organized")

	// 2. Calling function with one return value
	fmt.Println("\n--- Calling add() ---")
	result := add(10, 20)
	fmt.Printf("add(10, 20) = %d\n", result)

	sum := add(a, 15)
	fmt.Printf("add(%d, 15) = %d\n", a, sum)

	// 3. Calling function with two return values
	fmt.Println("\n--- Calling power() ---")
	powerLevel, isOverThousand := power("Goku")
	fmt.Printf("Power level: %d, Over 9000: %v\n", powerLevel, isOverThousand)

	// Ignoring one return value with _
	_, exists := power("Vegeta")
	fmt.Printf("Does Vegeta exist: %v\n", exists)

	// 4. Demonstrating discarding values
	fmt.Println(" Demonstrating Discard Values")
	demonstrateDiscardValue()
	demonstrateKeepOnlyFirst()

	// 5. Demonstrating named return values
	fmt.Println("\nDemonstrating Named Return Values")
	quotient, remainder := divideWithNames(17, 5)
	fmt.Printf("divideWithNames(17, 5) = quotient: %d, remainder: %d\n", quotient, remainder)

	// 6. Demonstrating all concepts together
	demonstrateAllConcepts()
}

// Function with no return value
func log(message string) {
	fmt.Println(message)
}

// Function with one return value
func add(a int, b int) int {
	return a + b
}

// Function with two return values
func power(name string) (int, bool) {
	return 9000, true
}

//EXPLANATION: Multiple Return Values
// Being able to return multiple values is something you'll use often.
// This is useful when a function needs to return both a result AND an error status.
//
// Example: A file reader might return (data, error) - the file content AND whether it succeeded.

// EXPLANATION: Using _ to Discard Values
// Sometimes you only care about one of the return values from a function.
// In those cases, you assign the other values to _ (the blank identifier).
//
// Example 1: We only care about the 'exists' value, not the 'powerLevel'
func demonstrateDiscardValue() {
	_, exists := power("Vegeta")
	fmt.Printf("Example 1 - Ignoring power level: Does Vegeta exist? %v\n", exists)
}

// Example 2: We only care about the power level, not the boolean
func demonstrateKeepOnlyFirst() {
	powerLevel, _ := power("Superman")
	fmt.Printf("Example 2 - Ignoring existence: Power level = %d\n", powerLevel)
}

//EXPLANATION: Named Return Values
// Named return values let you name the return values, making code more readable.
// They are less common but useful for documentation.

// Example without named returns:
// func divide(a, b int) (int, int) {
//     return a / b, a % b
// }

// Example WITH named returns (more readable):
func divideWithNames(a, b int) (quotient int, remainder int) {
	return a / b, a % b
}

//EXPLANATION: Parameter Declaration Shorthand
// When multiple parameters have the same type, you can use shorthand syntax.
//
// VERBOSE (old way):
// func oldAdd(a int, b int) int { return a + b }
//
// SHORTHAND (modern way):
// func add(a, b int) int { return a + b }

// FULL EXAMPLE IN ACTION
func demonstrateAllConcepts() {
	fmt.Println("\nFULL EXAMPLE: All Concepts Together ")

	// Using named return values
	q, r := divideWithNames(17, 5)
	fmt.Printf("17 divided by 5 = quotient: %d, remainder: %d\n", q, r)

	// Using shorthand parameter declaration: add(a, b int)
	total := add(100, 250)
	fmt.Printf("add(100, 250) = %d\n", total)

	// Discarding a return value with _
	_, isReal := power("Naruto")
	fmt.Printf("Is Naruto real? %v (ignoring power level)\n", isReal)
}

//Being able to return multiple values is something you’ll use often. You’ll also frequently use _ to discard a value. Named
//return values and the slightly less verbose parameter declaration aren’t that common. Still, you’ll run into all of these
//sooner than later so it’s important to know about them.
