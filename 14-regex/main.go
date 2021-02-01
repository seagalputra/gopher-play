package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	regex, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	str := regex.FindString(text)
	fmt.Println(str)

	idx := regex.FindStringIndex(text)
	fmt.Println(idx)
	fmt.Println(text[idx[0]:idx[1]])
}
