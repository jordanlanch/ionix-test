-- +goose Up
-- +goose StatementBegin
CREATE TABLE vaccinations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  drug_id INTEGER REFERENCES drugs(id),
  dose INTEGER,
  date TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS vaccinations;
