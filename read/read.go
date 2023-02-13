package read

import (
	"fmt"
	"log"

	"github.com/lanimatran/deaddrop-go/db"
	"github.com/lanimatran/deaddrop-go/logging"
	"github.com/lanimatran/deaddrop-go/session"
)

func ReadMessages(user string) {
	if !db.UserExists(user) {
		logging.LogMessage(user, "Failed to read messages")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		logging.LogMessage(user, "Could not authenticate to read messages")
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		fmt.Println(session.Decrypt(message))
	}

	logging.LogMessage(user, "Succesfully read messages")
}
