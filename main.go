package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
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

//function for generating a random number based on the size of the list passed in
//adapted from problem sheets 1
func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers))
	return answers[index]
}

//function for taking in user input and crafting a response based on what the userInput is
func Ask(userInput string) string {
	responses := buildResponseList()
	for _, resp := range responses { // look at every single response/pattern/answers
		if resp.re.MatchString(userInput) {

			Answer := getRandomAnswer(resp.responses) // get random element.
			return Answer
		}
	}
	// if we're down here, it means there were no matches;
	return "Im sorry I don't follow what you mean, could you clarify that for me?" // catch all.
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
		fmt.Println(Ask(userInput))
		//Quit program
		if strings.Compare(strings.ToLower(strings.TrimSpace(userInput)), "quit") == 0 {
			fmt.Println("Bye.")
			break
		}
	}
}
