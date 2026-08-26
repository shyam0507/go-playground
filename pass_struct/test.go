package main

import "fmt"

type User struct {
	Name   string
	Mobile []string
}

func main() {
	shyam := User{Name: "Shyam", Mobile: []string{"1", ""}}
	fmt.Printf("%p\n", &shyam)
	modUser(shyam)
	// fmt.Println(&shyam)
	fmt.Println(shyam)
}

func modUser(u User) {
	fmt.Printf("%p\n", &u)

	u.Mobile = append(u.Mobile, "10")
	u.Mobile[1] = "911"
	u.Name = "Shyamnath"
	fmt.Println(u)
}
