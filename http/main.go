package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Res struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Req struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// getRequest()
	postRequest()
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response Res
	json.Unmarshal(data, &response)

	fmt.Println(response)

	// fmt.Println(string(data))
}

func postRequest() {
	req := Req{
		Name:  "test",
		Email: "test@test.com",
	}

	data, _ := json.Marshal(&req)

	reqBody := bytes.NewBuffer(data)
	res, err := http.Post("https://postman-echo.com/post", "application/json", reqBody)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(response)
	fmt.Println(res.StatusCode)

	fmt.Println(string(body))
}
