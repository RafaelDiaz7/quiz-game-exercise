package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("You have", exerciseCheck(getCsvRecords()), "correct answers.")
}

func getCsvRecords() (csvRecords [][]string) {
	csvData, err := os.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(csvData)))
	r.Comma = ','

	csvRecords, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func exerciseCheck(csvRecords [][]string) (counter int) {
	counter = 0
	for _, record := range csvRecords {
		fmt.Println(record[0], "= ")
		correctAnswer := record[1]
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		userAnswer := scanner.Text()
		if userAnswer == correctAnswer {
			counter += 1
		}
	}

	return
}

//func getExercise() (question string, correctAnswer string) {
//	question = record[0]
//	correctAnswer = record[1]
//	return
//}
