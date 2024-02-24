-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE PaymentMethods (
    PaymentMethodID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    CardNumber BYTEA NOT NULL,
    ExpirationDate DATE NOT NULL,
    CVV BYTEA NOT NULL,
    CardholderName VARCHAR(255) NOT NULL,
    CustomerID UUID,
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
