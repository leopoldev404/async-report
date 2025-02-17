CREATE TABLE refresh_token (
	customer_id VARCHAR(26) REFERENCES customer(id) ON DELETE CASCADE,
	token_hash VARCHAR(500) NOT NULL UNIQUE,
	created_at TIMESTAMPTZ NOT NULL,
	expires_at TIMESTAMPTZ NOT NULL,
	PRIMARY KEY (customer_id, token_hash)
)