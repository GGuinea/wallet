-- +goose Up
CREATE TABLE IF NOT EXISTS entry (
	id text PRIMARY KEY,
	wallet_id text NOT NULL,
	op_type text NOT NULL,
	amount text NOT NULL,
	balance_after text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE entry;
