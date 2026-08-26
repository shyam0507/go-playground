package main

import "fmt"

type Shape interface {
	Area() int
}

type Rectangle struct {
	length int
	width  int
}

func main() {

	//creating a struct
	set := map[string]struct{}{}
	set["1"] = struct{}{}
	set["2"] = struct{}{}

	//
	var rect1 Shape
	rect1 = &Rectangle{
		length: 10,
		width:  20,
	}

	fmt.Println(rect1.Area())

}

func (r *Rectangle) Area() int {
	return r.length * r.width
}
