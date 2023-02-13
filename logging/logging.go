package logging

import (
	"os"
  "fmt"
  "log"
)

func logMessage() {
  fmt.Println(" Woohoo I got here")
  f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
		log.Fatal(err)
	}
  _, err = f.WriteString("testing"as)
  if err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
