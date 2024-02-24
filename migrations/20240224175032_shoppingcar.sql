-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE ShoppingCarts (
    CartID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    OwnerID VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ShoppingCartItems (
    ItemID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    CartID UUID NOT NULL,
    ProductID UUID NOT NULL,
    Quantity INT NOT NULL CHECK (Quantity > 0),
    FOREIGN KEY (CartID) REFERENCES ShoppingCarts(CartID) ON DELETE CASCADE,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
