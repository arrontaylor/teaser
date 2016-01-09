package teaser

type Account struct {
	AccountId int
	Username  string
	Password  string
	FirstName string
	LastName  string
	SessionId int
}

var LoadAccount func(username string) *Account
var CreateAccount func(username string, password string) *Account

var Accounts = make(map[string]*Account)

func GetAccount(username string) *Account {
	if Accounts[username] == nil {
		Accounts[username] = LoadAccount(username)
	}

	return Accounts[username]
}