package main

import "fmt"

func main() {
	dict := map[string]int{}
	dict["One"] = 1
	dict["Two"] = 2
	dict["Three"] = 3
	fmt.Println(dict)

	for k, v := range dict {
		fmt.Print(k)
		fmt.Println(v)
	}
}
