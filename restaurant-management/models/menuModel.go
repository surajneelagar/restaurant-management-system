package models

import (
	"time"
)

type Menu struct {
	ID         string     `bson:"_id"`
	Name       string     `json:"name"`
	Category   string     `json:"category"`
	Start_Date *time.Time `json:"start_date,omitempty"`
	End_Date   *time.Time `json:"end_date,omitempty"`
	Created_at *time.Time `json:"created_at,omitempty"`
	Updated_at *time.Time `json:"updated_at,omitempty"`
	Manu_id    string     `json:"manu_id"`
}
