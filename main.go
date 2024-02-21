package main

import "fmt"

func main() {
	// fmt.Println("Hello, playground")
	// fmt.Println(reflect.TypeOf(os.Args).Kind())
	// fmt.Println(len(os.Args))
	// fmt.Println(os.Args)

	// For loop
	// for i := 0; i < 5; i += 2 {
	// 	fmt.Println(i)
	// }
	// i := 0
	// for ; i < 7; i += 3 {
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)

	// // Foreach loop
	// slice := []string{"a", "b", "c", "d", "e"}
	// for i, val := range slice {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, val)
	// }
	// fmt.Println(multiply(2, 3))
	// fmt.Println(divide.Divide(10, 2))

	// x := 10
	// fmt.Println("Outer x:", x) // Prints 10

	// // Nested block
	// {
	// 	x := 20                    // Shadowing the outer x
	// 	fmt.Println("Inner x:", x) // Prints 20
	// }

	// fmt.Println("Outer x:", x) // Prints 10

	// If-else
	// if x > 5 {
	// 	fmt.Println("x is greater than 5")
	// } else if x < 5 {
	// 	fmt.Println("x is less than 5")
	// } else {
	// 	fmt.Println("x is equal to 5")
	// }

	// day := "Monday"

	// switch day {
	// case "Monday":
	//     fmt.Println("Today is Monday")
	// case "Tuesday":
	//     fmt.Println("Today is Tuesday")
	// default:
	//     fmt.Println("Some other day")
	// }

	// day := "Monday"
	// if day == "Monday" {
	// 	fmt.Println("Today is Monday")
	// }
	// if day == "Tuesday" {
	// 	fmt.Println("Today is Tuesday")
	// }
	// if day != "Monday" && day != "Tuesday" {
	// 	fmt.Println("Some other day")
	// }

	creditCard := "1357261567325172365"

	sumEven := 0
	sumOdd := 0
	for i, digit := range creditCard {
		if i%2 == 0 {
			fmt.Println("Digit:", string(digit))
			fmt.Println("ini ganjil")
		} else {
			fmt.Println("Digit:", string(digit))
			fmt.Println("ini genap")
			sumEven += int(digit * 2)
		}
	}

	sum := sumEven + sumOdd
	if sum%10 == 0 {
		fmt.Println("Valid Credit Card")
	} else {
		fmt.Println("Invalid Credit Card")
	}
}

func multiply(a, b int) int {
	return a * b
}
