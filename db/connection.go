package db

import "database/sql"
import _ "github.com/mattn/go-sqlite3"

var Connection, _ = sql.Open("sqlite3", "./db.sqlite")
