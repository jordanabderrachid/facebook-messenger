package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jordanabderrachid/facebook-messenger/messenger"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	accessToken := getInput("Bot access token: ", r)
	recipientID := getInput("Recipient id: ", r)
	text := getInput("Text: ", r)

	m := messenger.NewMessenger(accessToken)
	response, err := m.SendTextMessage(recipientID, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(response)
}

func getInput(message string, r *bufio.Reader) string {
	fmt.Print(message)
	value, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.Replace(value, "\n", "", -1)
}
