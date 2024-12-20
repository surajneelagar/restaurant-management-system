package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name"`
	Price      *float64           `json:"price"`
	Food_image *string            `json:"food_image"`
	Created_at time.Time          `json:"create_at"`
	Updated_at time.Time          `json:"updated_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"manu_id"`
}
