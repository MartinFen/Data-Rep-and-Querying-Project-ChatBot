/*	Author: Martin Fennell
	Date: 2017
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Print a greeting to the user.
	fmt.Println("Hello, I'm Eliza. How are you feeling today?")

	// Keep reading user input and printing Eliza's response until the user types 'quit'.
	for reader := bufio.NewReader(os.Stdin); ; {
		// Print user prompt.
		fmt.Print("> ")
		// Read user input.
		userinput, _ := reader.ReadString('\n')
		// Trim the user input's end of line characters.
		userinput = strings.Trim(userinput, "\r\n")
		//hard coded output
		output := "Nice to meet you"

		// Generate and print Eliza's response.
		fmt.Println(output)

		// If the user input was quit, then quit.
		// Note that Eliza gets to respond to quit before this happens.
		if strings.Compare(strings.ToLower(strings.TrimSpace(userinput)), "quit") == 0 {
			fmt.Println("Bye.")
			break
		}
	}
}
