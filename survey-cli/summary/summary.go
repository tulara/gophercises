package summary

type SurveyResponse struct {
	Email       string
	EmployeeId  int
	SubmittedAt string
	Responses   []int
}

// will need to model different types of responses
// type Response struct {
// 	Type  string
// 	Value int
// }

//

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
