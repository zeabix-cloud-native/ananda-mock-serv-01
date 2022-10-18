package balance

type BalanceAccountDTO struct {
	ID      uint `json:"id,omitempty"`
	Owner   uint `json:"owner"`
	Balance uint `json:"balance"`
}

type Service interface {
	CreateBalanceAccount(p *BalanceAccountDTO) (*BalanceAccountDTO, error)
	Debit(id uint, amt uint) (*BalanceAccountDTO, error)
	//Credit(id uint, amount uint) (*BalanceAccountDTO, error)
	Get(id uint) (*BalanceAccountDTO, error)
}

type service struct {
	R BalanceAccountRepository
}

func NewBalanceAccountService(repo BalanceAccountRepository) Service {
	return &service{
		R: repo,
	}
}

func (s *service) CreateBalanceAccount(p *BalanceAccountDTO) (*BalanceAccountDTO, error) {
	entity, err := s.R.CreateBalanceAccount(p.Owner)
	if err != nil {
		return nil, err
	}

	return entityToDTO(entity), nil
}

func (s *service) Get(id uint) (*BalanceAccountDTO, error) {
	entity, err := s.R.GetByID(id)
	if err != nil {
		return nil, err
	}

	return entityToDTO(entity), nil
}

func (s *service) Debit(id uint, amt uint) (*BalanceAccountDTO, error) {
	entity, err := s.R.Debit(id, amt)
	if err != nil {
		return nil, err
	}

	return entityToDTO(entity), nil
}

func entityToDTO(entity *BalanceAccount) *BalanceAccountDTO {
	return &BalanceAccountDTO{
		ID:      entity.ID,
		Owner:   entity.Owner,
		Balance: entity.Balance,
	}
}
