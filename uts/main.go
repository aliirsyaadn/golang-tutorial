package main

import "fmt"

func swap1(x int, y int) {
	fmt.Println("x:", x, "y:", y)
	x, y = y, x
	fmt.Println("x:", x, "y:", y)

}
func swap2(x *int, y *int) {
	*x, *y = *y, *x
}
func main() {
	v1, v2 := 2, 9
	swap1(v1, v2)
	fmt.Println(v1, v2)

	fmt.Println("-------------------")

	swap2(&v1, &v2)
	fmt.Println(v1, v2)

	fmt.Println("-------------------")

	v3 := 3
	fmt.Println("before:", v3)
	v3 = adding10(v3)
	fmt.Println("after:", v3)
}

func adding10(n int) int {
	n = n + 10
	return n
}
