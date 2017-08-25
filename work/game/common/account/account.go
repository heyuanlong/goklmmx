package account



type Account struct {
	AccountId int
}

func NewAccount() *Account  {
	a := new(Account)
	return a
}