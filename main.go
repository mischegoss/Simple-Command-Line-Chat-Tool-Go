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
	fmt.Println("Meet Rud-E, the AI with attitude")
}

func askQuestion() string {
	fmt.Print("So, what is your question already? ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func generateResponse(question string) string {
	response, err := client.Chat(context.TODO(), &cohere.ChatRequest{Message: question + " Make the response slightly rude and snarky."})
	if err != nil {
		fmt.Println("Error generating response:", err)
		fmt.Println("Error occurred. Please try again.")
	}
	return response.Text
}

func printResponse(question, response string) {
	fmt.Println("Here's the response to the question:", question)
	fmt.Println(response)
}

func askForMoreQuestions() bool {
	fmt.Print("Do you have another question? (y/n) ")
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
			fmt.Println("Finally some peach and quiet. You can exit now.")
			break
		}
	}
}
