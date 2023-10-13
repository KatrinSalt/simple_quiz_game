package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
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

var (
	correct_answers int
	wrong_answers   int
)

func main() {

	records, err := readData("test_questions.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output from the csv file: %v\n", records)

	// define a new object of struct Quiz
	// quiz_game := new(Quiz)

	quiz_game := &Quiz{}
	quiz_game.AddItem(records)

	fmt.Printf("Number of questions: %v\n", len(quiz_game.questions))
	fmt.Printf("Quiz answers: %v\n", quiz_game.answers)

	fmt.Println("Hello my friend. What is your name?")
	name, err := readUserInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Nice to meet you, %v! It's time to take a Quiz and have fun!\n", name)

	for num, question := range quiz_game.questions {

		fmt.Printf("Question number %v: %v\n", (num + 1), question)
		answer, err := readUserInput()
		if err != nil {
			log.Fatal(err)
		}

		calculateResult(answer, quiz_game.answers[num])
	}

	fmt.Printf("Congratulations, %v! You answered all the quiz questions. Here is your result\n", name)
	fmt.Printf("Number of correct answers: %v\nNumber of wrong answers %v\n", correct_answers, wrong_answers)
	fmt.Printf("You solved %v%% of the quiz. Now relax and drink your beer:)\n", (correct_answers)*100/len(quiz_game.questions))
}

func readData(fileName string) ([][]string, error) {

	// Multi-Dimensional slice:
	// [][]string Multi-dimensional slice are
	// just like the multidimensional array, except that slice does not contain the size.

	// Example:
	// Creating multi-dimensional slice
	// s1 := [][]int{{12, 34},
	// {56, 47},
	// {29, 40},
	// {46, 78}

	// Output: [[12 34] [56 47] [29 40] [46 78]]

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
	// fmt.Printf("%p\n", quiz_game)
	// fmt.Println(reflect.TypeOf(quiz_game))

	for _, record := range records {
		quiz_question := record[0]
		quiz_answer := record[1]

		quiz_game.questions = append(quiz_game.questions, quiz_question)
		quiz_game.answers = append(quiz_game.answers, quiz_answer)
	}

}

func readUserInput() (string, error) {
	for {
		fmt.Println("Your answer:")
		reader := bufio.NewReader(os.Stdin)
		user_answer, err := reader.ReadString('\n')
		if err != nil {
			return user_answer, err
		} else if user_answer != "\n" && strings.TrimSpace(user_answer) != "" {
			return strings.TrimSpace(user_answer), nil
		} else {
			fmt.Println("You should provide an answer even if it is wrong! Please, try again")
			continue
		}
	}
}

// do you always need to return error?
func checkAnswer(userAnswer string, correctAnswer string) bool {
	// strings.EqualFold() Function in Golang reports whether s and t,
	// interpreted as UTF-8 strings, are equal under Unicode case-folding,
	// which is a more general form of case-insensitivity.
	return strings.EqualFold(userAnswer, correctAnswer)
}

func calculateResult(userAnswer string, correctAnswer string) {
	if checkAnswer(userAnswer, correctAnswer) {
		correct_answers++
	} else {
		wrong_answers++
	}
}
