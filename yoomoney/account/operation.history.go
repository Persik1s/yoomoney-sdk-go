package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	_OperationHistoryUrl = "https://yoomoney.ru/api/operation-history"
)

const (
	StatusSuccess  = "success"     // success — платеж завершен успешно;
	StatusRefused  = "refused"     // refused — платеж отвергнут получателем или отменен отправителем;
	StatusProgress = "in_progress" // in_progress — платеж не завершен или перевод не принят получателем.
)

// Направление движения средств. Может принимать значения:
const (
	DirectionIn  = "in"  // in (приход);
	DirectionOut = "out" // out (расход);
)

const (
	TypePayment    = "payment-shop"      // payment-shop — исходящий платеж в магазин;
	TypeTransfer   = "outgoing-transfer" // outgoing-transfer — исходящий P2P-перевод любого типа;
	TypeDeposition = "deposition"        // deposition — зачисление;
	TypeIncoming   = "incoming-transfer" // incoming-transfer — входящий перевод.
)

type Operation struct {
	OperationId string    `json:"operation_id"`
	Status      string    `json:"status"`
	DateTime    time.Time `json:"datetime"`

	Title     string  `json:"title"`
	Direction string  `json:"direction"`
	Amount    float64 `json:"amount"`
	Label     string  `json:"label"`
	Type      string  `json:"type"`
}

// --------------------------------------------------------------------- //

// Типы операции получение или отправка
const (
	OperationTypeDeposition = "deposition"
	OperationTypePayment    = "payment"
)

type OperationRequest struct {
	Type        string        `url:"type"`                   // Тип операции
	Label       string        `url:"label"`                  // Поиск по label
	From        time.Duration `url:"from,omitempty"`         // Вывести все операции  начиная с N времяни
	Till        time.Duration `url:"till,omitempty"`         // Вывести все операции  до N времяни
	StartRecord string        `url:"start_record,omitempty"` // Eсли параметр присутствует, то будут отображены операции, начиная с номера start_record. Операции нумеруются с 0. Подробнее про постраничный вывод списка
	Records     int           `url:"records,omitempty"`      // Количество запрашиваемых записей истории операций. Допустимые значения: от 1 до 100, по умолчанию — 30.
}

type OperationRespone struct {
	Operations []Operation `json:"operations"`
	Error      string      `json:"error"`
	NextRecord string      `json:"next_record"`
}

func (y *YooAccount) GetOperations(data OperationRequest) (OperationRespone, error) {
	value, err := query.Values(data)
	if err != nil {
		return OperationRespone{}, err
	}
	valueEncode := value.Encode()

	if data.Records == 0 {
		data.Records = 30
	}
	fmt.Println(valueEncode)

	request, err := http.NewRequest(http.MethodPost, _OperationHistoryUrl, bytes.NewReader([]byte(valueEncode)))

	request.Header.Set("Content-Length", strconv.Itoa(len(valueEncode)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer ", y.Client.AccessToken))

	if err != nil {
		return OperationRespone{}, err
	}

	client := &http.Client{}
	respone, err := client.Do(request)
	if err != nil {
		return OperationRespone{}, err
	}

	operationsRespone := OperationRespone{
		Operations: make([]Operation, data.Records),
	}

	body, err := io.ReadAll(respone.Body)

	if err != nil {
		return OperationRespone{}, err
	}

	if err := json.Unmarshal(body, &operationsRespone); err != nil {
		return OperationRespone{}, err
	}

	return operationsRespone, nil
}
