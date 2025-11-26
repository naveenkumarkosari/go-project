-- +goose up

CREATE TABLE feed_follows(
  id uuid PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL,
  user_id uuid not null REFERENCES users(id) on delete cascade,
  feed_id uuid not null REFERENCES feeds(id) on delete cascade,
  UNIQUE(user_id,feed_id)
 );

-- +goose Down
DROP TABLE feed_follows;
