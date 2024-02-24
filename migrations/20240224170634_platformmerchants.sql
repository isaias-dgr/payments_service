-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE PlatformMerchants (
    PlatformID UUID,
    MerchantID UUID,
    FOREIGN KEY (PlatformID) REFERENCES OnlinePaymentPlatforms(PlatformID),
    FOREIGN KEY (MerchantID) REFERENCES Merchants(MerchantID),
    PRIMARY KEY (PlatformID, MerchantID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
