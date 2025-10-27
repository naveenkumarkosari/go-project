-- +goose up

CREATE TABLE feeds(
  id uuid PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL,
  content Text,
  createdBy uuid NOT NULL,
  CONSTRAINT  fk_userid 
  FOREIGN KEY (createdBy)
  REFERENCES users(id)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
