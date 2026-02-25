package main

import (
	"log"
	"os"

	"github.com/tulararogers-webster/gophercises/survey-cli/parser"
	"github.com/tulararogers-webster/gophercises/survey-cli/summary"
)

// input:
// first argument is name of csv file, second argument is responses.
// number of questions tells us how many columns to expect in responses.
// type of question tells us type to expect and validation to run

// responses include columns:
// Email
// Employee Id
// Submitted At Timestamp (if there is no submitted at timestamp, you can assume the user did not submit a survey)
// Each column from the fourth onwards are responses to survey questions.

// from this, make a survey summary containing:
// 1.The participation percentage and total participant counts of the survey.
// 2. The average for each rating question

// main will include reading from files.
// pachage delegation for validation and aggregation

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: survey-cli <responses.csv>")
	}

	//questionsFile := os.Args[1]
	responsesFile := os.Args[1]

	f, err := os.Open(responsesFile)
	if err != nil {
		log.Fatalf("Error opening %s:%v", responsesFile, err)
	}
	defer f.Close()

	responses, err := parser.ParseResponses(f)
	if err != nil {
		log.Fatalf("Error parsing responses: %v", err)
	}

	count := summary.CountTotalParticipants(responses)
	pct := summary.ParticipationPercentage(responses)
	log.Printf("Total participants: %d", count)
	log.Printf("Participation percentage: %.2f%%", pct)
}
