package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// Generate random number with fixed sequence of generated value
	//rand.Seed(10)
	//fmt.Println("random ke-1:", rand.Int())
	//fmt.Println("random ke-2:", rand.Int())
	//fmt.Println("random ke-3:", rand.Int())

	// Generate random unique number
	// by injecting seed with unique number
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1:", rand.Int())
	fmt.Println("random ke-2:", rand.Int())
	fmt.Println("random ke-3:", rand.Int())

	fmt.Println("random:", randomString(10))
}
