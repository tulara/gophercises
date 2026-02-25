package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tulararogers-webster/gophercises/survey-cli/parser"
	"github.com/tulararogers-webster/gophercises/survey-cli/summary"
)

func TestParseResponses_ValidInput(t *testing.T) {
	input := strings.NewReader(
		"employee1@abc.xyz,1,2014-07-28T20:35:41+00:00,5,4,3\n" +
			",2,2014-07-29T07:05:41+00:00,4,5,3",
	)

	responses, err := parser.ParseResponses(input)

	require.NoError(t, err)
	assert.Len(t, responses, 2)
	assert.Equal(t, summary.SurveyResponse{
		Email:       "employee1@abc.xyz",
		EmployeeId:  1,
		SubmittedAt: "2014-07-28T20:35:41+00:00",
		Responses:   []int{5, 4, 3},
	}, responses[0])
	assert.Equal(t, summary.SurveyResponse{
		Email:       "",
		EmployeeId:  2,
		SubmittedAt: "2014-07-29T07:05:41+00:00",
		Responses:   []int{4, 5, 3},
	}, responses[1])
}

func TestParseResponses_EmptyInput(t *testing.T) {
	input := strings.NewReader("")

	responses, err := parser.ParseResponses(input)

	assert.Nil(t, responses)
	assert.EqualError(t, err, "Survey response input was empty")
}

func TestParseResponses_EmptyResponseColumnsDefaultToZero(t *testing.T) {
	input := strings.NewReader("employee@abc.xyz,1,2014-07-28T20:35:41+00:00,,5,")

	responses, err := parser.ParseResponses(input)

	require.NoError(t, err)
	assert.Equal(t, []int{0, 5, 0}, responses[0].Responses)
}

func TestParseResponses_MissingBothEmailAndID(t *testing.T) {
	input := strings.NewReader(",,2014-07-28T20:35:41+00:00,5,4")

	_, err := parser.ParseResponses(input)

	assert.EqualError(t, err, "[row: 1] row must have employee id or email")
}

func TestParseResponses_EmptyEmployeeIdDefaultsToZero(t *testing.T) {
	// email is present but id is empty — passes the identity check but fails int parsing
	input := strings.NewReader("employee@abc.xyz,,2014-07-28T20:35:41+00:00,5")

	responses, err := parser.ParseResponses(input)

	require.NoError(t, err)
	require.Len(t, responses, 1)
	assert.Equal(t, 0, responses[0].EmployeeId)
}

func TestParseResponses_InvalidEmployeeID(t *testing.T) {
	input := strings.NewReader("employee@abc.xyz,not-a-number,2014-07-28T20:35:41+00:00,5")

	_, err := parser.ParseResponses(input)

	assert.ErrorContains(t, err, "[row: 1]")
	assert.ErrorContains(t, err, `invalid employee_id "not-a-number"`)
}

func TestParseResponses_InvalidResponseValue(t *testing.T) {
	// "not-a-number" is the 5th column (j=1 → j+4=5)
	input := strings.NewReader("employee@abc.xyz,1,2014-07-28T20:35:41+00:00,5,not-a-number,3")

	_, err := parser.ParseResponses(input)

	assert.ErrorContains(t, err, "[row: 1]")
	assert.ErrorContains(t, err, "[column: 5]")
	assert.ErrorContains(t, err, `"not-a-number" was not a number`)
}

func TestParseResponses_TooFewColumns(t *testing.T) {
	// minimum is 4 columns: email, id, submitted_at, and at least one response
	input := strings.NewReader("employee@abc.xyz,1,2014-07-28T20:35:41+00:00")

	_, err := parser.ParseResponses(input)

	assert.ErrorContains(t, err, "[row: 1]")
	assert.ErrorContains(t, err, "expected at least 4 columns, but got 3")
}
