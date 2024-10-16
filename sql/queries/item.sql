-- name: CreateItem :one
INSERT INTO Items (
        classid,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;