package main

import (
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(3, 7)
	if total != 10 {
		t.Errorf("Sum totaled to incorrect value, expected: %d but actual: %d", 10, total)
	}
}

func TestRunQuiz(t *testing.T) {

	problems := []problem{ 
		{"2+2", "4"},
		{"3+4", "7"},
	}

	t.Run("RunQuiz should report full marks when all problems answered correctly",
	 func(t *testing.T) {
		 correct, incorrect := runQuiz(problems, strings.NewReader("4\n7\n"))

		 if correct != 2 {
			 t.Errorf("Expected %d correct answers but was %d", 2, correct)
		 }
	 
		 if incorrect != 0 {
			 t.Errorf("Expected %d incorrect answers but was %d", 0, incorrect)
		 }
	 })

	 t.Run("RunQuiz should report incorrect answers when they are given",
		func(t *testing.T){
			c, i := runQuiz(problems, strings.NewReader("3\n12\n"))

			expectedCorrect := 0
			expectedIncorrect := 2
			if c != expectedCorrect {
				t.Errorf("Expected %d correct answers but was %d", expectedCorrect, c)
			}

			if i != expectedIncorrect {
				t.Errorf("Expected %d incorrect answers but was %d", expectedIncorrect, i)
			}
	})
}