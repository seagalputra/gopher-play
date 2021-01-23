package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person
	age        int
	department string
}

type student struct {
	name  string
	grade int
}

func main() {

	var s1 = student{}
	s1.name = "john wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("name   :", s1.name)
	fmt.Println("grade  :", s1.grade)
	fmt.Println("student 1 : ", s1.name)
	fmt.Println("student 2 : ", s2.name)
	fmt.Println("student 3 : ", s3.name)

	var s4 = student{name: "wick", grade: 2}
	var s5 *student = &s4

	fmt.Println("student 1, name :", s4.name)
	fmt.Println("student 4, name :", s5.name)

	s4.name = "ethan"

	fmt.Println("student 1, name :", s4.name)
	fmt.Println("student 4, name :", s5.name)

	var e1 = employee{}
	e1.name = "wick"
	e1.age = 21
	e1.department = "IT"
	e1.person.age = 22

	fmt.Println("name  :", e1.name)
	fmt.Println("age   :", e1.age)
	fmt.Println("age   :", e1.person.age)
	fmt.Println("department :", e1.department)

	var p2 = person{name: "wick", age: 21}
	var e2 = employee{person: p2, department: "Finance"}

	fmt.Println("name  :", e2.name)
	fmt.Println("age   :", e2.person.age)
	fmt.Println("department :", e2.department)

	/*
		This is an anonymous struct.
		If we only need temporary struct, we can use anonymous struct
	*/
	var e3 = struct {
		person
		grade int
	}{}
	e3.person = person{"wick", 21}
	e3.grade = 2

	fmt.Println("name   :", e3.name)
	fmt.Println("age    :", e3.age)
	fmt.Println("grade  :", e3.grade)

	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
