package main

import (
	"log"

	"github.com/supriya-kotturu/algorithms-in-go/quiz-game/utils"
)

func main() {
	filepath, limit, help, err := utils.ParseCommand()

	if err != nil {
		log.Fatalf("could not parse command : %v", err)
	}

	if help {
		utils.ShowHelp()
		return
	}

	questions, err := utils.GetQuestions(filepath)

	if err != nil {
		log.Fatalf("could not get questions : %v", err)
	}

	result, err := utils.Play(questions, limit)

	if err != nil {
		log.Fatalf("could not play the game : %v", err)
	}

	utils.ShowResult(result)
}
