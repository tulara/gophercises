package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

// do i need to export this?
type problem struct {
	question string
	answer string
} 

// does this need to be exported if it is in the same package? seems like no?
func runQuiz(problems []problem, handle io.Reader) (int, int){

	scanner := bufio.NewScanner(handle)
	var correctanswers, incorrectanswers int

	for i, problem := range problems {
		fmt.Printf("Question %d: %s:", i+1, problem.question)
		
		var useranswer string
		if scanner.Scan() {
			useranswer = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic (err)
		}

		if(useranswer == problem.answer) {
			correctanswers++
		} else {
			incorrectanswers++
		}
	}
	return correctanswers, incorrectanswers
}

func compileProblemsFromFile(filename string) [][]string {
	fmt.Println("Compiling quiz questions from", filename)
	file, err := os.Open(filename)
	// should return this 
	if err != nil { 
		fmt.Println("There was an issue opening your file", err)
		os.Exit(1)
	}
	defer file.Close()

	filereader := csv.NewReader(file)
	records, err := filereader.ReadAll() // assumes entire file will fit in memory
	if err != nil {
		fmt.Println("oh noes:", err)
	}

	return records
}

func parseProblems(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			question: record[0],
			answer: record[1],
		}
	}

	return problems
}


func main(){
	filenamePtr := flag.String("filename", "problems.csv", "point me to your problems")
	flag.Parse()

	records := compileProblemsFromFile(*filenamePtr)
	
	problems := parseProblems(records)
	correctanswers, incorrectanswers := runQuiz(problems, os.Stdin)
	fmt.Printf("Total correct: %d, Total incorrect: %d", correctanswers, incorrectanswers)
}