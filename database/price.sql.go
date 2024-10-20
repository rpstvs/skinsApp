// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: price.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addPrice = `-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING pricedate, item_id, price
`

type AddPriceParams struct {
	Pricedate time.Time
	ItemID    uuid.UUID
	Price     float64
}

func (q *Queries) AddPrice(ctx context.Context, arg AddPriceParams) ([]Price, error) {
	rows, err := q.db.QueryContext(ctx, addPrice, arg.Pricedate, arg.ItemID, arg.Price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Price
	for rows.Next() {
		var i Price
		if err := rows.Scan(&i.Pricedate, &i.ItemID, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPricebyId = `-- name: GetPricebyId :many
SELECT Price
FROM Prices
WHERE Item_id = $1
`

func (q *Queries) GetPricebyId(ctx context.Context, itemID uuid.UUID) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getPricebyId, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var price float64
		if err := rows.Scan(&price); err != nil {
			return nil, err
		}
		items = append(items, price)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
