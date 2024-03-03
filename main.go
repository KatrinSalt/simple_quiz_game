package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type QuizProblem struct {
	question string
	answer   string
}

var (
	correctAnswers int
	userAnswers    []string
)

func main() {
	filename, timeout := readArguments()
	records, err := readCSV(filename)
	if err != nil {
		fmt.Printf("Failed to open and parse the provided cvs file: %s. Error: %v\n", filename, err)
		os.Exit(1)
	}

	quizProblems := prepareQuiz(records)

	fmt.Println("Hello my friend! It's time to have some fun and take an easy-peasy quiz!")
	fmt.Printf("You will have %d seconds to complete the quiz. The time will start after you press \"Enter\" key. Good luck!", timeout)

	// Wait for the user to press "Enter"
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	userAnswers = startQuiz(quizProblems, timeout)
	calculateResult(userAnswers, quizProblems)

	fmt.Printf("\nQuiz is completed! You gave %d correct answers, which means you solved %v%% of the quiz!\n", correctAnswers, (correctAnswers)*100/len(quizProblems))
	fmt.Println("Now relax and grab some fika!")
}

// Adding flags: --filename, --timeout
func readArguments() (string, int) {
	filename := flag.String("filename", "quiz_problems.csv", "A csv file that contains quiz questions'")
	timeout := flag.Int("timeout", 30, "The time limit for the quiz in seconds")
	flag.Parse()
	return *filename, *timeout
}

func readCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

func prepareQuiz(records [][]string) []QuizProblem {
	quizProblems := make([]QuizProblem, len(records))
	for i, record := range records {
		quizProblems[i] = QuizProblem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return quizProblems
}

// Running quiz with a timeout
func startQuiz(quizProblems []QuizProblem, timeout int) []string {
	// Create Timer
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	userAnswerCh := make(chan string)

problemloop:
	for i, problem := range quizProblems {
		fmt.Printf("Question number %d: %s\n", (i + 1), problem.question)
		// Start go routine for getting user's answer
		go getUserAnswer(userAnswerCh)

		select {
		case <-timer.C:
			fmt.Println("\n\nOops, you are out of time!")
			break problemloop
		case answer := <-userAnswerCh:
			userAnswers = append(userAnswers, answer)
		}
	}
	close(userAnswerCh)
	return userAnswers
}

func getUserAnswer(userAnswerCh chan string) {
	fmt.Println("Your answer:")
	reader := bufio.NewReader(os.Stdin)
	userAnswer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		userAnswerCh <- strings.TrimSpace(userAnswer)
	}
}

func calculateResult(userAnswers []string, quizProblems []QuizProblem) {
	for i, answer := range userAnswers {
		if strings.EqualFold(answer, quizProblems[i].answer) {
			correctAnswers++
		}
	}
}
