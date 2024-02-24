-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Merchants (
    MerchantID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(255) NOT NULL,
    BusinessID VARCHAR(255) UNIQUE NOT NULL,
    AccountInfoID INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
