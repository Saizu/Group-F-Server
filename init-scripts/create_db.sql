CREATE DATABASE db;

\c db;

CREATE TABLE announces (
  id    INTEGER NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  title TEXT NOT NULL,
  body  TEXT NOT NULL,
  time  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
  id     INTEGER NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  name   TEXT NOT NULL,
  banned BOOLEAN NOT NULL DEFAULT FALSE,
  last_login TIMESTAMP
);

CREATE TABLE inquiries (
  id    INTEGER NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  usrid INTEGER NOT NULL,
  title TEXT NOT NULL,
  body  TEXT NOT NULL,
  time  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  reply TEXT
);

CREATE TABLE items (
  id   INTEGER NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  name TEXT NOT NULL
);

CREATE TABLE users_items (
  usrid INTEGER  NOT NULL REFERENCES users ( id ),
  itmid INTEGER  NOT NULL REFERENCES items ( id ),
  amount INTEGER NOT NULL,
  PRIMARY KEY ( usrid, itmid )
);
