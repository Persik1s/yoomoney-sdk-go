package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	_accountInfoUrl = "https://yoomoney.ru/api/account-info"
)

type Account struct {
	AccountId   string  `json:"account"`      // account id
	AccountType string  `json:"account_type"` // тип аккаунта
	Balance     float64 `json:"balance"`      // баланс
	Currency    string  `json:"currency"`
}

func (a *YooAccount) Info() (Account, error) {
	request, err := http.NewRequest(http.MethodPost, _accountInfoUrl, nil)
	if err != nil {
		return Account{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", a.Client.AccessToken))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", "0")

	client := &http.Client{}
	respone, err := client.Do(request)
	if err != nil {
		return Account{}, err
	}

	body, err := io.ReadAll(respone.Body)
	if err != nil {
		return Account{}, err
	}

	accountInfo := Account{}
	if err := json.Unmarshal(body, &accountInfo); err != nil {
		return Account{}, err
	}
	return accountInfo, nil
}
