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
	defaultTimeLimitSeconds = 10
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
		select {
		case <-cNextQuestion:
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
}

func main() {
	var totalQuestions, totalSuccessfulAnswers int
	var err error
	var timer *time.Timer
	cAnswer = make(chan int)
	cNextQuestion = make(chan int)
	cError = make(chan error)

	go quiz()
	timer = time.NewTimer(time.Duration(defaultTimeLimitSeconds) * time.Second)
	for {
		cNextQuestion <- 1
		select {
		case <-timer.C:
			log.Println("Time expired")
			timer.Reset(time.Duration(defaultTimeLimitSeconds) * time.Second)
			log.Println("Timer reset")
		case result := <-cAnswer:
			log.Println("Receive value")
			if result == 1 {
				totalSuccessfulAnswers++
			}
			timer.Stop()
			timer.Reset(time.Duration(defaultTimeLimitSeconds) * time.Second)
		case err = <-cError:
			if err != io.EOF {
				log.Panic(err)
			}
			timer.Stop()
		}
		log.Println("Increment total questions counter")
		totalQuestions++
		if err != nil {
			break
		}
	}
	log.Printf("Your result is total questions %d total correct answers %d", totalQuestions, totalSuccessfulAnswers)
}
