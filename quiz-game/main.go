package main

import (
	"github.com/supriya-kotturu/algorithms-in-go/quiz-game/utils"
)

func main() {
	deps := &utils.DefaultQuizDeps{}
	q := utils.NewQuiz(deps)
	utils.StartQuiz(q)
}
