package main

//Imports the packages needed for the program to run including Cohere API client
import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	_ "github.com/joho/godotenv/autoload"
	chalk "github.com/ttacon/chalk"
)

// Grabs the API key from the .env file
var apiKey = os.Getenv("COHERE_KEY")

// Creates a new client object with the API key
var client = cohereclient.NewClient(cohereclient.WithToken(apiKey))

// Creates reader object that will grab users' input in CLI
var reader = bufio.NewReader(os.Stdin)

// Defines styles used with chalk
var yellowOnBlack = chalk.Yellow.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold).
	Style
var blackOnYellow = chalk.Black.NewStyle().WithBackground(chalk.Yellow).WithTextStyle(chalk.Bold).
	Style

// Prints welcome message
func welcomeUser() {
	fmt.Println(blackOnYellow("Meet Arrgh-I, the AI pirate"))
}

// Asks a question and returns a string value of one line and trims the text
func askQuestion() string {
	fmt.Print(yellowOnBlack("Ahoy there, matey! Spit out yer question! "))
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Makes call to the API and returns response object or error message
func generateResponse(question string) string {
	response, err := client.Chat(context.TODO(),
		&cohere.ChatRequest{Message: question + " Make the response slightly snarky and in the tone of a pirate."})
	if err != nil {
		fmt.Println("Error generating response:", err)
		fmt.Println("Error occurred. Please try again.")
	}
	return response.Text
}

// This prints the response text from the response object
func printResponse(question, responseText string) {
	fmt.Println(yellowOnBlack("Blimey, here's the answer to the riddle you've been seeking: " + question))
	fmt.Println(blackOnYellow(responseText))
}

// Asks question and returns true/false outcome 'y/n'
func askForMoreQuestions() bool {
	fmt.Print(yellowOnBlack("Arr, got another riddle for me, have ye? (y/n) "))
	text, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(text)) == "y"
}

// Executes the program
func main() {
	//Prints welcome message
	welcomeUser()
	// This for loop will keep going until ask for more questions is false
	for {
		question := askQuestion()
		responseText := generateResponse(question)
		printResponse(question, responseText)

		if !askForMoreQuestions() {
			fmt.Println(yellowOnBlack("Be gone with ye, then! And watch out for sharks, they like the taste o' land fools."))
			break
		}
	}
}
