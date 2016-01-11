package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/arrontaylor/teaser/teaser"
)

var LoadAccount = func(username string) *teaser.Account {
	account := &teaser.Account{}

	row := Connection.QueryRow("SELECT accountid, username, password, firstname, lastname, gender FROM accounts WHERE username=?", username)

	err := row.Scan(&account.AccountId, &account.Username, &account.Password, &account.FirstName, &account.LastName, &account.Gender)

	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	row = Connection.QueryRow("SELECT count(*) FROM teases WHERE fromusername=?", username)

	err = row.Scan(&account.SentCount)

	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	return account
}
