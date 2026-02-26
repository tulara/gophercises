package domain

import "time"

type SurveyResponse struct {
	Email string
	// EmployeeId is parsed to 0 when empty.
	EmployeeId  int
	SubmittedAt time.Time

	// Responses are parsed to 0 when unanswered
	Responses []int
}
