package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Quiz struct {
	About string
	Data  Data
}

type Data struct {
	Questions []Question
}

type Question struct {
	Id             int
	Text           string
	Answers        []string
	AdditionalInfo string
}

func main() {
	questions := ShowAllQuestions()
	fmt.Println(questions)

	question := GetRandomQuestion()
	fmt.Println(question)
}

// ----------------------------------------------------------------------------

func GetRandomQuestion() int {
	jsonData := GetQuizJson()

	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	var quiz Quiz

	err := quizDecoder.Decode(&quiz)

	if err != nil {
		fmt.Println("Something weird happened trying to open the data file.")
		panic(fmt.Sprintf("%s", err))
	}

	numberOfQuestions := len(quiz.Data.Questions)

	return getRandomNumber(0, numberOfQuestions)
}

func getRandomNumber(lower int, upper int) int {
	return 3
}

func ShowAllQuestions() string {
	jsonData := GetQuizJson()

	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	showQuestions := ""
	for {
		var quiz Quiz

		if err := quizDecoder.Decode(&quiz); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Something weird happened trying to open the data file.")
			panic(fmt.Sprintf("%s", err))
		}

		fmt.Printf("Questions:\n")
		for _, question := range quiz.Data.Questions {
			showQuestions += strconv.Itoa(question.Id) + " - " + question.Text + "\n"
		}
	}

	return showQuestions
}

func GetQuizJson() string {
	// read the whole file at once - it's not that big
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	return string(data)
}
