package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	defaultFileName         = "problems.csv"
	defaultTimeLimitSeconds = 30
)

var (
	cTimer        chan time.Time
	cAnswer       chan int
	cNextQuestion chan int
	cError        chan error
)

func quiz() {
	file, err := os.Open("problems.csv")
	if err != nil {
		cError <- err
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		<-cNextQuestion
		record, err := r.Read()
		if err == io.EOF {
			cError <- err
		}
		if err != nil {
			cError <- err
		}
		var answer int
		log.Printf("%s ", record[0])
		_, err = fmt.Scanln(&answer)
		if err != nil {
			cError <- err
		}
		correctAnswer, err := strconv.Atoi(record[1])
		if err != nil {
			cError <- err
		}
		if answer == correctAnswer {
			cAnswer <- 1
		} else {
			cAnswer <- 0
		}
	}
}

func main() {
	var totalQuestions, totalSuccessfulAnswers int
	var err error
	cTimer = make(chan time.Time)
	cAnswer = make(chan int)
	cNextQuestion = make(chan int)
	cError = make(chan error)

	go quiz()
	for {
		cNextQuestion <- 1
		select {
		case <-cTimer:
			log.Println("Time expired")
		case result := <-cAnswer:
			log.Println("Receive value")
			if result == 1 {
				totalSuccessfulAnswers++
			}
		case err = <-cError:
			if err != io.EOF {
				log.Panic(err)
			}
		}
		totalQuestions++
		if err != nil {
			break
		}
	}
	log.Printf("Your result is total questions %d total correct answers %d", totalQuestions, totalSuccessfulAnswers)
}
