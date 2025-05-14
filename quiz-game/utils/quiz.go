package utils

import "log"

type QuizDependencies interface {
	parseCommand() (string, int, bool, error)
	getQuestions(string) ([]Question, error)
	play([]Question, int) (Result, error)
	showResult(Result)
	showHelp()
}

type Quiz struct {
	deps QuizDependencies
}

type DefaultQuizDeps struct{}

func (d *DefaultQuizDeps) parseCommand() (string, int, bool, error) {
	return parseCommand()
}

func (d *DefaultQuizDeps) getQuestions(filepath string) ([]Question, error) {
	return getQuestions(filepath)
}

func (d *DefaultQuizDeps) play(questions []Question, limit int) (Result, error) {
	return play(questions, limit)
}

func (d *DefaultQuizDeps) showResult(r Result) {
	showResult(r)
}

func (d *DefaultQuizDeps) showHelp() {
	showHelp()
}

func NewQuiz(deps QuizDependencies) *Quiz {

	if deps == nil {
		return &Quiz{
			deps: &DefaultQuizDeps{},
		}
	}

	return &Quiz{
		deps: deps,
	}
}

func StartQuiz(q *Quiz) {
	filepath, limit, help, err := q.deps.parseCommand()

	if err != nil {
		log.Fatalf("could not parse command : %v", err)
	}

	if help {
		q.deps.showHelp()
		return
	}

	questions, err := q.deps.getQuestions(filepath)

	if err != nil {
		log.Fatalf("could not get questions : %v", err)
	}

	result, err := q.deps.play(questions, limit)

	if err != nil {
		log.Fatalf("could not play the game : %v", err)
	}

	q.deps.showResult(result)

}
