# Simple Quiz Program in Golang

## Overview

This project is a command-line based quiz program implemented in Go. It allows users to take quizzes stored in CSV files, keeping track of correct and incorrect answers. The program also features a timer functionality where the quiz stops once the specified time limit is reached.

## Features

- Read quiz questions and answers from a CSV file.
- Customizable CSV file name and timer duration via flags.
- Timer functionality stops the quiz after the specified duration.
- Questions are asked one at a time, and the program waits for user input.
- Correct and incorrect answers are tracked and displayed at the end of the quiz.
- Handles CSV files with commas in questions or answers.
- String Trimming: The program performs string cleanup to ensure correct answers with extra whitespace or different capitalization are still considered correct.

## Installation

To use this quiz program, ensure you have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

After installing Go, follow these steps to set up the quiz program:

1. Clone or download the repository to your local machine.
2. Navigate to the directory containing the project files.
3. Build the program using the following command:

```bash
go build .
```

## Usage

To run the quiz program, use the following command format:

```bash
./simple_quiz_game [flags]
```

### Flags

- `-filename quiz_problems.csv`: Specifies the CSV file containing quiz questions and answers. If not provided, the program defaults to `problems.csv`.
- `-timeout time_in_seconds`: Sets the time limit for the quiz in seconds. The default timeout is 30s.

### Example Usage

1. Run the quiz with default settings (quiz.csv, 30s timeout)

   ```bash
   ./quiz
   ```

2. Customize the CSV file and set a timeout:

   ```bash
   ./quiz -csv my_quiz.csv -timeout 60
   ```

## Input Format

The CSV file should contain quiz questions and answers in the following format:

```csv
"What is the capital of Sweden?",Stockholm
"What is the capital of Russia?",Moscow
"When Stockholm was founded?",1252
"Which Swedish pop group won the Eurovision?",ABBA
```

Ensure that questions and answers are separated by commas. Questions with commas should be enclosed in double quotes.