CREATE TABLE accounts (
  accountid INTEGER PRIMARY KEY NOT NULL,
  username TEXT,
  password TEXT,
  firstname TEXT,
  lastname TEXT,
  gender TEXT
);

CREATE TABLE teases (
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
