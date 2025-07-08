package satim

import "errors"

var (
	TestEndpoint string = "https://test.satim.dz/payment/rest"
	LiveEndpoint string = "https://cib.satim.dz/payment/rest"

	CurrencyDZD SatimCurrency = "012"

	LanguageAR SatimLanguage = "AR"
	LanguageFR SatimLanguage = "FR"
	LanguageEN SatimLanguage = "EN"
)

type SatimCurrency = string
type SatimLanguage = string
type SatimOrderStatus = string

type SatimRegisterOrderResponse struct {
	OrderID      *string `json:"orderId"`
	FormUrl      *string `json:"formUrl"`
	ErrorCode    *string `json:"errorCode"`
	ErrorMessage *string `json:"errorMessage"`
}

func (r *SatimRegisterOrderResponse) IsSuccessful() bool {
	return r.ErrorCode == nil || *r.ErrorCode == "0"
}

type SatimOrderConfirmResponse struct {
	ErrorCode             *string        `json:"ErrorCode"`
	ErrorMessage          *string        `json:"ErrorMessage"`
	DepositAmount         *int64         `json:"depositAmount"`
	ApprovalCode          *string        `json:"approvalCode"`
	Currency              *SatimCurrency `json:"currency"`
	Params                interface{}    `json:"params"`
	ActionCode            *int64         `json:"actionCode"`
	ActionCodeDescription *string        `json:"actionCodeDescription"`
	OrderStatus           *int64         `json:"OrderStatus"`
	OrderNumber           *string        `json:"OrderNumber"`
	Pan                   *string        `json:"Pan"`
	Amount                *int64         `json:"Amount"`
	Expiration            *string        `json:"expiration"`
	Ip                    *string        `json:"Ip"`
	SvfeResponse          *string        `json:"SvfeResponse"`
}

func (r *SatimOrderConfirmResponse) IsSuccessful() bool {
	return r.ErrorCode == nil || *r.ErrorCode == "0"
}

func (r *SatimOrderConfirmResponse) IsAlreadyConfirmed() bool {
	return r.ErrorCode != nil && *r.ErrorCode == "2"
}

type SatimOrderStatusResponse struct {
	ErrorCode     *string        `json:"ErrorCode"`
	ErrorMessage  *string        `json:"ErrorMessage"`
	OrderNumber   *string        `json:"OrderNumber"`
	OrderStatus   *int64         `json:"OrderStatus"`
	DepositAmount *int64         `json:"depositAmount"`
	Currency      *SatimCurrency `json:"currency"`
	Amount        *int64         `json:"Amount"`
	Params        interface{}    `json:"params"`
	ApprovalCode  *string        `json:"approvalCode"`
	Pan           *string        `json:"Pan"`
	Expiration    *string        `json:"expiration"`
	Ip            *string        `json:"Ip"`
	SvfeResponse  *string        `json:"SvfeResponse"`
}

func (r *SatimOrderStatusResponse) IsSuccessful() bool {
	return r.ErrorCode == nil && *r.ErrorCode == "0"
}

func (r *SatimOrderStatusResponse) GetStatus() (SatimOrderStatus, error) {
	if r.OrderStatus != nil {
		return OrderStatusFromCode(*r.OrderStatus), nil
	}
	return "", errors.New("failed to get order status")
}

type SatimOrderRefundResponse struct {
	ErrorCode    *string `json:"errorCode"`
	ErrorMessage *string `json:"errorMessage"`
}
