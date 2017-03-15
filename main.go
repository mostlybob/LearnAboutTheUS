package main

import "github.com/mostlybob/LearnAboutTheUS/quiz"
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// questions := ShowAllQuestions()
	// fmt.Println(questions)

	// question := GetRandomQuestion()
	// questionId := question.Id
	// fmt.Println(question.Text)

	// question = GetQuestion(questionId)

	foo := quiz.CreateQuizFromJSON(GetQuizJson("data.json"))
	// question := quiz.GetQuestionById(1)
	// fmt.Println(question.Text)

	question := foo.GetRandomQuestion()
	fmt.Println(question.Text)
	fmt.Println("(press enter to show answers)")

	var input string
	fmt.Scanln(&input)

	for _, answer := range question.Answers {
		fmt.Println(answer)
	}
}

func GetQuizJson(pathToJson string) string {
	// read the whole file at once - it's not that big
	data, err := ioutil.ReadFile(pathToJson)
	if err != nil {
		panic(err)
	}

	return string(data)
}
