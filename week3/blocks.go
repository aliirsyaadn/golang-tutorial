package main

import "fmt"

func blocks() {
	// Blocks:
	// Blocks in Go are a sequence of statements enclosed in curly braces {}. They define a scope for variables and control structures. Here's an example:
	// Block 1
	{
		x := 10
		fmt.Println("Inside Block 1:", x)
	}

	// Block 2
	{
		x := 20 // Shadowing variable x from Block 1
		fmt.Println("Inside Block 2:", x)
	}

	// Shadowing Variables:
	// Shadowing occurs when a variable in a nested block has the same name as a variable in an outer block. In such cases, the variable in the inner block "shadows" the variable in the outer block. Here's an example:
	y := 10
	fmt.Println("Outer y:", y) // Prints 10

	// Nested block
	{
		y := 20                    // Shadowing the outer y
		fmt.Println("Inner y:", y) // Prints 20
	}

	// The outer y remains unchanged
	fmt.Println("Outer y after shadowing:", y) // Prints 10
}
