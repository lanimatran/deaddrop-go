package send

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lanimatran/deaddrop-go/db"
	"github.com/lanimatran/deaddrop-go/logging"
	"github.com/lanimatran/deaddrop-go/session"
	"strings"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(user string, to string) {
	if !db.UserExists(user) {
		logging.LogMessage(user, "Failed to send messages - User not recognized")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		logging.LogMessage(user, "Could not authenticate to send messages")
		log.Fatalf("Unable to authenticate user")
	}

	if !db.UserExists(to) {
		logging.LogMessage(to, "Failed to receive a message")
		log.Fatalf("Destination user does not exist")
	}

	message := strings.TrimSpace(getUserMessage())
	mac,_ := session.ProduceMAC(user, message)

	fmt.Println(session.ProduceUnhashedMAC(user, message))

	db.SaveMessage(session.Encrypt(message), user, to, mac)

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
