package accounts

type Account struct {
	owner   string
	balance int
}

// Newaccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
