package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	var input string
	fmt.Print("Type some number:")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
		fmt.Println("Address of number is", &number)
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}

	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
		}()
	}
}
