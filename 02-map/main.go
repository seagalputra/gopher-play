package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	var data map[string]int
	data = map[string]int{}
	data["one"] = 1

	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(chicken1)

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	// iterate using for - range
	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
