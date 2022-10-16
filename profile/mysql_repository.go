package profile

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLProfileRepository struct {
	DB *gorm.DB
}

func NewMySQLProfileRepository(dsn string) (ProfileRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Profile{})
	return &MySQLProfileRepository{
		DB: db,
	}, nil
}

func (r *MySQLProfileRepository) CreateProfile(firstname string, lastname string, email string) (*Profile, error) {

	p := Profile{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
	}

	result := r.DB.Create(&p)
	if result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

func (r *MySQLProfileRepository) GetProfile(ID uint) (*Profile, error) {
	var p Profile
	result := r.DB.First(&p, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}
