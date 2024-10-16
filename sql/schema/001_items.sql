-- +goose Up
CREATE TABLE Items (
    classid BIGINT UNIQUE PRIMARY KEY,
    ItemName TEXT UNIQUE NOT NULL,
    DayChange DECIMAL(10, 2) NOT NULL,
    WeekChange DECIMAL(10, 2) NOT NULL,
    ImageURL TEXT NOT NULL
);
-- +goose Down
DROP TABLE Items;