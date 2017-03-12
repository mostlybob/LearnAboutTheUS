package main_test

import (
	. "github.com/mostlybob/LearnAboutTheUS"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Quiz", func() {

	var (
		quiz Quiz
	)

	BeforeEach(func() {
		quiz = Quiz{}
		quiz.Data.Questions = append(quiz.Data.Questions, Question{Id: 1})
	})

	Describe("First test for Quiz", func() {
		Context("A Quiz has Questions", func() {
			It("should have at least one question", func() {
				Expect(len(quiz.Data.Questions)).ToNot(Equal(0))
			})
		})
	})

	Describe("Questions are part of a Quiz", func() {
		Context("A Question has some attributes", func() {
			question := quiz.Data.Questions[0]

			Expect(question.Id).ToNot(Equal(0))
		})
	})
})
