package summary

import (
	"math"

	"github.com/tulararogers-webster/gophercises/survey-cli/domain"
)

// will need to model different types of responses
// type Response struct {
// 	Type  string
// 	Value int
// }

// ParticipationPercentage calculates the ratio of users who submitted a survey (participants)
// to all users (including those who have unsubmitted surveys).
// Returns percentage to two decimal places
func ParticipationPercentage(responses []domain.SurveyResponse) float64 {
	if len(responses) == 0 {
		return 0.0
	}
	raw := float64(CountTotalParticipants(responses)) / float64(len(responses)) * 100
	return math.Round(raw*100) / 100
}

func CountTotalParticipants(responses []domain.SurveyResponse) int {
	count := 0
	for _, r := range responses {
		if !r.SubmittedAt.IsZero() {
			count = count + 1
		}
	}

	return count
}
