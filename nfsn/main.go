package main

import "github.com/mostlybob/LearnAboutTheUS/quiz"

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	http.HandleFunc("/", hello)
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

// func ShowAQuestion(quiz quiz.Quiz) {
// 	question := quiz.GetRandomQuestion()
// 	fmt.Printf("%d - %v\n", question.Id, question.Text)
// 	fmt.Println("(press Enter to show answers)")

// 	var input string
// 	fmt.Scanln(&input)

// 	for _, answer := range question.Answers {
// 		fmt.Println(answer)
// 	}
// }

// func GetQuizJson(pathToJson string) string {
// 	// read the whole file at once - it's not that big
// 	data, err := ioutil.ReadFile(pathToJson)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return string(data)
// }
