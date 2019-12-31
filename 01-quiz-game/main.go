package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var score, counter int
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(csvFile)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if answer, err := strconv.Atoi(record[1]); err == nil {
			fmt.Printf("what is %s, sir? ", record[0])
			scanner.Scan()
			text := scanner.Text()
			predict, _ := strconv.Atoi(text)

			if predict == answer {
				score++
			}
		}
		counter++
	}
	fmt.Printf("Your score is %d out of %d\n", score, counter)
}
