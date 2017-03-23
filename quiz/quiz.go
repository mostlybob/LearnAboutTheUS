package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	seed := time.Now().UnixNano()
	questionIds := quiz.GetQuestionIds()

	index := getRandomNumber(len(questionIds), seed)

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

func (quiz Quiz) ShowAllQuestions() []string {
	return []string{}
}

func getRandomNumber(upper int, seed int64) int {
	return rand.New(rand.NewSource(seed)).Intn(upper)
}
