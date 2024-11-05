package main

import (
	"log"

	"github.com/Isacco-B/go-safe-note-api/database"
	"github.com/Isacco-B/go-safe-note-api/repositories"
	"github.com/Isacco-B/go-safe-note-api/routes"
	"github.com/Isacco-B/go-safe-note-api/jobs"
)

func main() {
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	repositories.InitNoteCollection()

	r := routes.SetupRouter()
	cronjobs.StartCronJobs()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
