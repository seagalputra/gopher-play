package main

import (
	"fmt"
	"os"
	"time"
)

func myTimer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")

	//for true {
	//	fmt.Println("Hello !!")
	//	time.Sleep(1 * time.Second)
	//}

	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	//done := make(chan bool)
	//ticker := time.NewTicker(time.Second)

	//go func() {
	//	time.Sleep(10 * time.Second)
	//	done <- true
	//}()

	//for {
	//	select {
	//	case <-done:
	//		ticker.Stop()
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("Hello !!", t)
	//	}
	//}

	var timeout = 5

	go myTimer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}
