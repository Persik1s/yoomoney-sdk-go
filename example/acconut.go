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
		"URI",
	)

	acc := account.NewYooAccount(client)
	data, err := acc.Info()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

}
