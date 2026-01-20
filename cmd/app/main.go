package main

import (
	"log"

	"github.com/dmitrie43/LibrarySearcherUser/internal/app/services/rabbitmq"
)

func main() {
	rabbitmq.Listening()

	log.Print("Continue")
}
