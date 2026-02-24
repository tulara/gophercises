package main

// input:
// first argument is name of csv file, second argument is responses.
// number of questions tells us how many columns to expect in responses.
// type of question tells us type to expect and validation to run

// responses include columns:
// Email
// Employee Id
// Submitted At Timestamp (if there is no submitted at timestamp, you can assume the user did not submit a survey)
// Each column from the fourth onwards are responses to survey questions.

// from this, make a survey summary containing:
// 1.The participation percentage and total participant counts of the survey.
// 2. The average for each rating question

// main will include reading from files.
// pachage delegation for validation and aggregation
