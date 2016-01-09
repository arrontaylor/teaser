package db

import "time"
import "github.com/arrontaylor/teaser/server"

var CreateTease = func(fromusername string, tousername string) *server.Tease {
	tease := &server.Tease{
		FromUsername: fromusername,
		ToUsername:   tousername,
		Time:         time.Now(),
		Read:         false,
	}

	res, err := Connection.Exec("INSERT INTO teases (fromusername, tousername, time, read) VALUES (?,?,?,?)", tease.FromUsername, tease.ToUsername, tease.Time.Unix(), tease.Read)

	if err != nil {
		return nil
	}

	rowsAffected, err := res.RowsAffected()

	if rowsAffected != 1 || err != nil {
		return nil
	}

	return tease
}
