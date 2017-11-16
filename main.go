package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Response struct {
	re        *regexp.Regexp
	responses []string
}

func buildResponseList() []Response {
	//empty struct
	response := Response{}
	//empty array of structs
	allResponses := []Response{}
	//read in a file
	file, _ := os.Open("./data/patterns.dat")
	//create scanner
	scanner := bufio.NewScanner(file)
	//loop through scanner until theres nothing left to scan
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		patternStr := scanner.Text()
		response.re = regexp.MustCompile(patternStr)
		scanner.Scan() // move onto the next line which holds the answers
		answersAsStr := scanner.Text()
		//split responses up when a ; is in scope of scanner
		answerList := strings.Split(answersAsStr, ";")
		//add split responses
		response.responses = answerList
		allResponses = append(allResponses, response)
	}
	//fmt.Println(allResponses)
	return allResponses //return list
}

//main func where it loops until the user quits

func main() {
	fmt.Println("Hi my name is Eliza, whats yours?")
	for reader := bufio.NewReader(os.Stdin); ; {
		// Print user prompt.
		fmt.Print("> ")
		// Read user input.
		userInput, _ := reader.ReadString('\n')
		// Trim the user input's end of line characters.
		userInput = strings.Trim(userInput, "\r\n")
		fmt.Println(buildResponseList())
		//Quit program
		if strings.Compare(strings.ToLower(strings.TrimSpace(userInput)), "quit") == 0 {
			fmt.Println("Bye.")
			break
		}
	}
}
