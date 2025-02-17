CREATE TABLE report (
	customer_id VARCHAR(26) REFERENCES customer(id) ON DELETE CASCADE,
	id VARCHAR(26) PRIMARY KEY NOT NULL,
	type VARCHAR(20) NULL,
	output_file_path VARCHAR,
	download_url VARCHAR,
	download_url_expires_at TIMESTAMPTZ,
	error_message VARCHAR,
	created_at TIMESTAMPTZ NOT NULL,
	started_at TIMESTAMPTZ,
	failed_at TIMESTAMPTZ,
	completed_at TIMESTAMPTZ
)