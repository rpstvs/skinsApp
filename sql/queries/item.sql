-- name: CreateItem :exec
INSERT INTO Items (
        id,
        Classid,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetItemIDbyName :one
SELECT id
FROM Items
WHERE ItemName = $1;
-- name: UpdatePriceChange :exec
UPDATE Items
SET DayChange = $1,
    WeekChange = $2
WHERE Id = $3;