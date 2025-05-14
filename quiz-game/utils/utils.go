package utils

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	id       int
	question string
	answer   string
}

type Result struct {
	totalQuestions int
	correct        int
	attempted      int
}

func showHelp() {
	fmt.Println("--- QUIZ GAME HELP ---")
	fmt.Println("This program runs a quiz using questions from a CSV file.")
	fmt.Println("\nOptions:")
	fmt.Println("  -file string")
	fmt.Println("        Path to the question bank file in CSV format (default \"problem-set.csv\")")
	fmt.Println("  -help")
	fmt.Println("        Display this help information")
	fmt.Println("\nExample usage:")
	fmt.Println("  quiz-game -file=my-questions.csv")
}

func parseCommand() (string, int, bool, error) {
	file := flag.String("file", "problem-set.csv", "The file path to the question bank in csv format of (question, answer).")
	limit := flag.Int("limit", 30, "The time limit in seconds to solve each question.")
	help := flag.Bool("help", false, "Shows the instructions to start the quiz.")

	flag.Parse()

	_, err := os.Stat(*file)

	if err == nil {
		return *file, *limit, *help, nil
	}

	if os.IsNotExist(err) {
		return "", *limit, false, errors.New("file not found: " + *file + "")
	} else {
		return "", *limit, false, err
	}
}

func getQuestions(filePath string) ([]Question, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return []Question{}, errors.New("could not open file: " + file.Name())
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		return []Question{}, errors.New("could not parse the file in CSV format: " + file.Name())
	}

	questions := make([]Question, len(records))

	for idx, record := range records {
		questions[idx] = Question{
			id:       idx,
			question: strings.TrimSpace(record[0]),
			answer:   strings.TrimSpace(record[1]),
		}
	}

	return questions, nil
}

func play(questions []Question, limit int) (Result, error) {
	var result Result

	result.totalQuestions = len(questions)

	timer := time.NewTimer(time.Duration(limit) * time.Second)
	ansCh := make(chan string)

	for _, q := range questions {
		fmt.Println(q.question)

		// I/O is a blocking call. Wrapping it in a go routine makes it asynchronous
		go func() {
			var ans string

			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time limit exceeded!\n")
			return result, nil
		case ans := <-ansCh:
			if ans == q.answer {
				result.correct++
			}
			result.attempted++
		}
	}

	return result, nil
}

func showResult(r Result) {
	fmt.Println("--- RESULT ---")
	fmt.Printf("You attempted %d out of %d\n", r.attempted, r.totalQuestions)
	fmt.Printf("Your score : %d / %d\n", r.correct, r.totalQuestions)
}
