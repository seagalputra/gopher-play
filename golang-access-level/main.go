package main

import (
	f "fmt"
	. "golang-access-level/library"
)

func main() {
	//var s1 = Student{Name: "ethan", Grade: 21}
	//f.Println("name ", s1.Name)
	//f.Println("grade", s1.Grade)

	sayHello("ethan")

	f.Printf("Name : %s\n", Student.Name)
}
