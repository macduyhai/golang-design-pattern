package builder

// Co thể tạo ra nhiều biến thể khác nhau cho cùng một đối tượng

type BankAccount interface {
	WithDraw(atm uint64)
	Deposit(atm uint64)
	GetBalance() int64
}
type BankAccountBuilder interface {
	WithOwnerName(name string) BankAccountBuilder
	WithOwnerIdentity(identificationNumber uint64) BankAccountBuilder
	AtBranch(branch string) BankAccountBuilder
	OpeningBalance(balance int64) BankAccountBuilder
	Build() BankAccount
}
type bankAccount struct {
	ownerName            string
	identificationNumber uint64
	branch               string
	balance              int64
}

func (acc *bankAccount) WithDraw(amt uint64) {
	acc.balance -= int64(amt)
}

func (acc *bankAccount) Deposit(amt uint64) {
	acc.balance += int64(amt)

}

func (acc *bankAccount) GetBalance() int64 {
	return acc.balance
}

func (acc *bankAccount) WithOwnerName(name string) BankAccountBuilder {
	acc.ownerName = name
	return acc
}

func (acc *bankAccount) WithOwnerIdentity(identificationNumber uint64) BankAccountBuilder {
	acc.identificationNumber = identificationNumber
	return acc
}

func (acc *bankAccount) AtBranch(branch string) BankAccountBuilder {
	acc.branch = branch
	return acc
}

func (acc *bankAccount) OpeningBalance(balance int64) BankAccountBuilder {
	acc.balance = int64(balance)
	return acc
}

func (acc *bankAccount) Build() BankAccount {
	return acc
}

func NewBankAccountBuilder() BankAccountBuilder {
	return &bankAccount{}
}
