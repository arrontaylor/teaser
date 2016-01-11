package db

import "time"
import "github.com/arrontaylor/teaser/teaser"

var LoadTeasesList = func(username string) teaser.TeasesList {
	teaseslist := make([]teaser.Tease, 0)

	rows, err := Connection.Query("SELECT teaseid, fromusername, time, read FROM teases WHERE tousername=? AND read=0", username)

	if err != nil {
		return nil
	}

	for rows.Next() {
		tease := teaser.Tease{ToUsername: username}
		var t int64

		rows.Scan(&tease.TeaseId, &tease.FromUsername, &t, &tease.Read)
		tease.Time = time.Unix(t, 0)

		teaseslist = append(teaseslist, tease)
	}

	return teaseslist
}
