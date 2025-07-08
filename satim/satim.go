package satim

import (
	"encoding/json"
	"errors"
	"sync"
)

type SatimClient struct {
	httpClient *SatimHttpClient
	username   string
	password   string
	terminalId string
}

func NewSatim(username, password, terminalId string, testMode bool) *SatimClient {
	endpoint := LiveEndpoint
	if testMode {
		endpoint = TestEndpoint
	}
	return &SatimClient{
		httpClient: newSatimHttpClient(endpoint),
		username:   username,
		password:   password,
		terminalId: terminalId,
	}
}

type SatimOrderBuilder struct {
	orderNumber       *string
	amount            *int64
	currency          SatimCurrency
	language          SatimLanguage
	returnUrl         *string
	failUrl           *string
	description       *string
	userDefinedFields map[string]string
	userDefinedMu     sync.RWMutex
}

func (s *SatimClient) NewOrder() *SatimOrderBuilder {
	return &SatimOrderBuilder{
		currency:          CurrencyDZD,
		language:          LanguageFR,
		userDefinedFields: make(map[string]string),
		userDefinedMu:     sync.RWMutex{},
	}
}

func (b *SatimOrderBuilder) WithOrderNumber(orderNumber string) *SatimOrderBuilder {
	b.orderNumber = &orderNumber
	return b
}

func (b *SatimOrderBuilder) WithAmount(amount int64) *SatimOrderBuilder {
	b.amount = &amount
	return b
}

func (b *SatimOrderBuilder) WithCurrency(currency SatimCurrency) *SatimOrderBuilder {
	b.currency = currency
	return b
}

func (b *SatimOrderBuilder) WithLanguage(language SatimLanguage) *SatimOrderBuilder {
	b.language = language
	return b
}

func (b *SatimOrderBuilder) WithReturnUrl(returnUrl string) *SatimOrderBuilder {
	b.returnUrl = &returnUrl
	return b
}

func (b *SatimOrderBuilder) WithFailUrl(failUrl string) *SatimOrderBuilder {
	b.failUrl = &failUrl
	return b
}

func (b *SatimOrderBuilder) WithDescription(description string) *SatimOrderBuilder {
	b.description = &description
	return b
}

func (b *SatimOrderBuilder) WithUserDefinedField(key, value string) *SatimOrderBuilder {
	b.userDefinedMu.Lock()
	defer b.userDefinedMu.Unlock()
	b.userDefinedFields[key] = value
	return b
}

func (b *SatimOrderBuilder) validate() error {
	if b.orderNumber == nil {
		return errors.New("orderNumber is required")
	}

	if b.amount == nil {
		return errors.New("amount is required")
	}

	// Currently with DZD Currency
	if *b.amount < 5000 {
		return errors.New("amount must be greater than 5000")
	}

	if b.returnUrl == nil {
		return errors.New("returnUrl is required")
	}

	b.userDefinedMu.RLock()
	defer b.userDefinedMu.RUnlock()
	// check if userDefinedFields values and keys are less than 20 chars
	for k, v := range b.userDefinedFields {
		if v == "" {
			return errors.New("userDefinedFields must have a value")
		}

		if len(k) > 20 {
			return errors.New("userDefinedFields must not have a length greater than 20 characters")
		}

		if len(v) > 20 {
			return errors.New("userDefinedFields values must not have a length more than 20 characters")
		}
	}

	return nil
}

func (b *SatimOrderBuilder) GenerateOrderDetails() (map[string]interface{}, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	finalData := map[string]interface{}{
		"amount":      *b.amount,
		"currency":    b.currency,
		"language":    b.language,
		"returnUrl":   *b.returnUrl,
		"orderNumber": *b.orderNumber,
	}

	if b.failUrl != nil {
		finalData["failUrl"] = *b.failUrl
	}

	if b.description != nil {
		finalData["description"] = *b.description
	}

	b.userDefinedMu.RLock()
	defer b.userDefinedMu.RUnlock()

	result := map[string]interface{}{
		"data":              finalData,
		"userDefinedFields": b.userDefinedFields,
	}

	return result, nil
}

func (s *SatimClient) RegisterOrder(orderData map[string]interface{}) (*SatimRegisterOrderResponse, error) {
	// Turn orderData into query params but first append orderData["userDefinedFields"]["force_terminal_id"] = terminalId
	finalData := orderData["data"].(map[string]interface{})
	finalData["userName"] = s.username
	finalData["password"] = s.password
	userDefinedFields := orderData["userDefinedFields"].(map[string]string)
	userDefinedFields["force_terminal_id"] = s.terminalId

	query := map[string]interface{}{}
	for key, value := range finalData {
		query[key] = value
	}

	//json enconde userDefinedFields into query["jsonParams"]
	jsonParams, err := json.Marshal(userDefinedFields)
	if err != nil {
		return nil, err
	}
	query["jsonParams"] = string(jsonParams)

	return s.httpClient.RegisterOrderQuery(query)
}

func (s *SatimClient) ConfirmOrder(orderId string) (*SatimOrderConfirmResponse, error) {
	finalData := make(map[string]interface{})
	finalData["userName"] = s.username
	finalData["password"] = s.password
	finalData["orderId"] = orderId

	query := map[string]interface{}{}
	for key, value := range finalData {
		query[key] = value
	}

	return s.httpClient.ConfirmOrderQuery(query)
}

func (s *SatimClient) GetOrderStatus(orderId string) (*SatimOrderStatusResponse, error) {
	finalData := make(map[string]interface{})
	finalData["userName"] = s.username
	finalData["password"] = s.password
	finalData["orderId"] = orderId

	query := map[string]interface{}{}
	for key, value := range finalData {
		query[key] = value
	}

	return s.httpClient.GetOrderStatusQuery(query)
}

func (s *SatimClient) RefundOrder(orderId string, amount int64) (*SatimOrderRefundResponse, error) {
	finalData := make(map[string]interface{})
	finalData["userName"] = s.username
	finalData["password"] = s.password
	finalData["orderId"] = orderId
	finalData["amount"] = amount

	query := map[string]interface{}{}
	for key, value := range finalData {
		query[key] = value
	}

	return s.httpClient.RefundOrderQuery(query)
}
