package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `json:"id"`         // Unique identifier
	Text       string             `json:"text"`       // Note content
	Title      string             `json:"title"`      // Title of the note
	Created_at time.Time          `json:"created_at"` // Creation timestamp
	Updated_at time.Time          `json:"updated_at"` // Last updated timestamp
	Note_id    string             `json:"note_id"`    // Unique Note ID
}
