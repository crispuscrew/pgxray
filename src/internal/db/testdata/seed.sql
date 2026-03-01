CREATE TABLE users (
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    email      TEXT UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE orders (
    id      SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    status  TEXT CHECK (status IN ('pending', 'done'))
);

INSERT INTO users (name, email) VALUES
    ('Alice', 'alice@example.com'),
    ('Bob',   NULL);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_orders_pending ON orders(status) WHERE status = 'pending';