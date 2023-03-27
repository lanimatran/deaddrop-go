package db

import (
	"log"
)

// GetMessagesForUser assumes that a user has already been
// authenticated through a call to session.Authenticate(user)
// and then returns all the messages stored for that user
func GetMessagesForUser(user string) ([][]byte, []string, []string) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT data, user, checksum FROM Messages, Users
		WHERE recipient = (
			SELECT id FROM Users WHERE user = ?
		) AND sender = Users.id
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	// marshall rows into an array
	messages := make([][]byte, 0)
	senders := make([]string, 0)
	macs := make([]string, 0)
	for rows.Next() {
		var message []byte
		var sender string
		var mac string
		err := rows.Scan(&message, &sender, &mac)
		if err != nil {
			log.Fatalf("unable to scan row")
		}
		messages = append(messages, message)
		senders = append(senders, sender)
		macs = append(macs, mac)
	}

	return messages, senders, macs
}

// saveMessage will process the transaction to place a message
// into the database
func SaveMessage(message, sender string, recipient string, mac string) {
	database := Connect().Db

	database.Exec(`
		INSERT INTO Messages (sender, recipient, data, checksum)
		VALUES (
			(SELECT id FROM Users WHERE user = ?),
			(SELECT id FROM Users WHERE user = ?),
			?,
			?
		);
	`, sender, recipient, message, mac)
}
