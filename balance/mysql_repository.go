package balance

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLBalanceAccountRepository struct {
	DB *gorm.DB
}

func NewMySQLBalanceAccountRepository(dsn string) (BalanceAccountRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&BalanceAccount{})
	return &MySQLBalanceAccountRepository{
		DB: db,
	}, nil
}

func (r *MySQLBalanceAccountRepository) CreateBalanceAccount(owner uint) (*BalanceAccount, error) {
	acc := BalanceAccount{
		Owner:   owner,
		Balance: uint(0),
	}

	result := r.DB.Create(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}

func (r *MySQLBalanceAccountRepository) GetByID(id uint) (*BalanceAccount, error) {
	var acc BalanceAccount
	result := r.DB.First(&acc, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}

func (r *MySQLBalanceAccountRepository) GetByOwner(uid uint) (*BalanceAccount, error) {
	var acc BalanceAccount
	result := r.DB.Where("owner = ?", strconv.FormatUint(uint64(uid), 10)).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}

func (r *MySQLBalanceAccountRepository) Debit(id uint, amt uint) (*BalanceAccount, error) {
	var acc BalanceAccount
	result := r.DB.First(&acc, id)
	if result.Error != nil {
		return nil, result.Error
	}

	acc.Balance += amt

	result = r.DB.Save(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}
