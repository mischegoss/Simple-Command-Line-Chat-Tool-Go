package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

var client = cohereclient.NewClient(cohereclient.WithToken("prCT9e0FiD1a79YtwnnHGJwGd0SlerSikMSRgmYo"))
var reader = bufio.NewReader(os.Stdin)
var question (string)
var anotherQuestion (string)

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
			return // Handle the error gracefully
		}

		anotherQuestion = strings.ToLower(strings.TrimSpace(anotherQuestion))

		if anotherQuestion == "y" {
			interact()
		} else {
			fmt.Println("Goodbye.")
			break
		}
	}
}

func interact() {
	grabQuestion()
	printResponse()
}

func main() {
	welcomeUser()
	interact()
	moreQuestion()
}
