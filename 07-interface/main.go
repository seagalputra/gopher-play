package main

import (
	"fmt"
	"math"
	"strings"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungLain interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

type person struct {
	name string
	age  int
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas        :", bangunDatar.luas())
	fmt.Println("keliling    :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas        :", bangunDatar.luas())
	fmt.Println("keliling    :", bangunDatar.keliling())
	fmt.Println("jari-jari   :", bangunDatar.(lingkaran).jariJari())

	var bangunRuang hitungLain = &kubus{4}
	fmt.Println("===== kubus")
	fmt.Println("luas        :", bangunRuang.luas())
	fmt.Println("keliling    :", bangunRuang.keliling())
	fmt.Println("volume      :", bangunRuang.volume())

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "mango", "banana"},
	}

	fmt.Println(data)

	var another interface{}

	another = 2
	var number = another.(int) * 10
	fmt.Println(another, "multiplied by 10 is :", number)

	another = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(another.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	another = &person{name: "wick", age: 27}
	var name = another.(*person).name
	fmt.Println(name)
}
