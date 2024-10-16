-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING *;