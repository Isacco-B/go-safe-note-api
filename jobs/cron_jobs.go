package cronjobs

import (
	"log"
	"time"

	"github.com/Isacco-B/go-safe-note-api/models"
	"github.com/Isacco-B/go-safe-note-api/repositories"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartCronJobs() {
	c := cron.New()
	_, err := c.AddFunc("*/1 * * * *", func() {
		var notes []models.Note
		err := repositories.GetAllNotes(&notes)
		if err != nil {
			log.Fatal(err)
		}

		if len(notes) == 0 {
			log.Println("No notes found")
		}

		for _, note := range notes {
			if note.ExpiredAt < primitive.NewDateTimeFromTime(time.Now()) {
				err = repositories.DeleteNoteById(note.ID.Hex())
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		log.Println("Running job at: ", time.Now())
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
