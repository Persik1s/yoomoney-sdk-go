package account

import "sdk/yoomoney"

type YooAccount struct {
	Client *yoomoney.Client
}

func NewYooAccount(client *yoomoney.Client) *YooAccount {
	return &YooAccount{
		Client: client,
	}
}
