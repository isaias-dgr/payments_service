-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE Customers DROP COLUMN PaymentMethodID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
