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

// Define two constants for inform when an answer is wrong or not
const (
	correctAnswer = 1
	wrongAnswer   = 0
)

// Variables we use on this quiz, different channels for sync between go routines and the two flags
// fo setup the file with the quiz and the limit of time
var (
	cAnswer       chan int
	cNextQuestion chan int
	cError        chan error

	fileName         = flag.String("file", "problems.csv", "This flag determine the file name to user")
	timeLimitSeconds = flag.Int("limit", 10, "This is the limit of time in seconds to answer each question")
)

// This method open a csv file and read row per row as soon as receive a signal
// throught the cNextQuestion channel, then verify the answer and return throught
// the cError or cAnswer channels, an error or the result of the answer
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
					cAnswer <- correctAnswer
				} else {
					cAnswer <- wrongAnswer
				}
			}()
		}
	}
}

func main() {
	// Initialize all the variables we are going to use
	flag.Parse()
	var totalQuestions, totalSuccessfulAnswers int
	var err error
	var timer *time.Timer
	cAnswer = make(chan int)
	cNextQuestion = make(chan int)
	cError = make(chan error)

	// Start the go routine where we manage all related with the quiz file
	go quiz()
	// Start a new timer applying a limit in time
	timer = time.NewTimer(time.Duration(*timeLimitSeconds) * time.Second)
	for {
		// Send new question flag for show the next question
		cNextQuestion <- 1
		// Listen to all the channels we have for sync between the quiz go routine and the main
		// go routine and the timer
		select {
		case <-timer.C:
			timer.Reset(time.Duration(*timeLimitSeconds) * time.Second)
		case result := <-cAnswer:
			if result == correctAnswer {
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
