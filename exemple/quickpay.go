package main

import (
	"fmt"
	"sdk/yoomoney/operation"
)

func main() {
	quickpay := operation.NewYooQuickPay()

	url, err := quickpay.GenerateQuickPay(operation.QuickPayForm{
		Label:       "USER",
		Sum:         10,
		AccountId:   "4100118697218483",
		QuickPay:    operation.QuickPayShop,
		PayMentType: operation.PayMentPC,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url)
}
