package account

import "github.com/Persik1s/yoomoney-sdk-go/yoomoney"

type YooAccount struct {
	Client *yoomoney.Client
}

func NewYooAccount(client *yoomoney.Client) *YooAccount {
	return &YooAccount{
		Client: client,
	}
}
