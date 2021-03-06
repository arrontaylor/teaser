package db

import "fmt"
import "github.com/arrontaylor/teaser/teaser"

var CreateAccount = func(username string, password string, firstname string, lastname string, gender string) *teaser.Account {
	account := teaser.GetAccount(username)

	if account != nil {
		return account
	} else {
		account = &teaser.Account{
			Username:  username,
			Password:  password,
			FirstName: firstname,
			LastName:  lastname,
			Gender:    gender,
		}
	}

	res, err := Connection.Exec("INSERT INTO accounts (username, password, firstname, lastname, gender) VALUES (?, ?, ?, ?, ?)", username, password, firstname, lastname, gender)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	accountId, err := res.LastInsertId()
	account.AccountId = int(accountId)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return account
}
