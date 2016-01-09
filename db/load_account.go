package db

import "github.com/arrontaylor/teaser/server"

var LoadAccount = func(username string) *server.Account {
	account := &server.Account{}

	row := Connection.QueryRow("SELECT accountid, username, password, firstname, lastname, gender FROM accounts WHERE username=?", username)

	err := row.Scan(&account.AccountId, &account.Username, &account.Password, &account.FirstName, &account.LastName, &account.Gender)

	if err != nil {
		return nil
	}

	return account
}
