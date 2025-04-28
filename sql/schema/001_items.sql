-- +goose Up
CREATE TABLE Items (
    ClassId TEXT UNIQUE NOT NULL,
    ItemName TEXT UNIQUE NOT NULL,
    DayChange DECIMAL(10, 2) NOT NULL,
    WeekChange DECIMAL(10, 2) NOT NULL,
    ImageURL TEXT NOT NULL
);
-- +goose Down
DROP TABLE Items;