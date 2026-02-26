package summary_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tulararogers-webster/gophercises/survey-cli/domain"
	"github.com/tulararogers-webster/gophercises/survey-cli/summary"
)

func Test_ShouldCalculateTotalParticipantCount(t *testing.T) {
	surveyResponses := []domain.SurveyResponse{
		{
			Email:       "employee1@abc.xyz",
			EmployeeId:  1,
			SubmittedAt: mustParseTime(t, "2014-07-28T20:35:41+00:00"),
			Responses:   []int{5, 5, 5, 4, 4},
		},
		{
			Email:       "",
			EmployeeId:  2,
			SubmittedAt: mustParseTime(t, "2014-07-29T07:05:41+00:00"),
			Responses:   []int{4, 5, 5, 3, 3},
		},
	}

	count := summary.CountTotalParticipants(surveyResponses)
	assert.Equal(t, 2, count, "total count did not match expected")
}

func Test_ShouldHandleUnsubmittedSurveys(t *testing.T) {
	surveyResponses := []domain.SurveyResponse{
		{
			Email:       "employee1@abc.xyz",
			EmployeeId:  1,
			SubmittedAt: mustParseTime(t, "2014-07-28T20:35:41+00:00"),
			Responses:   []int{5, 5, 5, 4, 4},
		},
		{
			Email:       "",
			EmployeeId:  2,
			SubmittedAt: mustParseTime(t, "2014-07-29T07:05:41+00:00"),
			Responses:   []int{4, 5, 5, 3, 3},
		},
		{
			Email:       "",
			EmployeeId:  3,
			SubmittedAt: time.Time{},
			Responses:   []int{5, 5, 5, 5, 4},
		},
	}

	count := summary.CountTotalParticipants(surveyResponses)
	assert.Equal(t, 2, count, "total count did not match expected")
}

func TestParticipationPercentage_AllParticipated(t *testing.T) {
	responses := []domain.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: mustParseTime(t, "2014-07-28T20:35:41+00:00")},
		{EmployeeId: 2, SubmittedAt: mustParseTime(t, "2014-07-29T07:05:41+00:00")},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 100.0, pct)
}

func TestParticipationPercentage_NoneParticipated(t *testing.T) {
	responses := []domain.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: time.Time{}},
		{EmployeeId: 2, SubmittedAt: time.Time{}},
		{EmployeeId: 3, SubmittedAt: time.Time{}},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 0.0, pct)
}

func TestParticipationPercentage_PartialParticipation(t *testing.T) {
	responses := []domain.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: mustParseTime(t, "2014-07-28T20:35:41+00:00")},
		{EmployeeId: 2, SubmittedAt: mustParseTime(t, "2014-07-29T07:05:41+00:00")},
		{EmployeeId: 3, SubmittedAt: time.Time{}},
		{EmployeeId: 4, SubmittedAt: time.Time{}},
	}

	pct := summary.ParticipationPercentage(responses)

	assert.Equal(t, 50.0, pct)
}

func TestParticipationPercentage_FractionalResult(t *testing.T) {
	responses := []domain.SurveyResponse{
		{EmployeeId: 1, SubmittedAt: mustParseTime(t, "2014-07-28T20:35:41+00:00")},
		{EmployeeId: 2, SubmittedAt: time.Time{}},
		{EmployeeId: 3, SubmittedAt: time.Time{}},
	}

	pct := summary.ParticipationPercentage(responses)
	assert.Equal(t, 33.33, pct)
}

func TestParticipationPercentage_EmptyInput(t *testing.T) {
	pct := summary.ParticipationPercentage([]domain.SurveyResponse{})

	// no division by zero — zero total means zero participation
	assert.Equal(t, 0.0, pct)
}

func mustParseTime(t *testing.T, s string) time.Time {
	time, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("invalid time: %v", err)
	}
	return time
}
