package domain

type SurveyResponse struct {
	Email string
	// EmployeeId is parsed to 0 when empty.
	EmployeeId  int
	SubmittedAt string

	// Responses are parsed to 0 when unanswered
	Responses []int
}
