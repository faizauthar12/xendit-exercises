package responses

type XenditCreateInvoiceResponse struct {
	ExternalID string `json:"external_id"`
	ID         string `json:"id"`
	InvoiceURL string `json:"invoice_url"`
}
