-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Products (
    ProductID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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
