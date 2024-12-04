CREATE DATABASE db;

\c db;

CREATE TABLE players (
  id   text PRIMARY KEY,
  name text NOT NULL
);
