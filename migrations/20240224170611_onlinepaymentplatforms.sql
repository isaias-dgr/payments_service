-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE OnlinePaymentPlatforms (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    PlatformID UUID UNIQUE,
    Name VARCHAR(255) NOT NULL,
    SupportedCards TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
