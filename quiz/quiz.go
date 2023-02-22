package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const QuizFileDefault = "problems.csv"

func main() {
	// accept input parameter
	// it is how to get a rough list of arguments
	//fmt.Println(os.Args)

	quizFileName := flag.String("quiz_file", QuizFileDefault, "the file to read quiz from.")
	flag.Parse()

	fmt.Println(*quizFileName)
	// read a text file as a whole
	var data []byte
	var err error

	data, err = os.ReadFile("~/quiz.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	// parse string as csv
	// write a loop
	// write conditionals
	// manipulate variables
	// interpolate string
	// print message

}
