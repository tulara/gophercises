package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scan is just testing out how bufio scanning works
func Scan() {
	scanner := bufio.NewScanner(os.Stdin)
	var answer string

	if scanner.Scan() {
		answer = scanner.Text()
	} 
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("echo:", []byte(answer))
}