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
var question (string)
var anotherQuestion (string)
var apiKey = os.Getenv("cohere_key")
var client = cohereclient.NewClient(cohereclient.WithToken(apiKey))

func welcomeUser() {
	fmt.Println("Welcome to Ask Me Anything - AI Edition")
}

func grabQuestion() string {
	fmt.Print("What is your question? ")
	input, _ := reader.ReadString('\n')
	question = input
	return strings.TrimSpace(question)
}

func printResponse() {
	response, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: question,
		},
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Here is the response to the question ", question)
	fmt.Println(response.Text)
}

func moreQuestion() {
	for {
		fmt.Print("Do you have another question? (y/n) ")
		anotherQuestion, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		anotherQuestion = strings.ToLower(strings.TrimSpace(anotherQuestion))

		if anotherQuestion == "y" {
			grabQuestion()
			printResponse()
		} else {
			fmt.Println("Goodbye.")
			break
		}
	}
}

func main() {
	welcomeUser()
	grabQuestion()
	printResponse()
	moreQuestion()
}
