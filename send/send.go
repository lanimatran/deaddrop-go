package send

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lanimatran/deaddrop-go/db"
	"github.com/lanimatran/deaddrop-go/logging"
	"github.com/lanimatran/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to string) {
	if !db.UserExists(to) {
		logging.LogMessage(to, "Failed to receive a message")
		log.Fatalf("Destination user does not exist")
	}

	message := getUserMessage()

	db.SaveMessage(session.Encrypt(message), to)

	logging.LogMessage(to, "Succesfully received a message")
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {
	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return text
}
