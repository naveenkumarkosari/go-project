-- +goose up

CREATE TABLE users(
  id uuid PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
