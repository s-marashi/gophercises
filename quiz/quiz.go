package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const QuizFileDefault = "problems.csv"
const MaxQuizLength = 100

type Quiz struct {
	questions []Question
	asked     int
	grade     int
}

type Question struct {
	Q string
	A string
}

func main() {
	quizFileName := flag.String("quiz_file", QuizFileDefault, "the file to read quiz from.")
	flag.Parse()

	f, err := os.Open(*quizFileName)
	if err != nil {
		log.Fatal("Failed to open file.")
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	quiz := newQuiz()
	quiz.loadFromCsv(csvReader)

	for {
		err := quiz.askNext()
		if err != nil {
			break
		}
	}

	fmt.Printf("Your grade is : %d\n", quiz.grade)
}

func newQuiz() *Quiz {
	return &Quiz{
		questions: make([]Question, 0, MaxQuizLength),
	}
}

func (quiz *Quiz) addQuestion(question Question) {
	quiz.questions = append(quiz.questions, question)
}

func (quiz *Quiz) loadFromCsv(reader *csv.Reader) {
	for {
		questionCSV, err := reader.Read()
		if err == io.EOF {
			break
		}

		question := Question{
			Q: questionCSV[0],
			A: questionCSV[1],
		}
		quiz.addQuestion(question)
	}
}

func (quiz *Quiz) askNext() error {
	if len(quiz.questions) <= quiz.asked {
		return errors.New("quiz is finished")
	}
	quiz.grade += quiz.questions[quiz.asked].ask()
	quiz.asked += 1

	return nil
}

func (question *Question) ask() int {
	fmt.Println(question.Q)
	var ans string
	_, err := fmt.Scanln(&ans)
	if err != nil {
		log.Fatal("Failed to read user input.")
	}

	if strings.TrimSpace(ans) == strings.TrimSpace(question.A) {
		return 1
	} else {
		return 0
	}
}
