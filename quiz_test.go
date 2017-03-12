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
})
