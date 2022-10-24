package preference

type Repository interface {
	CreatePreference(profileId uint, lang string) (uint, error)
	GetLanguagePreference(profileId uint) (string, error)
}
