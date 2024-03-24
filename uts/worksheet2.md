# Worksheet 2
## 1. 
Write a program that defines a variable named greetings of type slice of strings with
the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
Create a subslice containing the first two values; a second subslice with the second,
third, and fourth values; and a third subslice with the fourth and fifth values. Print out
all four slices.
## Answer:
```go
package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	subslice1 := greetings[:2]
	subslice2 := greetings[1:4]
	subslice3 := greetings[3:]

	fmt.Println("Original Slice:", greetings)
	fmt.Println("Subslice 1:", subslice1)
	fmt.Println("Subslice 2:", subslice2)
	fmt.Println("Subslice 3:", subslice3)
}
```
## 2. 
Write a program that defines a string variable called message with the value "Hi ☺
and " and prints the fourth rune in it as a character, not a number.
## Answer:
```go
package main

import "fmt"

func main() {
	message := "Hi ☺ and "
	fmt.Println(string([]rune(message)[3]))
}

```

## 3. 
Write a program that defines a struct called Employee with three fields: firstName,
lastName, and id. The first two fields are of type string, and the last field (id) is of
type int. Create three instances of this struct using whatever values you’d like.
Initialize the first one using the struct literal style without names, the second using the
struct literal style with names, and the third with a var declaration. Use dot notation to
populate the fields in the third struct. Print out all three structs.
## Answer:
```go
package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	// First instance using struct literal style without names
	emp1 := Employee{"John", "Doe", 1001}

	// Second instance using struct literal style with names
	emp2 := Employee{firstName: "Alice", lastName: "Smith", id: 1002}

	// Third instance using var declaration and dot notation to populate fields
	var emp3 Employee
	emp3.firstName = "Bob"
	emp3.lastName = "Jones"
	emp3.id = 1003

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)
	fmt.Println("Employee 3:", emp3)
}
```

## 4. 
Define a function minMax which accepts a non-empty slice of integers and returns
the minimum and maximum elements of the given slice.
Examples of usage:
```go
ls := []int{3, 1, 7, 0, -3, 5, 6, 3}
fmt.Println(minMax(ls)) //-3 7
```
## Answer:
```go
package main

import "fmt"

func minMax(nums []int) (min, max int) {
	min = nums[0]
	max = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	ls := []int{3, 1, 7, 0, -3, 5, 6, 3}
	fmt.Println(minMax(ls)) // Output: -3 7
}
```

## 5. 
Find and explain the output of the following program snippet:
```go
x := 25
px := &x
ppx := &px
fmt.Printf("%T -- %v -- %T -- %v\n", ppx, **ppx, px, *px)
```
## Answer:
* The program snippet declares a variable x with value 25.
* It then declares a pointer px pointing to the address of x, and a pointer ppx pointing to the address of px.
* The Printf statement prints the type and value of ppx (pointer to pointer), the dereferenced value of ppx (double-dereferenced value), the type and value of px (pointer), and the dereferenced value of px.
* Output:
```go
**int -- 25 -- *int -- 25
```

## 6. 
Find and explain the output of the following program snippet:
```go
func main() {
    a := 22
    f := func() {
    fmt.Println(a)
    a := 33
    }
    f()
    fmt.Println(a)
}
```
## Answer:
* The program declares a variable a with value 22.
* It defines a function f that prints the value of a and then declares a new variable a with value 33.
* Inside the main function, f is called, which prints the value of the inner a (which is 33) and then returns.
* Finally, the main function prints the value of the outer a (which is 22).
* Output:
```go
22
22
```

## 7. 
Find and explain the output of this program snippet:
```go
func swap1(x int, y int) {
    x, y = y, x
}
func swap2(x *int, y *int) {
    *x, *y = *y, *x
}
func main() {
    v1, v2 := 2, 9
    swap1(v1, v2)
    fmt.Println(v1, v2)
    swap2(&v1, &v2)
    fmt.Println(v1, v2)
}
```
## Answer:
* The program defines two functions, swap1 and swap2, to swap the values of two integers.
* In the main function, it declares two integers v1 and v2 with initial values 2 and 9, respectively.
* It calls swap1 with v1 and v2, but since swap1 takes parameters by value, it doesn't affect the original variables.
* Then, it calls swap2 with pointers to v1 and v2, so it swaps the values of v1 and v2.
* Finally, it prints the values of v1 and v2.
* Output:
```go
2 9
9 2
```