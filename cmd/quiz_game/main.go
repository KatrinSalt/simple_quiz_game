package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

// type Question struct {
// 	question string
// }

// type Answer struct {
// 	answer string
// }

type Quiz struct {
	questions []string
	answers   []string
}

func main() {
	records, err := readData("test_questions.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output from the csv file: %v\n", records)

	// define a new object of struct Quiz
	// quiz_game := new(Quiz)

	quiz_game := &Quiz{
		questions: []string{},
	}

	// //printing address
	// fmt.Printf("%p\n", quiz_game)
	// fmt.Println(reflect.TypeOf(quiz_game))

	quiz_game.AddItem(records)

	fmt.Printf("%v\n", *quiz_game)
	fmt.Printf("Quiz questions: %v\n", quiz_game.questions)
	fmt.Printf("Quiz answers: %v\n", quiz_game.answers)

}

func readData(fileName string) ([][]string, error) {

	file, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	// // HOW the first line is skipped??
	// // skip first line, check if there are errors
	// if _, err := r.Read(); err != nil {
	// 	return [][]string{}, err
	// }

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func (quiz_game *Quiz) AddItem(records [][]string) {

	//printing address
	fmt.Printf("%p\n", quiz_game)
	fmt.Println(reflect.TypeOf(quiz_game))

	for _, record := range records {
		quiz_question := record[0]
		quiz_answer := record[1]

		quiz_game.questions = append(quiz_game.questions, quiz_question)
		quiz_game.answers = append(quiz_game.answers, quiz_answer)
	}

}

// creating a function for a struct Quiz (creating a method for a class)
// func (quiz *Quiz) AddItems(question Question, answer Answer) ([]Question, []Answer) {

// 	quiz.questions = append(quiz.questions, question)
// 	quiz.answers = append(quiz.answers, answer)

// 	return quiz.questions, quiz.answers
// }
