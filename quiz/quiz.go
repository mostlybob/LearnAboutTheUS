package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Quiz struct {
	About     string
	Questions []Question
	Src       QuizSource
}

type Question struct {
	Id             int
	Text           string
	Answers        []string
	AdditionalInfo string
}

/*
2017-03-22 1255h
Made a bit of a mess here. Yes, I know this isn't the Go way
of doing interfaces - I guess my C# is showing. One thing 
I did learn is that seeding the random number generator
for getting the flash card question belongs in the quiz itself,
since seeds only have to be generated once. I spent some time
looking for something like a constructor and it ended up being
simple, I think. Wherever the Quiz object is constructed, the
Source property, whatever it's called gets set there. Then
when GetRandomQuestion gets called, the Source is already
built and all it has to do is use it. Increasingly, I'm not sure
about where getRandomNumber is living. I can't see where else
to put it, but it's feeling like it doesn't belong where it is.
*/

type QuizSource struct {
	Int63() int64
    Seed(seed int64)
}

func (quiz Quiz) GetRandomQuestion() Question {
	questionIds := quiz.GetQuestionIds()

	index := getRandomNumber(len(questionIds), quiz.Src)

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

// func ShowAllQuestions() string {
// 	jsonData := GetQuizJson()

// 	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

// 	showQuestions := ""
// 	for {
// 		var quiz Quiz

// 		if err := quizDecoder.Decode(&quiz); err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Println("Something weird happened trying to open the data file.")
// 			panic(fmt.Sprintf("%s", err))
// 		}

// 		fmt.Printf("Questions:\n")
// 		for _, question := range quiz.Questions {
// 			showQuestions += strconv.Itoa(question.Id) + " - " + question.Text + "\n"
// 		}
// 	}

// 	return showQuestions
// }

func getRandomNumber(upper int, source Source) int {
	r := rand.New(source)

	randomNumber := r.Intn(upper)

	return randomNumber
}
