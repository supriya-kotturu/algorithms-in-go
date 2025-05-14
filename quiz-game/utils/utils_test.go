package utils

import (
	"testing"
)

// MockQuizDeps implements QuizDependencies interface for testing
type MockQuizDeps struct {
	ParseCommandCalled bool
	GetQuestionsCalled bool
	PlayCalled         bool
	ShowHelpCalled     bool
	ShowResultCalled   bool

	// Return values for mocked functions
	FilePath          string
	TimeLimit         int
	HelpFlag          bool
	ParseCommandError error
	Questions         []Question
	GetQuestionsError error
	Result            Result
	PlayError         error
}

func (m *MockQuizDeps) parseCommand() (string, int, bool, error) {
	m.ParseCommandCalled = true
	return m.FilePath, m.TimeLimit, m.HelpFlag, m.ParseCommandError
}

func (m *MockQuizDeps) getQuestions(filepath string) ([]Question, error) {
	m.GetQuestionsCalled = true
	return m.Questions, m.GetQuestionsError
}

func (m *MockQuizDeps) play(q []Question, limit int) (Result, error) {
	m.PlayCalled = true
	return m.Result, m.PlayError
}

func (m *MockQuizDeps) showResult(r Result) {
	m.ShowResultCalled = true
}

func (m *MockQuizDeps) showHelp() {
	m.ShowHelpCalled = true
}

func TestQuizNormalFlow(t *testing.T) {
	// Setup mock with expected values
	mockDeps := &MockQuizDeps{
		FilePath:  "test.csv",
		TimeLimit: 30,
		HelpFlag:  false,
		Questions: []Question{
			{id: 1, question: "1+1?", answer: "2"},
			{id: 2, question: "2+2?", answer: "4"},
		},
		Result: Result{
			totalQuestions: 2,
			correct:        1,
			attempted:      2,
		},
	}

	// Create quiz with mock dependencies
	quiz := NewQuiz(mockDeps)

	// Run the quiz
	StartQuiz(quiz)

	// Verify all expected functions were called
	if !mockDeps.ParseCommandCalled {
		t.Error("parseCommand was not called")
	}
	if !mockDeps.GetQuestionsCalled {
		t.Error("getQuestions was not called")
	}
	if !mockDeps.PlayCalled {
		t.Error("play was not called")
	}
	if !mockDeps.ShowResultCalled {
		t.Error("showResult was not called")
	}
	if mockDeps.ShowHelpCalled {
		t.Error("showHelp was called unexpectedly")
	}
}

func TestQuizHelpFlow(t *testing.T) {
	// Setup mock with help flag set to true
	mockDeps := &MockQuizDeps{
		FilePath:  "test.csv",
		TimeLimit: 30,
		HelpFlag:  true,
	}

	// Create quiz with mock dependencies
	quiz := NewQuiz(mockDeps)

	// Run the quiz
	StartQuiz(quiz)

	// Verify only parseCommand and showHelp were called
	if !mockDeps.ParseCommandCalled {
		t.Error("parseCommand was not called")
	}
	if !mockDeps.ShowHelpCalled {
		t.Error("showHelp was not called")
	}
	if mockDeps.GetQuestionsCalled {
		t.Error("getQuestions was called unexpectedly")
	}
	if mockDeps.PlayCalled {
		t.Error("play was called unexpectedly")
	}
	if mockDeps.ShowResultCalled {
		t.Error("showResult was called unexpectedly")
	}
}
