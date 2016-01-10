package db

var AddFriendToList = func(username1 string, username2 string) {
	Connection.Exec("INSERT INTO friends (username1, username2) VALUES (?, ?)", username1, username2)
}
