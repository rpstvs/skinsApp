-- +goose Up
CREATE TABLE Prices (
    PriceDate DATE NOT NULL,
    item_id TEXT NOT NULL REFERENCES Items(ClassId),
    Price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY(item_id, PriceDate),
    UNIQUE(item_id, PriceDate)
);
-- +goose Down
DROP TABLE Prices;