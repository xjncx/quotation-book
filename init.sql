CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS quotes (
    id SERIAL PRIMARY KEY,
    quote_uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
    author VARCHAR(255),
    quote_text VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS quote_unique_idx
ON quotes (author, quote_text);

CREATE UNIQUE INDEX IF NOT EXISTS quotes_quote_uuid_idx
ON quotes (quote_uuid);
