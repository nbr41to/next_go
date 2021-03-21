CREATE TABLE IF NOT EXISTS todos(
    id  SERIAL PRIMARY KEY,
    text TEXT,
    created TIMESTAMP
);
