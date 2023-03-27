package read

import (
	"fmt"
	"log"

	"github.com/lanimatran/deaddrop-go/db"
	"github.com/lanimatran/deaddrop-go/logging"
	"github.com/lanimatran/deaddrop-go/session"
	"strings"
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

	messages, senders, macs := db.GetMessagesForUser(user)
	for i, message := range messages {
		decryptedMessage := strings.TrimSpace(session.Decrypt(message))
		mac := session.ProduceUnhashedMAC(senders[i], decryptedMessage)
		err := session.CompareMACs(macs[i], mac)
		if (err != nil) {
			fmt.Println("A message is hidden due to not passing integrity check")
			logging.LogMessage(user, "A read message failed integrity check")
		} else {
			fmt.Println("From: " + senders[i])
			fmt.Println(decryptedMessage)
		}
		fmt.Println("==== End of Message ====")
	}

	logging.LogMessage(user, "Succesfully read messages")
}
