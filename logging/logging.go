package logging

import (
	"os"
  "log"
  "time"
)

func LogMessage(username string, message string) {
  f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
		log.Fatal(err)
	}
  _, err = f.WriteString(time.Now().String() + "\t" + username + "\t" + message + "\n")
  if err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
