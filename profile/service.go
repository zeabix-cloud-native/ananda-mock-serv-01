package profile

import (
	"errors"

	"github.com/zeabix-cloud-native/ananda-mock-serv-01/clients"
)

type ProfileDTO struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Balance   uint   `json:"balance"`
}

type Service interface {
	CreateProfile(p *ProfileDTO) (*ProfileDTO, error)
	GetProfile(id uint) (*ProfileDTO, error)
	SetupAccount(id uint) (*ProfileDTO, error)
}

type service struct {
	R   ProfileRepository
	Acc clients.AccountService
}

var (
	ErrProfileNotFound = errors.New("Profile not found")
)

func NewProfileService(repo ProfileRepository, accService clients.AccountService) Service {
	return &service{
		R:   repo,
		Acc: accService,
	}
}

func (s *service) CreateProfile(p *ProfileDTO) (*ProfileDTO, error) {
	entity, err := s.R.CreateProfile(p.FirstName, p.LastName, p.Email)
	if err != nil {
		return nil, err
	}

	return s.SetupAccount(entity.ID)
}

func (s *service) GetProfile(id uint) (*ProfileDTO, error) {
	entity, err := s.R.GetProfile(id)
	if err != nil {
		return nil, ErrProfileNotFound
	}

	dto := ProfileDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
	}

	if entity.AccountID != uint(0) {
		// Get Balance from another service
		response, err := s.Acc.GetBalance(entity.AccountID)
		if err != nil {
			return nil, err
		}

		dto.Balance = response.Balance
	}

	return &dto, nil
}

func (s *service) SetupAccount(id uint) (*ProfileDTO, error) {
	acc, err := s.Acc.CreateAccount(id)
	if err != nil {
		return nil, err
	}

	updated, err := s.R.UpdateAccount(id, acc.ID)
	if err != nil {
		return nil, err
	}

	response := entityToDTO(updated)
	response.Balance = acc.Balance
	return response, nil
}

func entityToDTO(entity *Profile) *ProfileDTO {
	return &ProfileDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
	}
}
