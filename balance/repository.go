package balance

type BalanceAccount struct {
	ID      uint `gorm:"primaryKey"`
	Owner   uint
	Balance uint
}

type BalanceAccountRepository interface {
	CreateBalanceAccount(owner uint) (*BalanceAccount, error)
	GetByID(id uint) (*BalanceAccount, error)
	GetByOwner(uid uint) (*BalanceAccount, error)
	Debit(id uint, amt uint) (*BalanceAccount, error)
}
