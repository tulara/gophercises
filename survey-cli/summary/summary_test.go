package summary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tulararogers-webster/gophercises/survey-cli/summary"
)

func Test_ShouldCalculateTotalParticipantCount(t *testing.T) {
	surveyResponses := []summary.SurveyResponse{
		{
			Email:       "employee1@abc.xyz",
			EmployeeId:  1,
			SubmittedAt: "2014-07-28T20:35:41+00:00",
			Responses:   []int{5, 5, 5, 4, 4},
		},
		{
			Email:       "",
			EmployeeId:  2,
			SubmittedAt: "2014-07-29T07:05:41+00:00",
			Responses:   []int{4, 5, 5, 3, 3},
		},
	}

	count := summary.CountTotalParticipants(surveyResponses)
	assert.Equal(t, 2, count, "total count did not match expected")
}

func Test_ShouldHandleUnsubmittedSurveys(t *testing.T) {
	surveyResponses := []summary.SurveyResponse{
		{
			Email:       "employee1@abc.xyz",
			EmployeeId:  1,
			SubmittedAt: "2014-07-28T20:35:41+00:00",
			Responses:   []int{5, 5, 5, 4, 4},
		},
		{
			Email:       "",
			EmployeeId:  2,
			SubmittedAt: "2014-07-29T07:05:41+00:00",
			Responses:   []int{4, 5, 5, 3, 3},
		},
		{
			Email:       "",
			EmployeeId:  3,
			SubmittedAt: "",
			Responses:   []int{5, 5, 5, 5, 4},
		},
	}

	count := summary.CountTotalParticipants(surveyResponses)
	assert.Equal(t, 2, count, "total count did not match expected")
}

func TestParticipationPercentage_AllParticipated(t *testing.T) {
	responses := []summary.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: "2014-07-28T20:35:41+00:00"},
		{EmployeeId: 2, SubmittedAt: "2014-07-29T07:05:41+00:00"},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 100.0, pct)
}

func TestParticipationPercentage_NoneParticipated(t *testing.T) {
	responses := []summary.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: ""},
		{EmployeeId: 2, SubmittedAt: ""},
		{EmployeeId: 3, SubmittedAt: ""},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 0.0, pct)
}

func TestParticipationPercentage_PartialParticipation(t *testing.T) {
	responses := []summary.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: "2014-07-28T20:35:41+00:00"},
		{EmployeeId: 2, SubmittedAt: "2014-07-29T07:05:41+00:00"},
		{EmployeeId: 3, SubmittedAt: ""},
		{EmployeeId: 4, SubmittedAt: ""},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 50.0, pct)
}

func TestParticipationPercentage_FractionalResult(t *testing.T) {
	responses := []summary.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: "2014-07-28T20:35:41+00:00"},
		{EmployeeId: 2, SubmittedAt: ""},
		{EmployeeId: 3, SubmittedAt: ""},
	}

	pct := summary.ParticipationPercentage(responses)
	assert.Equal(t, 33.33, pct)
}

func TestParticipationPercentage_EmptyInput(t *testing.T) {
	pct := summary.ParticipationPercentage([]summary.SurveyResponse{})

	// no division by zero — zero total means zero participation
	assert.Equal(t, 0.0, pct)
}
