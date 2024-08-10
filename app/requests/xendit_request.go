package requests

import "time"

type XenditInvoiceItem struct {
	Name     string  `json:"name" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Quantity int64   `json:"quantity" validate:"required"`
	Url      string  `json:"url" validate:"required"`
	Category string  `json:"category"`
}

type XenditCreateInvoiceRequest struct {
	CustomerPhoneNumber string              `json:"customer_phone_number" validate:"required"`
	CustomerName        string              `json:"customer_name" validate:"required"`
	CustomerEmail       string              `json:"customer_email" validate:"required"`
	CustomerUUID        string              `json:"customer_uuid" validate:"required"`
	CustomerAddress     string              `json:"customer_address" validate:"required"`
	CustomerCountry     string              `json:"customer_country" validate:"required"`
	Description         string              `json:"description" validate:"required"`
	InvoiceItems        []XenditInvoiceItem `json:"invoice_items" validate:"required"`
}

type XenditGetInvoiceRequest struct {
	InvoiceID     string     `json:"invoice_id"`
	ExternalID    string     `json:"external_id"`
	Limit         int64      `json:"limit"`
	CreatedAfter  *time.Time `json:"created_after"`
	CreatedBefore *time.Time `json:"created_before"`
	PaidAfter     *time.Time `json:"paid_after"`
	PaidBefore    *time.Time `json:"paid_before"`
}
