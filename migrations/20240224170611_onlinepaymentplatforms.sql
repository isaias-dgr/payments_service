-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE OnlinePaymentPlatforms (
    PlatformID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(255) NOT NULL,
    SupportedCards TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
