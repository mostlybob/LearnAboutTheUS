package main_test

import (
	. "github.com/mostlybob/LearnAboutTheUS"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

var _ = Describe("Quiz", func() {
	fmt.Println("Here's my obligatory intro, since Go doesn't allow unused imports :-)")
	var (
		quiz Quiz
	)

	BeforeEach(func() {
		quiz = Quiz{
			Questions: []Question{
				Question{
					Id:             1,
					Text:           "Here's some text?",
					Answers:        []string{"aaaaa"},
					AdditionalInfo: "some additional information",
				},
			},
		}
	})

	Describe("First test for Quiz", func() {
		Context("A Quiz has Questions", func() {
			It("should have at least one question", func() {
				Expect(len(quiz.Questions)).ToNot(Equal(0))
			})
		})
	})

	Describe("Questions are part of a Quiz", func() {
		Context("A Question has some attributes", func() {
			var question Question

			BeforeEach(func() {
				question = quiz.Questions[0]
			})

			It("has an id", func() {
				Expect(question.Id).ToNot(Equal(0))
			})

			It("has the question text", func() {
				Expect(question.Text).ToNot(BeEmpty())
			})

			It("has a collection of answers", func() {
				Expect(len(question.Answers)).ToNot(Equal(0))
			})

			It("has additional information", func() {
				Expect(question.AdditionalInfo).ToNot(BeEmpty())
			})
		})
	})

	Describe("Working with JSON", func() {
		It("can create a Quiz from some JSON", func() {
			quiz = CreateQuizFromJSON(`{
  "About": "This is the test quiz JSON",
  "Questions": [
    {
      "Id": 1,
      "Text": "What is JSON?",
      "Answers": [
        "Javascript Object Notation",
        "a data communication standard",
        "much, much better than XML for data communication"
      ],
      "AdditionalInfo": "JSON is a language-independent data format. It derives from JavaScript, but as of 2017 many programming languages include code to generate and parse JSON-format data. The official Internet media type for JSON is application/json. JSON filenames use the extension .json."
    }
  ]
}`)
		})
	})
})
