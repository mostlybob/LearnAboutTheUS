package main

// import "github.com/mostlybob/LearnAboutTheUS/quiz"
import "../quiz"

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func getAQuestion(w http.ResponseWriter, r *http.Request) {
	learnUs := quiz.CreateQuizFromJSON(getQuizJson("../LearnAboutTheUS.json"))
	question := showAQuestion(learnUs)
	// should probably look at getting this as a json object, since that's what
	// I want to return from the end point. Simple string for now.
	fmt.Fprintf(w, question)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/getaquestion", getAQuestion)
	cgi.Serve(nil)
}

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	learnUs := quiz.CreateQuizFromJSON(GetQuizJson("../LearnAboutTheUS.json"))

// 	ShowAQuestion(learnUs)
// }

func showAQuestion(quiz quiz.Quiz) string {
	question := quiz.GetRandomQuestion()

	return string(question.Id) + " - " + question.Text

	// fmt.Printf("%d - %v\n", question.Id, question.Text)
	// fmt.Println("(press Enter to show answers)")

	// var input string
	// fmt.Scanln(&input)

	// for _, answer := range question.Answers {
	// 	fmt.Println(answer)
	// }
}

func getQuizJson(pathToJson string) string {
	// read the whole file at once - it's not that big
	data, err := ioutil.ReadFile(pathToJson)
	if err != nil {
		panic(err)
	}

	return string(data)
}
