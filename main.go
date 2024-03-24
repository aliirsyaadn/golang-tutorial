package main

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

	// creditCard := "1357261567325172365"

	// sumEven := 0
	// sumOdd := 0
	// for i, digit := range creditCard {
	// 	if i%2 == 0 {
	// 		fmt.Println("Digit:", string(digit))
	// 		fmt.Println("ini ganjil")
	// 	} else {
	// 		fmt.Println("Digit:", string(digit))
	// 		fmt.Println("ini genap")
	// 		sumEven += int(digit * 2)
	// 	}
	// }

	// sum := sumEven + sumOdd
	// if sum%10 == 0 {
	// 	fmt.Println("Valid Credit Card")
	// } else {
	// 	fmt.Println("Invalid Credit Card")
	// }

	// a := "1"
	// digitInt, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(err)
	// fmt.Println(digitInt)

	// a = "1"
	// var digitInt2 int64
	// digitInt2, err = strconv.ParseInt(a, 10, 32)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(digitInt2)

	// creditNumber := "1263576153762512"

	// for i, digit := range creditNumber {
	// 	fmt.Println("Index:", i, "Digit:", string(digit))
	// }

	// for i := 0; i < len(creditNumber); i++ {
	// 	fmt.Println("Index:", i, "Digit:", string(creditNumber[i]))
	// }

	// fmt.Println(creditNumber)

	// Memory allocation 0x1 0x2 0x3 0x4 0x5 0x6

	// x := []string{"a", "b", "c", "d", "e"}
	// x := make([]string, 5, 6)
	// x[0] = "a"
	// x[1] = "b"
	// x[2] = "c"
	// x[3] = "d"
	// x[4] = "e"
	// fmt.Println("cap x:", cap(x))
	// y := x[2:4] // 0x3 0x4 0x5
	// fmt.Println("cap y:", cap(y))
	// y = append(y, "f") // 0x3 0x4 0x5 0x6
	// fmt.Println("setelah appendcap y:", cap(y))
	// fmt.Println("x: ", x) // [a b c d e]
	// fmt.Println("y: ", y) // [c d e f]

	// x := make([]int, 4, 6)
	// fmt.Println("cap x:", cap(x)) // 6
	// fmt.Println("len x:", len(x)) // 5
	// fmt.Println("x:", x)          // [0 0 0 0 0]
	// [0x812378379 0x812378380 0x812378381 0x812378382 0x812378383]

	// ["a", "b", "c", "d", "e"]

	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2:2]
	// z := x[2:4:4]
	// fmt.Println(cap(x), cap(y), cap(z))
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// y = append(y, "i", "j", "k")
	// x = append(x, "88")
	// z = append(z, "99")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := []string{"a", "b", "c", "d", "e"}
	// y := x[2:4:4]
	// y = append(y, "f")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// xArray := [5]int{1, 2, 3, 4, 5}
	// xSlice := xArray[:]
	// fmt.Println("tipe xArray:", reflect.TypeOf(xArray).Kind())
	// fmt.Println("tipe xSlice:", reflect.TypeOf(xSlice).Kind())
	// fmt.Println("xArray:", xArray)
	// fmt.Println("xSlice:", xSlice)
	// fmt.Println("cap xArray:", cap(xArray))
	// fmt.Println("cap xSlice:", cap(xSlice))
	// xSlice = append(xSlice, 6)
	// fmt.Println("Setelah Append")
	// fmt.Println("xArray:", xArray)
	// fmt.Println("xSlice:", xSlice)
	// fmt.Println("cap xArray:", cap(xArray))
	// fmt.Println("cap xSlice:", cap(xSlice))
	// xSlice = append(xSlice, 7, 8, 9, 10, 11)
	// fmt.Println("cap xSlice:", cap(xSlice))

	// xSlice := []int{1, 2, 3, 4, 5}
	// xArray := [4]int(xSlice)
	// fmt.Println("xArray:", xArray)
	// fmt.Println("xSlice:", xSlice)
}

func multiply(a, b int) int {
	return a * b
}
