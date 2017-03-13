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
			quiz = CreateQuizFromJSON(GetTestJSON())
			Expect(len(quiz.Questions)).ToNot(Equal(0))
		})

		Context("Testing the quiz created from JSON", func() {

			var question Question

			BeforeEach(func() {
				quiz = CreateQuizFromJSON(GetTestJSON())
				question = quiz.Questions[0]
			})

			It("should see the expected values", func() {
				Expect(quiz.About).To(Equal("This is the test quiz JSON"))
				Expect(question.Id).To(Equal(1))
				Expect(question.Text).To(Equal("What is JSON?"))
				Expect(question.AdditionalInfo).To(Equal("JSON is a language-independent data format. It derives from JavaScript, but as of 2017 many programming languages include code to generate and parse JSON-format data. The official Internet media type for JSON is application/json. JSON filenames use the extension .json."))
				Expect(len(question.Answers)).To(Equal(3))
				Expect(question.Answers[0]).To(Equal("Javascript Object Notation"))
				Expect(question.Answers[1]).To(Equal("a data communication standard"))
				Expect(question.Answers[2]).To(Equal("much, much better than XML for data communication"))
			})
		})
	})

	Describe("Functions of the Quiz object", func() {
	})

})

func GetTestJSON() string {
	return `
				{
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
			    }
			`
}

func GetTestQuizObject() Quiz {
	return Quiz{
		Questions: []Question{
			Question{
				Id:             1,
				Text:           "Question 1?",
				Answers:        []string{"aaaaa1", "aaaaa2", "aaaaa3", "aaaaa4"},
				AdditionalInfo: "some additional information",
			},
			Question{
				Id:             2,
				Text:           "Question 2?",
				Answers:        []string{"bbbbb1", "bbbbb2", "bbbbb3", "bbbbb4"},
				AdditionalInfo: "some additional information",
			},
			Question{
				Id:             3,
				Text:           "Question 3?",
				Answers:        []string{"ccccc1", "ccccc2", "ccccc3", "ccccc4"},
				AdditionalInfo: "some additional information",
			},
			Question{
				Text:           "Question 4?",
				Answers:        []string{"ddddd1", "ddddd2", "ddddd3", "ddddd4"},
				AdditionalInfo: "some additional information",
			},
			Question{
				Id:             5,
				Text:           "Question 5?",
				Answers:        []string{"eeeee1", "eeeee2", "eeeee3", "eeeee4"},
				AdditionalInfo: "some additional information",
			},
		},
	}
}
