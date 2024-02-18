package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	val := 5
	fmt.Println(val)

	val = 10
	fmt.Println(val)

	fmt.Println(val * 10)

	// write code quicksorting
	arr = []int{5, 4, 3, 2, 1}

	var i int = 10
	var f float64 = 10.21231231231
	str := "ABCDEF"
	fmt.Printf("Integer i: %d\n", i)
	fmt.Printf("Floating-point f: %f\n", f)
	fmt.Printf("String str: %s\n", str)
	i = 11

	fmt.Printf("Ini adalah integer %d\n dan ini adalah float %.2f\n dan ini adalah string %s\n", i, f, str)
}
