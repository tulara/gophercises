package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main(){
	file, err := os.Open("problems.csv")
	if err != nil { fmt.Println("There was an issue opening your file", err) }
	defer file.Close()

	filereader := csv.NewReader(file)
	inputreader := bufio.NewReader(os.Stdin)
	var correctanswers, incorrectanswers int

	records, err := filereader.ReadAll() // assumes entire file will fit in memory
	if err != nil {
		fmt.Println("oh noes:", err)
	}

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