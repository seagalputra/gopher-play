package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

func sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func main() {
	var firstName string = "John"

	var lastName string
	lastName = "Wick"

	myName := "Dio"

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name := new(string)

	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Println(greeting(firstName + " " + lastName))
	fmt.Println(greeting(myName))
	fmt.Println(first, second, third)
	fmt.Println(seventh, eight, ninth)
	fmt.Println("Alamat memori dari variabel name : ", name)
	fmt.Println("Isi data dari variabel name : ", *name)
	fmt.Println("10 + 20 =", sum(10, 20))

	// condition with temporary variabel
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Println("Percent :", percent)
	} else {
		fmt.Println("oops")
	}

	switch point {
	case 10:
		fmt.Println("Wohooo..")
	case 5, 8840.0:
		fmt.Println("ahaay")
	default:
		fmt.Println("No No")
	}

	var names [4]string
	names[0] = "trafaglar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi semua elemen \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	var fruitBuckets = make([]string, 2)
	fruitBuckets[0] = "apple"
	fruitBuckets[1] = "manggo"

	fmt.Println(fruitBuckets)

	var fruitSlice = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruitSlice[0:2]

	fmt.Println(newFruits)

	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	var cFruits = append(fruitSlice, "papaya")

	fmt.Println(fruits)
	fmt.Println(cFruits)

	dst := make([]string, 3)
	src := []string{"watermelon", "pinneaple", "apple", "orange"}
	n := copy(dst, src)
	copied := append(dst, "oooo")
	fmt.Println(n)
	fmt.Println(copied)
}
