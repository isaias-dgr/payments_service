-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE PlatformMerchants (
    PlatformID INT,
    MerchantID INT,
    FOREIGN KEY (PlatformID) REFERENCES OnlinePaymentPlatforms(id),
    FOREIGN KEY (MerchantID) REFERENCES Merchants(id),
    PRIMARY KEY (PlatformID, MerchantID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

