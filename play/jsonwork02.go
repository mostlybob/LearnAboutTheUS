package main

import (
  "encoding/json"
  "fmt"
  "io"
  "log"
  "strings"
)

func main() {
  testString := getTestString()

  //fmt.Println(testString)

  type Question struct {
    id             int
    text           string
    answers        []string
    additionalInfo string
  }

  type Quiz struct {
    About string
    data  []Question
  }

  quizDecoder := json.NewDecoder(strings.NewReader(testString))

  for {
    var quiz Quiz
    if err := quizDecoder.Decode(&quiz); err == io.EOF {
      break
    } else if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("About: %s\n", quiz.About)
  }
}

// ----------------------------------------------------------------------------

func getTestString() string {
  return `{
  "About": "This is the main data file for the quiz.",
  "data": {
    "questions": [
      {
        "id": 1,
        "text": "What is the supreme law of the land?",
        "answers": [
          "the Constitution"
        ],
        "additionalInfo": "The Founding Fathers of the United States wrote the Constitution in 1787. The Constitution is the “supreme law of the land.” The U.S. Constitution has lasted longer than any other country’s constitution. It establishes the basic principles of the United States government. The Constitution establishes a system of government called “representative democracy.” In a representative democracy, citizens choose representatives to make the laws. U.S. citizens also choose a president to lead the executive branch of government. The Constitution lists fundamental rights for all citizens and other people living in the United States. Laws made in the United States must follow the Constitution."
      },
      {
        "id": 2,
        "text": "What does the Constitution do?",
        "answers": [
          "sets up the government",
          "defines the government",
          "protects basic rights of Americans"
        ],
        "additionalInfo": "The Constitution of the United States divides government power between the national government and state governments. The name for this division of power is \"federalism.\" Federalism is an important idea in the Constitution. We call the Founding Fathers who wrote the Constitution the \"Framers\" of the Constitution. The Framers wanted to limit the powers of the government, so they separated the powers into three branches: executive, legislative, and judicial. The Constitution explains the power of each branch. The Constitution also includes changes and additions, called \"amendments.\" The first 10 amendments are called the \"Bill of Rights.\" The Bill of Rights established the individual rights and liberties of all Americans."
      },
      {
        "id": 3,
        "text": "The idea of self-government is in the first three words of the Constitution. What are these words?",
        "answers": [
          "We the People"
        ],
        "additionalInfo": "The Constitution says:\n\n\"We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, provide for the common defence, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, do ordain and establish this Constitution for the United States of America.\"\n\nWith the words \"We the People,\" the Constitution states that the people set up the government. The government works for the people and protects the rights of people. In the United States, the power to govern comes from the people, who are the highest power. This is called \"popular sovereignty.\" The people elect representatives to make laws."
      },
      {
        "id": 4,
        "text": "What is an amendment?",
        "answers": [
          "a change (to the Constitution)",
          "an addition (to the Constitution)"
        ],
        "additionalInfo": "An amendment is a change or addition to the Constitution. The Framers of the Constitution knew that laws can change as a country grows. They did not want to make it too easy to modify the Constitution, the supreme law of the land. The Framers also did not want the Constitution to lose its meaning. For this reason, the Framers decided that Congress could pass amendments in only two ways: by a two-thirds vote in the U.S. Senate and the House of Representatives or by a special convention. A special convention has to be requested by two-thirds of the states. After an amendment has passed in Congress or by a special convention, the amendment must then be ratified (accepted) by the legislatures of three-fourths of the states. The amendment can also be ratified by a special convention in three-fourths of the states. Not all proposed amendments are ratified. Six times in U.S. history amendments have passed in Congress but were not approved by enough states to be ratified."
      },
      {
        "id": 5,
        "text": "What do we call the first ten amendments to the Constitution?",
        "answers": [
          "the Bill of Rights"
        ],
        "additionalInfo": "The Bill of Rights is the first 10 amendments to the Constitution. When the Framers wrote the Constitution, they did not focus on individual rights. They focused on creating the system and structure of government. Many Americans believed that the Constitution should guarantee the rights of the people, and they wanted a list of all the things a government could not do. They were afraid that a strong government would take away the rights people won in the Revolutionary War. James Madison, one of the Framers of the Constitution, wrote a list of individual rights and limits on the government. These rights appear in the first 10 amendments, called the Bill of Rights. Some of these rights include freedom of expression, the right to bear arms, freedom from search without warrant, freedom not to be tried twice for the same crime, the right to not testify against yourself, the right to a trial by a jury of your peers, the right to an attorney, and protection against excessive fines and unusual punishments. The Bill of Rights was ratified in 1791."
      }
    ]
  }
}`
}
