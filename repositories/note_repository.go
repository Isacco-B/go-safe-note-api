package repositories

import (
	"context"
	"log"
	"time"

	"github.com/Isacco-B/go-safe-note-api/database"
	"github.com/Isacco-B/go-safe-note-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection

func InitNoteCollection() {
	var err error
	noteCollection, err = database.GetCollection("safe_note", "notes")
	if err != nil {
		log.Fatalf("Failed to get notes collection: %v", err)
	}
}

func CreateUser(note models.Note) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := noteCollection.InsertOne(ctx, note)
	return result, err
}

func GetNoteBy(link string) (models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var note models.Note
	err := noteCollection.FindOne(ctx, bson.M{"link": link}).Decode(&note)
	return note, err
}

func GetAllNotes(notes *[]models.Note) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := noteCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	return cursor.All(ctx, notes)
}

func DeleteNoteById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
	_, err = noteCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
