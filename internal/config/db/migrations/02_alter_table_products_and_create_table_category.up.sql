CREATE TABLE categories (
                        id SERIAL PRIMARY KEY,
                        name varchar(50) NOT NULL
);

ALTER TABLE products
    ADD COLUMN compare_at_price DOUBLE PRECISION NULL,
    ADD COLUMN category_id INTEGER REFERENCES categories(id);

