package quiz_test

import (
	. "github.com/mostlybob/LearnAboutTheUS/quiz"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

var _ = Describe("Quiz", func() {
	fmt.Println("Here's my obligatory intro, since Go doesn't allow unused imports :-)")
	var (
		testQuiz Quiz
	)

	BeforeEach(func() {
		testQuiz = Quiz{
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
				Expect(len(testQuiz.Questions)).ToNot(Equal(0))
			})
		})
	})

	Describe("Questions are part of a Quiz", func() {
		Context("A Question has some attributes", func() {
			var question Question

			BeforeEach(func() {
				question = testQuiz.Questions[0]
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
			testQuiz = CreateQuizFromJSON(GetTestJSON())
			Expect(len(testQuiz.Questions)).ToNot(Equal(0))
		})

		Context("Testing the quiz created from JSON", func() {
			var question Question

			BeforeEach(func() {
				testQuiz = CreateQuizFromJSON(GetTestJSON())
				question = testQuiz.Questions[0]
			})

			It("should see the expected values", func() {
				Expect(testQuiz.About).To(Equal("This is the test quiz JSON"))
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
		It("can create a test Quiz object", func() {
			testQuiz = GetTestQuizObject()
			Expect(len(testQuiz.Questions)).ToNot(Equal(0))
			Expect(len(testQuiz.Questions)).To(Equal(5))
		})

		Context("Testing the ability to get a question", func() {
			BeforeEach(func() {
				testQuiz = GetTestQuizObject()
			})

			It("should get a question by Id", func() {
				question := testQuiz.GetQuestionById(1)

				Expect(question.Id).To(Equal(1))
			})

			It("should get a list of question ids", func() {

				refIds := []int{1, 2, 3, 4, 5}
				questionIds := testQuiz.GetQuestionIds()

				for i := 0; i < len(questionIds); i++ {
					Expect(refIds).To(ContainElement(questionIds[i]))
				}

				for i := 0; i < len(refIds); i++ {
					Expect(questionIds).To(ContainElement(refIds[i]))
				}
			})

			It("should be able to get a random question", func() {
				/*
					I want to have a way to to initialize a counter variable for each question
					in the quiz and then run, say, 10 000 iterations of GetRandomQuestion &
					check the counts of each of the variables to see that they are within
					a certain amount of each other. If it's "random" then the higher the number
					of iterations, the closer to equal the counts in all the variables should be.
					I feel a statistics problem here, or maybe calculus, that there should be a formula
					to define some figure for tolerance for variation between the counts.

					e.g. in a true random sampling of the questions, at infinite iterations, the
					counts of occurrences of each of the questions would be equal. As you come back
					from infinity (go toward 0?) the amount of occurrences of any question will differ
					from any other and go up as the number of iterations decrease. The tolerance for
					that difference would have to rise too, perhaps as an invers proportion of the
					number of iterations?

					No doubt I'm overthinking this, but it's a fun little exercise of some very
					rusty math skills.

					Not surprisingly, @darrencauthon had the simple answer that was teasing just
					outside my thinking. Make the number of iterations 10x the number of options
					and make sure all options get hit at least once (or twice or 3 times maybe),
					the idea being that something broken in the random function would leave at least
					one of the counter variables at 0.
				*/
				questionIds := testQuiz.GetQuestionIds()

				numberOfQuestions := len(questionIds)
				iterations := numberOfQuestions * 10
				for i := 0; i < iterations; i++ {
					question := testQuiz.GetRandomQuestion()
					Expect(questionIds).To(ContainElement(question.Id))
				}

				var counters []int

				for i := 0; i < numberOfQuestions; i++ {
					counters = append(counters, 0)
				}

				for i := 0; i < iterations; i++ {
					question := testQuiz.GetRandomQuestion()
					// - will need a different way to enumerate the counters
					// - this way assumes the questions ids are sequential, starting at 1
					// - thinking something like a map (whatever the golang equivalent of a c#
					// dictionary is) would do it

					counters[question.Id-1]++
				}

				fmt.Printf("\n%s: %v\n", "generated array of random questions:", counters)

				for questionId, counterValue := range counters {
					Expect(counterValue > 0).To(BeTrue(), "question id "+string(questionId)+" was never selected")
				}
			})

			It("should be able to display all the questions", func() {
				// questionIds := testQuiz.GetQuestionIds()
				// var testQuestionsDisplay []string

				// for _, questionId := range questionIds {
				// 	questionText := testQuiz.GetQuestionById(questionId)
				// 	questionsLine := string(questionId) + " - " + questionText.Text
				// 	testQuestionsDisplay = append(testQuestionsDisplay, questionsLine)
				// }

				// questions := testQuiz.ShowAllQuestions()

				// for _, question := range questions {
				// 	Expect(testQuestionsDisplay).To(ContainElement(question))
				// }

				fmt.Printf("\n%s\n", "(not sure how ShowAllQuestions should behave yet)")
			})
		})
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
				AdditionalInfo: "some additional information for 1",
			},
			Question{
				Id:             2,
				Text:           "Question 2?",
				Answers:        []string{"bbbbb1", "bbbbb2", "bbbbb3", "bbbbb4"},
				AdditionalInfo: "some additional information for 2",
			},
			Question{
				Id:             3,
				Text:           "Question 3?",
				Answers:        []string{"ccccc1", "ccccc2", "ccccc3", "ccccc4"},
				AdditionalInfo: "some additional information for 3",
			},
			Question{
				Id:             4,
				Text:           "Question 4?",
				Answers:        []string{"ddddd1", "ddddd2", "ddddd3", "ddddd4"},
				AdditionalInfo: "some additional information for 4",
			},
			Question{
				Id:             5,
				Text:           "Question 5?",
				Answers:        []string{"eeeee1", "eeeee2", "eeeee3", "eeeee4"},
				AdditionalInfo: "some additional information for 5",
			},
		},
	}
}
