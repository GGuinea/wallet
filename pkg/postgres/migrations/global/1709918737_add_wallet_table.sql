-- +goose Up
CREATE TABLE IF NOT EXISTS wallet (
	id text PRIMARY KEY,
	balance text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE wallet;
