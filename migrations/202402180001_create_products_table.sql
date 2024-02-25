-- +goose Up

-- +goose StatementBegin 
CREATE TYPE category AS ENUM ('VIDA', 'AUTO', 'VIAGEM', 'RESIDENCIAL', 'PATRIMONIAL');

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name varchar(255) UNIQUE,
    category category,
    base_price decimal,
    tariffed_price decimal,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 

DROP TABLE products;
Drop Type category;

-- +goose StatementEnd