package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Quiz struct {
	About string
	Data  Data
}

type Data struct {
	Questions []struct {
		Id             int
		Text           string
		Answers        []string
		AdditionalInfo string
	}
}

func main() {
	jsonData := GetQuizJson()

	quizDecoder := json.NewDecoder(strings.NewReader(jsonData))

	for {
		var quiz Quiz

		if err := quizDecoder.Decode(&quiz); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Something weird happened trying to open the data file.")
			panic(fmt.Sprintf("%s", err))
		}

		fmt.Printf("%s\n", quiz.About)

		fmt.Printf("Questions:\n")
		for _, question := range quiz.Data.Questions {
			fmt.Printf("\t%d - %s\n", question.Id, question.Text)
		}
	}
}

// ----------------------------------------------------------------------------

func GetQuizJson() string {
	// read the whole file at once - it's not that big
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	return string(data)
}
