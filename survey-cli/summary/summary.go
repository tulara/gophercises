package summary

type SurveyResponse struct {
	Email string
	// EmployeeId is parsed to 0 when empty.
	EmployeeId  int
	SubmittedAt string

	// Responses are parsed to 0 when unanswered
	Responses []int
}

// will need to model different types of responses
// type Response struct {
// 	Type  string
// 	Value int
// }

func CountTotalParticipants(responses []SurveyResponse) int {
	// make sure unique, here or before here
	count := 0
	for _, r := range responses {
		if r.SubmittedAt != "" {
			count = count + 1
		}
	}

	return count
}
