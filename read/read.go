package read

import (
	"fmt"
	"log"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {
	if !db.UserExists(user) {
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		fmt.Println(message)
	}
}
