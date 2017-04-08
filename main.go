package main

import "github.com/mostlybob/LearnAboutTheUS/quiz"
import (
	"fmt"
	"io/ioutil"
)

func main() {
	learnUs := quiz.CreateQuizFromJSON(GetQuizJson("LearnAboutTheUS.json"))

	ShowAQuestion(learnUs)
}

func ShowAQuestion(quiz quiz.Quiz) {
	question := quiz.GetRandomQuestion()
	fmt.Printf("%d - %v\n", question.Id, question.Text)
	fmt.Println("(press Enter to show answers)")

	var input string
	fmt.Scanln(&input)

	for _, answer := range question.Answers {
		fmt.Println(answer)
	}
	fmt.Println("")
}

func GetQuizJson(pathToJson string) string {
	// read the whole file at once - it's not that big
	data, err := ioutil.ReadFile(pathToJson)
	if err != nil {
		panic(err)
	}

	return string(data)
}
