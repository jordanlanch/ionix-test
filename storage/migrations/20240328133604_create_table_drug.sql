-- +goose Up
-- +goose StatementBegin
CREATE TABLE drugs (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  approved BOOLEAN,
  min_dose INTEGER,
  max_dose INTEGER,
  available_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS drugs;
