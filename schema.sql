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