package main

import "fmt"

// Control Structures:
// Go provides various control structures like if, for, switch, and goto. Let's cover each with examples:
func main() {
	// If Statement:
	// The if statement allows conditional execution of code blocks.
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// For Statement:
	// The for statement is used for looping in Go. There are different variations like the complete for, the condition-only for, and the infinite for.
	// Complete for statement
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Condition-only for statement (like while loop)
	sum := 0
	for sum < 10 {
		sum += 2
		fmt.Println("Sum:", sum)
	}

	// Infinite for statement (usually used with break)
	for {
		fmt.Println("This is an infinite loop")
		break
	}

	// Break and Continue:
	// break is used to terminate a loop prematurely, and continue is used to skip the current iteration and move to the next iteration of the loop.
	// Using break to terminate loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Break - Iteration:", i)
	}

	// Using continue to skip iteration
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Continue - Iteration:", i)
	}

	// For-Range Statement:
	// The for-range statement is used to iterate over elements in various data structures like arrays, slices, strings, maps, etc.
	// Iterating over a slice
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating over a map
	colors := map[string]string{"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Switch Statement:
	// The switch statement is used for multi-way branching based on different cases.
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Some other day")
	}

	// Blank Switches and Goto:
	// Blank switches are useful when handling complex error conditions. goto is rarely used in modern Go code due to its potential to create unreadable and hard-to-debug code.

	y := 10

	// Blank switch
	switch {
	case y > 5:
		fmt.Println("y is greater than 5")
	case y < 5:
		fmt.Println("y is less than 5")
	default:
		fmt.Println("y is equal to 5")
	}

	// Goto example (not recommended)
	goto skip
	fmt.Println("This won't be printed")
skip:
	fmt.Println("Execution continues here")
}
