-- name: CreateItem :one
INSERT INTO Items (
        id,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetItemIDbyName :one
SELECT id
FROM Items
WHERE ItemName = $1;