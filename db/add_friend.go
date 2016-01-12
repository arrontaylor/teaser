package db

import (
	"time"
)

var AddFriendToList = func(username1 string, username2 string) {
	Connection.Exec("INSERT INTO friends (username1, username2, time) VALUES (?, ?, ?)", username1, username2, time.Now().Unix())
}
