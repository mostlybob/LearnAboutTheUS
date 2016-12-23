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
	questions := ShowQuestions()
	fmt.Println(questions)
}

// ----------------------------------------------------------------------------

func ShowQuestions() string {
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
