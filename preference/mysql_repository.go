package preference

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	DB *gorm.DB
}

type Preference struct {
	ID       uint `gorm:"primaryKey"`
	Profile  uint
	Language string
}

func NewMySQLRepository(dsn string) (Repository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Preference{})

	return &mysqlRepository{
		DB: db,
	}, nil
}

func (r *mysqlRepository) CreatePreference(profileId uint, lang string) (uint, error) {
	p := Preference{
		Profile:  profileId,
		Language: lang,
	}

	result := r.DB.Create(&p)
	if result.Error != nil {
		return 0, result.Error
	}

	return p.ID, nil
}

func (r *mysqlRepository) GetLanguagePreference(profileId uint) (string, error) {
	var p Preference
	result := r.DB.Where("profile = ?", strconv.FormatUint(uint64(profileId), 10)).First(&p)
	if result.Error != nil {
		return "", result.Error
	}

	return p.Language, nil
}
