package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zeabix-cloud-native/ananda-mock-serv-01/balance"
)

type AccountService interface {
	CreateAccount(owner uint) (*balance.BalanceAccountDTO, error)
	GetBalance(acc_id uint) (*balance.BalanceAccountDTO, error)
}

type accountservice struct {
	Endpoint string
}

func NewAccountService(api string) AccountService {
	return &accountservice{
		Endpoint: api,
	}
}

func (s *accountservice) CreateAccount(owner uint) (*balance.BalanceAccountDTO, error) {
	url := fmt.Sprintf("%s/%s", s.Endpoint, "balance/accounts")

	reqJson, err := json.Marshal(balance.BalanceAccountDTO{
		Owner: owner,
	})
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var acc balance.BalanceAccountDTO
	err = json.Unmarshal(responseData, &acc)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (s *accountservice) GetBalance(acc_id uint) (*balance.BalanceAccountDTO, error) {
	url := fmt.Sprintf("%s/%s/%d", s.Endpoint, "balance/accounts", acc_id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var acc balance.BalanceAccountDTO
	err = json.Unmarshal(responseData, &acc)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}
