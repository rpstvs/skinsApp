-- name: CreateItem :exec
INSERT INTO Items (
        Classid,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdatePriceChange :exec
UPDATE Items
SET DayChange = $1,
    WeekChange = $2
WHERE Classid = $3;