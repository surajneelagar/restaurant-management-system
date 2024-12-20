package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`           // Unique ID as a string
	First_name    *string            `json:"first_name"`    // Optional field
	Last_name     *string            `json:"last_name"`     // Optional field
	Password      *string            `json:"password"`      // Optional field
	Email         *string            `json:"email"`         // Optional email
	Avatar        *string            `json:"avatar"`        // Optional profile picture URL
	Phone         *string            `json:"phone"`         // Optional phone number
	Token         *string            `json:"token"`         // Authentication token
	Refresh_Token *string            `json:"refresh_token"` // Refresh token for sessions
	Created_at    time.Time          `json:"created_at"`    // Creation timestamp
	Updated_at    time.Time          `json:"updated_at"`    // Last updated timestamp
	User_id       string             `json:"user_id"`       // Unique User ID
}
