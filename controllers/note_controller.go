package controllers

import (
	"net/http"
	"time"

	"github.com/Isacco-B/go-safe-note-api/models"
	"github.com/Isacco-B/go-safe-note-api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateNote(c *gin.Context) {
    var note models.Note
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	note.Link = uuid.NewString()
	note.ExpiredAt = primitive.NewDateTimeFromTime(time.Now().Add(24 * time.Hour))

    _, err := repositories.CreateUser(note)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"link": note.Link})
}

func GetNote(c *gin.Context) {
    link := c.Param("link")
    note, err := repositories.GetNoteBy(link)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

	if note.ExpiredAt < primitive.NewDateTimeFromTime(time.Now()) {
		err = repositories.DeleteNoteById(note.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Note expired"})
	}

	err = repositories.DeleteNoteById(note.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, note)
}
