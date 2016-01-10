package db

import "github.com/arrontaylor/teaser/teaser"

var LoadFriendsList = func(username string) teaser.FriendsList {
	friendslist := make([]string, 0)

	rows, err := Connection.Query("SELECT username2 FROM friends WHERE username1=?", username)

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
