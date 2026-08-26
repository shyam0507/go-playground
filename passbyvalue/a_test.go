package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	processArray(arr1)
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	processArrayPoint(&arr1)
	fmt.Println(arr1)
}

func processArray(v [3]int) {
	v[2] = 10
}

func processArrayPoint(v *[3]int) {
	v[2] = 10
}
