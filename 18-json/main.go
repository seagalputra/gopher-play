package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	var arrayJson = `[
    {"Name": "john wick", "Age": 27},
    {"Name": "ethan hunt", "Age": 32}
]`

	var users []User
	err = json.Unmarshal([]byte(arrayJson), &users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", users[0].FullName)
	fmt.Println("user 2:", users[1].FullName)
}
