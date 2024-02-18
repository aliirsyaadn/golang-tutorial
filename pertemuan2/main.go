package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) multiplyAgeWith10() int {
	return p.age * 10
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	p := person{name: "John", age: 30}
	fmt.Println(p.multiplyAgeWith10())

	// array := [5]int{}
	// fmt.Println(reflect.TypeOf(array).Kind())
	// fmt.Printf("Length of array %d\n", len(array))
	// fmt.Printf("Capacity of array %d\n", cap(array))
	// fmt.Println(array)

	// array2 := [...]string{"a", "b", "c", "d", "e"}
	// fmt.Println(reflect.TypeOf(array2).Kind())
	// fmt.Printf("Length of array2 %d\n", len(array2))
	// fmt.Printf("Capacity of array2 %d\n", cap(array2))
	// fmt.Println(array2)

	// array3 := array2
	// array2[0] = "z"
	// fmt.Println(array3)
	// fmt.Println(array2)

	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(reflect.TypeOf(slice).Kind())
	// fmt.Println(slice)
	// fmt.Printf("Length of slice %d\n", len(slice))
	// fmt.Printf("Capacity of slice %d\n", cap(slice))

	// slice1 := slice[:]
	// slice[0] = 100
	// fmt.Println(slice1)
	// fmt.Println(slice)

	// subSliceWithCap := slice[1:3:5]
	// fmt.Println(subSliceWithCap)
	// fmt.Printf("Length of subSliceWithCap %d\n", len(subSliceWithCap))
	// fmt.Printf("Capacity of subSliceWithCap %d\n", cap(subSliceWithCap))

	// slice2 := make([]string, 5, 6)
	// fmt.Println(reflect.TypeOf(slice2).Kind())
	// fmt.Printf("Length of slice2 %d\n", len(slice2))
	// fmt.Printf("Capacity of slice2 %d\n", cap(slice2))
	// fmt.Println(slice2)
	// slice2 = append(slice2, "a")
	// slice2 = append(slice2, "b")
	// fmt.Println(slice2)
	// fmt.Printf("Length of slice2 %d\n", len(slice2))
	// fmt.Printf("Capacity of slice2 %d\n", cap(slice2))

	// slice3 := append(slice2, "c")
	// fmt.Println(slice3)

	// Declare a rune variable
	// var r rune = 'A'

	// Print the rune and its corresponding code point
	// fmt.Printf("Rune: %c\nCode Point: %U\n", r, r)

	// var b1 byte = 'A'
	// fmt.Printf("Byte: %c\nCode Point: %U\n", b1, b1)

	// Maps
	// var m map[string]int
	// fmt.Println(m)

	// m = make(map[string]int)
	// fmt.Println(m)
	// m["a"] = 1
	// m["b"] = 2
	// fmt.Println(m)

	// delete(m, "a")
	// fmt.Println(m)

	// val, ok := m["a"]
	// fmt.Println(val)
	// fmt.Println(ok)

	// val, ok = m["b"]
	// fmt.Println(val)
	// fmt.Println(ok)

	// Sets
	// set := make(map[string]any)
	// set["a"] = nil
	// set["b"] = nil
	// set["c"] = nil
	// fmt.Println(set)
	// delete(set, "a")
	// fmt.Println(set)

	fmt.Println(multiply(2, 3))
}

func multiply(a, b int) int {
	return a * b
}
