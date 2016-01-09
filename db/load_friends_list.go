package db

import "github.com/arrontaylor/teaser/server"

var LoadFriendsList = func(username string) server.FriendsList {
	friendslist := make([]string, 0)

	rows, err := Connection.Query("SELECT username FROM accounts")

	if err != nil {
		return nil
	}

	for rows.Next() {
		var friend string
		rows.Scan(&friend)
		friendslist = append(friendslist, friend)
	}

	return friendslist
}
