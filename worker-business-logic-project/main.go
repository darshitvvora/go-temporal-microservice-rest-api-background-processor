package main

import (
	"log"
)

func main() {
	if err := startWorker(); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
