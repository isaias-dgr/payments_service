-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE BankAccounts (
    AccountInfoID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    AccountNumber BYTEA NOT NULL,
    BankName VARCHAR(255) NOT NULL,
    BranchCode VARCHAR(255) NOT NULL,
    MerchantID UUID,
    FOREIGN KEY (MerchantID) REFERENCES Merchants(MerchantID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
