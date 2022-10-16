package profile

type ProfileDTO struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Service interface {
	CreateProfile(p *ProfileDTO) (*ProfileDTO, error)
	GetProfile(id uint) (*ProfileDTO, error)
}

type service struct {
	R ProfileRepository
}

func NewProfileService(repo ProfileRepository) Service {
	return &service{
		R: repo,
	}
}

func (s *service) CreateProfile(p *ProfileDTO) (*ProfileDTO, error) {
	entity, err := s.R.CreateProfile(p.FirstName, p.LastName, p.Email)
	if err != nil {
		return nil, err
	}

	dto := ProfileDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
	}

	return &dto, nil
}

func (s *service) GetProfile(id uint) (*ProfileDTO, error) {
	entity, err := s.R.GetProfile(id)
	if err != nil {
		return nil, err
	}

	dto := ProfileDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
	}
	return &dto, nil
}
