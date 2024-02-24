-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE AcquiringBanks (
    BankID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(255) NOT NULL,
    BankCode VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
