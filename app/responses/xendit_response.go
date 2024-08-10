package responses

type XenditCreateInvoiceResponse struct {
	ExternalID string `json:"external_id"`
	ID         string `json:"id"`
	InvoiceURL string `json:"invoice_url"`
}

type XenditGetInvoiceResponse struct {
	ID                        string `json:"id"`
	UserID                    string `json:"user_id"`
	ExternalID                string `json:"external_id"`
	Status                    string `json:"status"`
	MerchantName              string `json:"merchant_name"`
	MerchantProfilePictureURL string `json:"merchant_profile_picture_url"`
	Amount                    int64  `json:"amount"`
	PayerEmail                string `json:"payer_email"`
	Description               string `json:"description"`
	InvoiceURL                string `json:"invoice_url"`
}
