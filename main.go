package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	_ "github.com/joho/godotenv/autoload"
)

var reader = bufio.NewReader(os.Stdin)
var apiKey = os.Getenv("COHERE_KEY")
var client = cohereclient.NewClient(cohereclient.WithToken(apiKey))

func welcomeUser() {
	fmt.Println("Meet Arrgh-I, the AI pirate")
}

func askQuestion() string {
	fmt.Print("Ahoy there, matey! Spit out yer question! ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func generateResponse(question string) string {
	response, err := client.Chat(context.TODO(), &cohere.ChatRequest{Message: question + " Make the response slightly snarky and in the tone of a pirate."})
	if err != nil {
		fmt.Println("Error generating response:", err)
		fmt.Println("Error occurred. Please try again.")
	}
	return response.Text
}

func printResponse(question, response string) {
	fmt.Println("Blimey, here's the answer to the riddle you've been seeking:", question)
	fmt.Println(response)
}

func askForMoreQuestions() bool {
	fmt.Print("Arr, got another riddle for me, have ye?  (y/n) ")
	text, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(text)) == "y"
}

func main() {
	welcomeUser()
	for {
		question := askQuestion()
		response := generateResponse(question)
		printResponse(question, response)

		if !askForMoreQuestions() {
			fmt.Println("Be gone with ye, then! And watch out for sharks, they like the taste o' fools.")
			break
		}
	}
}
