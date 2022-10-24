package preference

import "errors"

type Service interface {
	CreatePreference(p *PreferenceDTO) (*PreferenceDTO, error)
	GetLanguagePreference(profileId uint) (*PreferenceDTO, error)
}

var (
	ErrPreferenceNotFound = errors.New("Unable to find preference for given profile")
)

type service struct {
	R Repository
}

func NewPreferenceService(repo Repository) Service {
	return &service{
		R: repo,
	}
}

func (s *service) CreatePreference(p *PreferenceDTO) (*PreferenceDTO, error) {
	_, err := s.R.CreatePreference(p.ProfileID, p.Language)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *service) GetLanguagePreference(profileId uint) (*PreferenceDTO, error) {
	lang, err := s.R.GetLanguagePreference(profileId)
	if err != nil {
		return nil, err
	}

	return &PreferenceDTO{
		ProfileID: profileId,
		Language:  lang,
	}, nil
}
