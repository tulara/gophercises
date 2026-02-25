package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"

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

	return responses, nil
}

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
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		return summary.SurveyResponse{}, fmt.Errorf("invalid employee_id %q: %w", row[1], err)
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

	// could add extra validation for valid email and timestamp.
	return summary.SurveyResponse{
		Email:       email,
		EmployeeId:  employeeID,
		SubmittedAt: submittedAt,
		Responses:   parsedResponses,
	}, nil
}
