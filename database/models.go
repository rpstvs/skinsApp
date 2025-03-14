// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID         uuid.UUID
	Classid    string
	Itemname   string
	Daychange  float64
	Weekchange float64
	Imageurl   string
}

type Price struct {
	Pricedate time.Time
	ItemID    uuid.UUID
	Price     float64
}
