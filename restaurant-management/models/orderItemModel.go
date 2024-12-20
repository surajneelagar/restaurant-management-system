package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `json:"id"`            // Unique identifier as a string
	Quantity      *string            `json:"quantity"`      // Optional field for quantity
	Unit_price    *float64           `json:"unit_price"`    // Optional field for unit price
	Created_at    time.Time          `json:"created_at"`    // Timestamp for creation
	Updated_at    time.Time          `json:"updated_at"`    // Timestamp for last update
	Food_id       *string            `json:"food_id"`       // Optional field for associated food item
	Order_item_id string             `json:"order_item_id"` // Unique identifier for the order item
	Order_id      string             `json:"order_id"`      // Identifier for the associated order
}
