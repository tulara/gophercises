package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer string
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

func runQuiz(problems []problem, reader bufio.Reader) (int, int){
	var correctanswers, incorrectanswers int

	for i, problem := range problems {
		fmt.Printf("Question %d: %s:", i+1, problem.question)
		useranswer, err := reader.ReadString('\n')
		if err != nil { panic(err) }

		if(strings.TrimSpace(useranswer) == problem.answer){
			correctanswers++
		} else {
			incorrectanswers++
		}
	}
	return correctanswers, incorrectanswers
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
	inputReader := bufio.NewReader(os.Stdin)
	
	problems := parseProblems(records)
	correctanswers, incorrectanswers := runQuiz(problems, *inputReader)
	fmt.Printf("Total correct: %d, Total incorrect: %d", correctanswers, incorrectanswers)
}