package operation

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

const (
	_quickPayUrl = "https://yoomoney.ru/quickpay/confirm.xml?"
)

const (
	QuickPayShop   = "shop"
	QuickPayDonate = "donate"
	QuickPaySmall  = "small"
)

const (
	PayMentPC = "pc" // оплата с кошелька
	PatMentAC = "ac" // оплата с банковской карты
)

type QuickPayForm struct {
	AccountId   string  `url:"receiver"`
	QuickPay    string  `url:"quickpay-form"`
	PayMentType string  `url:"paymentType"`
	Sum         float64 `url:"sum"`
	Label       string  `url:"label,omitempty"`
}

type YooQuickPay struct{}

func NewYooQuickPay() *YooQuickPay {
	return &YooQuickPay{}
}

func (y *YooQuickPay) GenerateQuickPay(form QuickPayForm) (string, error) {
	value, err := query.Values(form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", _quickPayUrl, value.Encode()), err
}
