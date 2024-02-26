-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Products (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    ProductID UUID UNIQUE,
    Name VARCHAR(255) NOT NULL,
    Price DECIMAL NOT NULL,
    Description TEXT,
    StockQuantity INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
