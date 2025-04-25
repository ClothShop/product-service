CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price DOUBLE PRECISION NOT NULL,
                          created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE images (
                        id SERIAL PRIMARY KEY,
                        product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                        url TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT now()
);
