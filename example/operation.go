package main

import (
	"fmt"
	"github.com/Persik1s/yoomoney-sdk-go/yoomoney"
	"github.com/Persik1s/yoomoney-sdk-go/yoomoney/account"
)

func main() {
	client := yoomoney.NewClient(
		"CLIENT_ID",
		"ACCESS_TOKEN",
		"REDIRECT_URI",
	)

	acc := account.NewYooAccount(client)

	operations, err := acc.GetOperations(account.OperationRequest{
		Label:   "USER",
		Records: 30,
		Type:    account.OperationTypeDeposition,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(operations)
}
