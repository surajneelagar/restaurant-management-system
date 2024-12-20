package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`                                                // MongoDB's Object ID
	Number_of_guests *int               `json:"number_of_guests" bson:"number_of_guests,omitempty"` // Optional field
	Table_number     *int               `json:"table_number" bson:"table_number,omitempty"`         // Optional field
	Created_at       time.Time          `json:"created_at" bson:"created_at"`                       // Timestamps
	Updated_at       time.Time          `json:"updated_at" bson:"updated_at"`
	Table_id         string             `json:"table_id" bson:"table_id"` // Unique identifier for the table
}
