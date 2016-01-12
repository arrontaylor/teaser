CREATE TABLE accounts (
  username TEXT,
  password TEXT,
  firstname TEXT,
  lastname TEXT,
  gender TEXT
);

CREATE TABLE teases (
  teaseid INTEGER PRIMARY KEY NOT NULL,
  fromusername TEXT,
  tousername TEXT,
  time INTEGER,
  read INTEGER
);

CREATE TABLE friends (
  username1 TEXT,
  username2 TEXT,
  time INTEGER
);

CREATE TABLE ranks (
  count INTEGER,
  name TEXT
);