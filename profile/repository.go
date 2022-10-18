package profile

type Profile struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	AccountID uint
}

type ProfileRepository interface {
	CreateProfile(firstname string, lastname string, email string) (*Profile, error)
	GetProfile(ID uint) (*Profile, error)
	UpdateAccount(id uint, accId uint) (*Profile, error)
}
