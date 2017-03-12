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

func main() {
	// questions := ShowAllQuestions()
	// fmt.Println(questions)

	question := GetRandomQuestion()
	questionId := question.Id
	fmt.Println(question.Text)

	question = GetQuestion(questionId)

}

// ----------------------------------------------------------------------------

func GetQuestion(id int) Question {
	jsonData := GetQuizJson()

	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	var quiz Quiz

	err := quizDecoder.Decode(&quiz)

	if err != nil {
		fmt.Println("Something weird happened trying to open the data file.")
		panic(fmt.Sprintf("%s", err))
	}

	for _, question := range quiz.Questions {
		if question.Id == id {
			return question
		}
	}

	var blankQuestion Question
	return blankQuestion
}

func GetRandomQuestion() Question {
	jsonData := GetQuizJson()

	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	var quiz Quiz

	err := quizDecoder.Decode(&quiz)

	if err != nil {
		fmt.Println("Something weird happened trying to open the data file.")
		panic(fmt.Sprintf("%s", err))
	}

	numberOfQuestions := len(quiz.Questions)
	randomQuestionIndex := getRandomNumber(numberOfQuestions)

	randomQuestion := quiz.Questions[randomQuestionIndex]

	return randomQuestion
}

func getRandomNumber(upper int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNumber := r.Intn(upper - 1)

	return randomNumber
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
