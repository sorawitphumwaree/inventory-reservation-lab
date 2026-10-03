CREATE TABLE products (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(btrim(name)) > 0),
    initial_stock BIGINT NOT NULL CHECK (initial_stock >= 0),
    available_stock BIGINT NOT NULL,
    CONSTRAINT products_available_stock_check
        CHECK (available_stock BETWEEN 0 AND initial_stock)
);
