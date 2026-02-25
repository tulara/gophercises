package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"strconv"
	"time"

	"github.com/tulararogers-webster/gophercises/survey-cli/summary"
)

func ParseResponses(r io.Reader) ([]summary.SurveyResponse, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 0 // use the first row to determine how many fields each row should have

	// read by row if we expect csv file to be bigger than memory
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("Survey response input was empty")
	}

	var responses []summary.SurveyResponse
	for i, row := range rows {
		resp, err := parseSurveyResponseRow(row)
		if err != nil {
			// could feasibly attempt to parse rest of rows but this is simpler for now
			return nil, fmt.Errorf("[row: %d] %w", i+1, err)
		}
		responses = append(responses, resp)
	}

	return deduplicate(responses), nil
}

func deduplicate(responses []summary.SurveyResponse) []summary.SurveyResponse {
	byEmail := map[string]int{}
	byID := map[int]int{}
	var result []summary.SurveyResponse

	for _, r := range responses {
		existingIdx := -1
		if r.Email != "" {
			if idx, ok := byEmail[r.Email]; ok {
				existingIdx = idx
			}
		}
		if existingIdx == -1 && r.EmployeeId != 0 {
			if idx, ok := byID[r.EmployeeId]; ok {
				existingIdx = idx
			}
		}

		if existingIdx == -1 {
			idx := len(result)
			result = append(result, r)
			if r.Email != "" {
				byEmail[r.Email] = idx
			}
			if r.EmployeeId != 0 {
				byID[r.EmployeeId] = idx
			}
		} else {
			result[existingIdx] = wasSubmittedLater(result[existingIdx], r)
		}
	}

	return result
}

func wasSubmittedLater(existing, incoming summary.SurveyResponse) summary.SurveyResponse {
	if existing.SubmittedAt != "" {
		// discard unsubmitted surveys if there is a submitted one.
		if incoming.SubmittedAt == "" {
			return existing
		}

		// we know times will be valid because invalid times will have been discarded during parsing.
		et, _ := time.Parse(time.RFC3339, existing.SubmittedAt)
		it, _ := time.Parse(time.RFC3339, incoming.SubmittedAt)

		if et.After(it) {
			return existing
		}
	}
	return incoming
}

// Ensures all parsed survey rows are valid
// Errors abort parsing completely. A more graceful solution would be to skip the row and continue with the rest of the report.
func parseSurveyResponseRow(row []string) (summary.SurveyResponse, error) {
	if len(row) < 4 {
		return summary.SurveyResponse{}, fmt.Errorf("expected at least 4 columns, but got %d", len(row))
	}

	email := row[0]
	id := row[1]
	submittedAt := row[2]
	responses := row[3:]

	// if we don't have employeeid or email, we cant count distinct participants
	if email == "" && id == "" {
		return summary.SurveyResponse{}, errors.New("row must have employee id or email")

	}

	employeeID := 0
	var err error
	if id != "" {
		employeeID, err = strconv.Atoi(id)
		if err != nil {
			return summary.SurveyResponse{}, fmt.Errorf("invalid employee_id %q: %w", id, err)
		}
	}

	if email != "" {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return summary.SurveyResponse{}, fmt.Errorf("invalid email %q: %w", email, err)
		}
	}

	if submittedAt != "" {
		_, err := time.Parse(time.RFC3339, submittedAt)
		if err != nil {
			return summary.SurveyResponse{}, fmt.Errorf("invalid submittedAt timestamp %q: %w", submittedAt, err)
		}
	}

	var parsedResponses []int
	for j, r := range responses {
		if r == "" {
			parsedResponses = append(parsedResponses, 0)
			continue
		}

		val, err := strconv.Atoi(r)
		if err != nil {
			// again, it would more appropriate to skip this column rather than abort completely.
			return summary.SurveyResponse{}, fmt.Errorf("[column: %d] %q was not a number. Error: %w", j+4, r, err)
		}
		parsedResponses = append(parsedResponses, val)
	}

	return summary.SurveyResponse{
		Email:       email,
		EmployeeId:  employeeID,
		SubmittedAt: submittedAt,
		Responses:   parsedResponses,
	}, nil
}
