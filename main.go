package main

import (
	"fmt"
	. "github.com/mostlybob/LearnAboutTheUS/main"
)

func main() {
	// questions := ShowAllQuestions()
	// fmt.Println(questions)

	// question := GetRandomQuestion()
	// questionId := question.Id
	// fmt.Println(question.Text)

	// question = GetQuestion(questionId)

	quiz := CreateQuizFromJSON(GetQuizJson())
	// question := quiz.GetQuestionById(1)
	// fmt.Println(question.Text)

	question := quiz.GetRandomQuestion()
	fmt.Println(question.Text)
	fmt.Println("(press enter to show answers)")

	var input string
	fmt.Scanln(&input)

	for _, answer := range question.Answers {
		fmt.Println(answer)
	}
}
