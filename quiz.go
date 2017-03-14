package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Quiz struct {
	About     string
	Questions []Question
}

type Question struct {
	Id             int
	Text           string
	Answers        []string
	AdditionalInfo string
}

func (quiz Quiz) GetRandomQuestion() Question {
	questionIds := quiz.GetQuestionIds()

	index := getRandomNumber(len(questionIds))

	questionId := questionIds[index]

	return quiz.GetQuestionById(questionId)
}

func (quiz Quiz) GetQuestionIds() []int {
	var ids []int

	for _, question := range quiz.Questions {
		ids = append(ids, question.Id)
	}

	return ids
}

func (quiz Quiz) GetQuestionById(id int) Question {
	for _, question := range quiz.Questions {
		if question.Id == id {
			return question
		}
	}

	return Question{} //do we want to throw when question isn't found?
}

func CreateQuizFromJSON(jsonData string) Quiz {
	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	var quiz Quiz

	err := quizDecoder.Decode(&quiz)

	if err != nil {
		fmt.Println("There was an error creating the Quiz from the JSON.")
		panic(fmt.Sprintf("%s", err))
	}

	return quiz
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
		for _, question := range quiz.Questions {
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

func getRandomNumber(upper int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNumber := r.Intn(upper - 1)

	return randomNumber
}
