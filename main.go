package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//empty map
var reflections map[string]string

//empty struct
type Response struct {
	re        *regexp.Regexp
	responses []string
}

//Func for reading in file contents and building a list of structs with each REGEX pattern its corrisponding responses.
//adapted from a group session
func buildResponseList() []Response {
	//empty struct
	response := Response{}
	//empty array of structs
	allResponses := []Response{}
	//read in a file
	file, err := os.Open("./data/patterns.dat")
	//theres an error crash
	if err != nil {
		panic(err)
	}
	//close file
	defer file.Close()
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

//Func for generating a randon number
//adapted from https://github.com/MartinFen/Go-Problem-Sheet-2-Answers/tree/master/Prob-5
func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers))
	return answers[index]
}

//Func for swapping certain words by spliting up input and swapping words and returning the string
//adapted from group session
func subWords(original string) string {
	//adapted from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
	if reflections == nil { // map hasn't been made yet
		reflections = map[string]string{ // will only happen once.
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "me",
			"me":     "you",
			"myself": "yourself",
		}
	}
	//split up string that user input into a string array then and loop through array such as you:me
	words := strings.Split(original, " ")
	for index, word := range words {
		//change the word if it's found in the map
		val, ok := reflections[word]
		if ok {
			//swap with the value
			words[index] = val
		}
	}
	return strings.Join(words, " ")
}

//Func for taking in user input and crafting elizas response
func Ask(userInput string) string {
	responses := buildResponseList()
	for _, resp := range responses { // look at every single response/answers
		if resp.re.MatchString(userInput) {
			match := resp.re.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]
			//fmt.Println(match[0])
			// remove punctuation here
			captured = subWords(captured)
			formatAnswer := getRandomAnswer(resp.responses) // get random element.
			if strings.Contains(formatAnswer, "%s") {       // string needs to be formatted
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer
		}
	}
	// if we're down here, it means there were no matches;
	return "Im sorry I don't follow what you mean, could you clarify that for me?" // catch all.
}

//Func for handling the user input from the webapp and sending the response back to the page
//adapted from adapted from https://github.com/data-representation/go-ajax
func userinputhandler(w http.ResponseWriter, r *http.Request) {

	userInput := strings.ToLower(r.URL.Query().Get("value"))
	fmt.Fprintf(w, "%s", Ask(userInput))
}

//adapted from adapted from https://github.com/data-representation/go-ajax
//Func for serving the SPA and linking the userinputhandler
func main() {
	//adapted from ajax
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
//Func was used for checking if the userinput and responses were working on the cli
/*func main() {
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
}*/
