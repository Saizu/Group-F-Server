CREATE TABLE announces (
  id    INTEGER NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  title TEXT NOT NULL,
  body  TEXT NOT NULL,
  time  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
