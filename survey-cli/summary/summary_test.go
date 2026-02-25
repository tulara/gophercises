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
