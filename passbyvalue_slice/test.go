package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	changeSlice(slice1)
	fmt.Println(slice1)

	// slice2 := []int{5, 6, 7}
	slice2 := make([]int, 0, 4)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 6)
	slice2 = append(slice2, 7)
	test := &slice2[0]
	fmt.Println(test)
	fmt.Println(slice2)
	addSlice(slice2)
	fmt.Println(slice2)
}

func changeSlice(v []int) {
	v[2] = 10
}

func addSlice(v []int) {
	v[0] = 10
	v = append(v, 100)
	v[0] = 11
	fmt.Println(v)
	fmt.Println(&v[0])

}

// func processsliceayPoint(v *[3]int) {
// 	v[2] = 10
// }
