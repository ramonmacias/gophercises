package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	cTimer        chan time.Time
	cAnswer       chan int
	cNextQuestion chan int
	cError        chan error

	fileName         = flag.String("file", "problems.csv", "This flag determine the file name to user")
	timeLimitSeconds = flag.Int("limit", 10, "This is the limit of time in seconds to answer each question")
)

func quiz() {
	file, err := os.Open(*fileName)
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
			go func() {
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
			}()
		}
	}
}

func main() {
	flag.Parse()
	var totalQuestions, totalSuccessfulAnswers int
	var err error
	var timer *time.Timer
	cAnswer = make(chan int)
	cNextQuestion = make(chan int)
	cError = make(chan error)

	go quiz()
	timer = time.NewTimer(time.Duration(*timeLimitSeconds) * time.Second)
	for {
		cNextQuestion <- 1
		select {
		case <-timer.C:
			timer.Reset(time.Duration(*timeLimitSeconds) * time.Second)
		case result := <-cAnswer:
			if result == 1 {
				totalSuccessfulAnswers++
			}
			timer.Stop()
			timer.Reset(time.Duration(*timeLimitSeconds) * time.Second)
		case err = <-cError:
			if err != io.EOF {
				log.Panic(err)
			}
			timer.Stop()
		}
		totalQuestions++
		if err != nil {
			break
		}
	}
	log.Printf("Your result is total questions %d total correct answers %d", totalQuestions, totalSuccessfulAnswers)
}
