package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/arrontaylor/teaser/teaser"
)

func LoadFrequentsList(username string) teaser.FrequentsList {
	rows, err := Connection.Query("SELECT tousername, count(*) FROM teases WHERE fromusername=? GROUP BY tousername", username)

	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	usernames := make(map[string]int)

	for rows.Next() {
		var username string
		var count int
		rows.Scan(&username, &count)

		usernames[username] = count
	}

	usernames = pareDownList(usernames)

	return usernames
}

func pareDownList(usernames map[string]int) map[string]int {
	for len(usernames) > teaser.FREQUENTS_LIST_LENGTH {
		var smallestfriend = ""
		var smallestcount = -1

		for friend, count := range usernames {
			if smallestcount < 0 || smallestcount > count {
				smallestfriend = friend
				smallestcount = count
			}
		}

		delete(usernames, smallestfriend)
	}

	return usernames
}
