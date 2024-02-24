-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE Payments (
    PaymentID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    PaymentMethodID UUID,
    CartID UUID,
    Total DECIMAL NOT NULL,
    Status VARCHAR(15) NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (CartID) REFERENCES ShoppingCarts(CartID),
    FOREIGN KEY (PaymentMethodID) REFERENCES PaymentMethods(PaymentMethodID)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
