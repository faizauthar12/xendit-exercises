package models

import "time"

type XenditAvailableEWalletItem struct {
	EWalletType string `json:"ewallet_type"`
}

type XenditAvailablePaylaterItem struct {
	PaylaterType string `json:"paylater_type"`
}

type XenditAvailableQRCodeItem struct {
	QRCodeType string `json:"qr_code_type"`
}

type XenditAvailabileRetailOutletItem struct {
	RetailOutletName string `json:"retail_outlet_name"`
}

type XenditInvoice struct {
	AvailableEWallets       []XenditAvailableEWalletItem       `json:"available_ewallets"`
	AvailablePaylaters      []XenditAvailablePaylaterItem      `json:"available_paylaters"`
	AvailableQRCodes        []XenditAvailableQRCodeItem        `json:"available_qr_codes"`
	AvailableRetailOutlets  []XenditAvailabileRetailOutletItem `json:"available_retail_outlets"`
	Created                 time.Time                          `json:"created"`
	Currency                string                             `json:"currency"`
	ExpiryDate              time.Time                          `json:"expiry_date"`
	ExternalID              string                             `json:"external_id"`
	ID                      string                             `json:"id"`
	InvoiceURL              string                             `json:"invoice_url"`
	MerchantName            string                             `json:"merchant_name"`
	MerchantProfilePicture  string                             `json:"merchant_profile_picture"`
	ShouldExcludeCreditCard bool                               `json:"should_exclude_credit_card"`
	ShouldSendEmail         bool                               `json:"should_send_email"`
	Status                  string                             `json:"status"`
	Updated                 time.Time                          `json:"updated"`
	UserID                  string                             `json:"user_id"`
}

type XenditCustomer struct {
	GivenName    string `json:"given_name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Addresses    string `json:"addresses"`
}

type XenditInvoiceWebhook struct {
	ID                     string    `json:"id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 int64     `json:"amount"`
	PaidAmount             int64     `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 time.Time `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	AdjustedReceivedAmount int64     `json:"adjusted_received_amount"`
	FeesPaidAmount         int64     `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
}
