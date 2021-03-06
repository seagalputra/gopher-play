package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func PrintMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data"`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout, no activities under 5 seconds")
			break loop
		}
	}
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Println("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda: ", menu)
}

func main() {
	runtime.GOMAXPROCS(2)

	//go print(5, "halo")
	//print(5, "apa kabar")

	//var input string
	//fmt.Scanln(&input)

	var messages = make(chan string)

	//	var sayHello = func(who string) {
	//		var data = fmt.Sprintf("hello %s", who)
	//		messages <- data
	//	}
	//
	//	go sayHello("john wick")
	//	go sayHello("ethan hunt")
	//	go sayHello("jason bourne")
	//
	//	var message1 = <-messages
	//	fmt.Println(message1)
	//
	//	var message2 = <-messages
	//	fmt.Println(message2)
	//
	//	var message3 = <-messages
	//	fmt.Println(message3)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	anotherMessages := make(chan int, 2)

	go func() {
		for {
			i := <-anotherMessages

			fmt.Println("Received data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("sending data", i)
		anotherMessages <- i
	}

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}

	var ch3 = make(chan string)
	go sendMessage(ch3)
	PrintMessage(ch3)

	rand.Seed(time.Now().Unix())
	var ch4 = make(chan int)

	go sendData(ch4)
	retrieveData(ch4)

	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	orderSomeFood("pizza")
	orderSomeFood("burger")
}
