-- +goose Up
CREATE TABLE IF NOT EXISTS product (
    id serial PRIMARY KEY,
    sku text,
    stock int,
    version bigint
);

CREATE TABLE IF NOT EXISTS orders (
    id serial PRIMARY KEY,
    product_id bigint,
    quantity int
);

INSERT INTO product ("sku", "stock", "version") VALUES ('x',5,0);

-- +goose Down
DROP TABLE product;
DROP TABLE orders;
