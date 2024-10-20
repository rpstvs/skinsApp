-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPricebyId :many
SELECT Price
FROM Prices
WHERE Item_id = $1;