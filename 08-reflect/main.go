package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	fmt.Println("tipe variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())

	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()
	var reflectStudent = reflect.ValueOf(s1)
	var method = reflectStudent.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("john wick"),
	})

	fmt.Println("nama :", s1.Name)
}
