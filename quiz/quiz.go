package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	defaultFileName = "problems.csv"
)

func main() {
	var totalQuestions, totalSuccessfulAnswers int

	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var answer int
		log.Printf("%s ", record[0])
		_, err = fmt.Scanln(&answer)
		if err != nil {
			log.Panic(err)
		}
		correctAnswer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		if answer == correctAnswer {
			totalSuccessfulAnswers++
		}
		totalQuestions++
	}
	log.Printf("Your result is total questions %d total correct answers %d", totalQuestions, totalSuccessfulAnswers)
}
