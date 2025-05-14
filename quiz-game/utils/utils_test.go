package utils

import (
	"log"
	"testing"
)

// Valid mock dependencies
type ValidMockQuizDeps struct{}

func (m *ValidMockQuizDeps) parseCommand() (file string, limit int, help bool, err error) {
	return file, limit, help, err
}

func (m *ValidMockQuizDeps) getQuestions(filepath string) ([]Question, error) {
	if filepath == "invalid" {
		return []Question{}, nil
	}

	return []Question{
		{question: "7+3", answer: "10"},
		{question: "4+4", answer: "8"},
		{question: "1+3", answer: "5"},
		{question: "2+4", answer: "6"},
		{question: "4+67", answer: "71"},
		{question: "2+34", answer: "36"},
		{question: "13+3", answer: "16"},
		{question: "12+4", answer: "16"},
		{question: "8+3", answer: "12"},
		{question: "24+4", answer: "29"},
		{question: "2+43", answer: "45"},
		{question: "23+4", answer: "27"},
	}, nil
}

func (m *ValidMockQuizDeps) play(q []Question, limit int) (Result, error) {
	return Result{
		correct:   5,
		attempted: 13,
	}, nil
}

func (m *ValidMockQuizDeps) showResult(r Result) {}

func (m *ValidMockQuizDeps) showHelp() {}

// Invalid mock dependencies
type InvalidMockQuizDeps struct{}

func (m *InvalidMockQuizDeps) parseCommand() (file string, limit int, help bool, err error) {
	return file, limit, help, err
}

func (m *InvalidMockQuizDeps) getQuestions(filepath string) ([]Question, error) {
	if filepath == "invalid" {
		return []Question{}, nil
	}

	return []Question{}, nil
}

func (m *InvalidMockQuizDeps) play(q []Question, limit int) (Result, error) {
	return Result{
		correct:   0,
		attempted: 0,
	}, nil
}

func (m *InvalidMockQuizDeps) showResult(r Result) {}

func (m *InvalidMockQuizDeps) showHelp() {}

func TestQuiz(t *testing.T) {
	testcases := []struct {
		name              string
		deps              QuizDependencies
		expectedCorrect   int
		expectedAttempted int
	}{
		{
			name:              "Valid Quiz",
			deps:              &ValidMockQuizDeps{},
			expectedCorrect:   5,
			expectedAttempted: 13,
		},
		{
			name:              "Invalid Quiz",
			deps:              &ValidMockQuizDeps{},
			expectedCorrect:   0,
			expectedAttempted: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			questions, err := tc.deps.getQuestions("filepath")

			if err != nil {
				log.Fatalf("\nCouldn't fetch the questions : %v\n", err)
			}

			res, err := tc.deps.play(questions, 1)

			if err != nil {
				log.Fatalf("\nCouldn't fetch the result : %v\n", err)
			}

			tc.deps.showResult(res)

			q := NewQuiz(tc.deps)
			StartQuiz(q)

		})
	}
}
