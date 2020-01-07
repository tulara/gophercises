package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func compileProblems(filename string) [][]string {
	fmt.Println("Compiling quiz questions from", filename)
	file, err := os.Open(filename)
	// should return this 
	if err != nil { fmt.Println("There was an issue opening your file", err) }
	defer file.Close()

	filereader := csv.NewReader(file)
	records, err := filereader.ReadAll() // assumes entire file will fit in memory
	if err != nil {
		fmt.Println("oh noes:", err)
	}

	return records
}

func main(){
	filenamePtr := flag.String("filename", "problems.csv", "point me to your problems")
	flag.Parse()

	records := compileProblems(*filenamePtr)
	inputreader := bufio.NewReader(os.Stdin)
	var correctanswers, incorrectanswers int

	for _, record := range records {
		fmt.Println(record[0],"?")
		useranswer, err := inputreader.ReadString('\n')
		if err != nil {panic(err)}

		if(strings.TrimRight(useranswer, "\r\n") == record[1]){
			correctanswers++ 
		} else {
			incorrectanswers++
		}
	}

	fmt.Printf("Total correct: %d, Total incorrect: %d", correctanswers, incorrectanswers)
}