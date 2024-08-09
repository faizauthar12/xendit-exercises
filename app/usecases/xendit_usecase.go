package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xendit/xendit-go/v6/invoice"
	"net/http"
	"os"
	"xendit-exercises/app/models"
	"xendit-exercises/app/requests"
	"xendit-exercises/app/responses"
	"xendit-exercises/app/utils"

	"github.com/google/uuid"
	"github.com/xendit/xendit-go/v6"
)

type XenditUseCaseInterface interface {
	CreateInvoice(request *requests.XenditCreateInvoiceRequest) (*responses.XenditCreateInvoiceResponse, *models.ErrorLog)
}

type XenditUseCase struct {
	ctx                context.Context
	XenditClient       *xendit.APIClient
	SuccessRedirectURL *string
	FailureRedirectURL *string
}

func InitXenditUseCaseInterface(
	ctx context.Context,
	XenditClient *xendit.APIClient,
) XenditUseCaseInterface {
	successRedirectURL := fmt.Sprintf("%s%s", "http://localhost:8080", "/success")
	failureRedirectURL := fmt.Sprintf("%s%s", "http://localhost:8080", "/failure")

	return &XenditUseCase{
		ctx:                ctx,
		XenditClient:       XenditClient,
		SuccessRedirectURL: &successRedirectURL,
		FailureRedirectURL: &failureRedirectURL,
	}
}

// full documentation of Xendit Create Invoice API can be found here: https://developers.xendit.co/api-reference/#create-invoice
func (u *XenditUseCase) CreateInvoice(request *requests.XenditCreateInvoiceRequest) (*responses.XenditCreateInvoiceResponse, *models.ErrorLog) {

	response := &responses.XenditCreateInvoiceResponse{}

	// externalId will be used as invoice ID
	externalId := uuid.New().String()
	fmt.Println("External ID: ", externalId)

	// create Customer Address Object
	customerAddress := *invoice.NewAddressObject()
	customerAddress.SetStreetLine1(request.CustomerAddress)
	customerAddress.SetCountry(request.CustomerCountry)

	// create Customer Object
	customer := *invoice.NewCustomerObject()
	customer.SetId(request.CustomerUUID)
	customer.SetGivenNames(request.CustomerName)
	customer.SetEmail(request.CustomerEmail)
	customer.SetAddresses([]invoice.AddressObject{customerAddress})
	customer.SetPhoneNumber(request.CustomerPhoneNumber)

	// create invoice items
	totalAmount := float64(0)
	invoiceItems := []invoice.InvoiceItem{}
	if len(request.InvoiceItems) > 0 {
		for _, item := range request.InvoiceItems {
			invoiceItem := invoice.InvoiceItem{
				Name:     item.Name,
				Price:    item.Price,
				Quantity: float32(item.Quantity),
				Url:      &item.Url,
			}

			totalAmount += float64(item.Price * float32(item.Quantity))
			invoiceItems = append(invoiceItems, invoiceItem)
		}
	}

	createInvoiceRequest := invoice.CreateInvoiceRequest{
		ExternalId:         externalId,
		Amount:             float64(totalAmount),
		Description:        &request.Description,
		Customer:           &customer,
		SuccessRedirectUrl: u.SuccessRedirectURL,
		FailureRedirectUrl: u.FailureRedirectURL,
		Items:              invoiceItems,
	}

	resp, r, errorXenditSdk := u.XenditClient.InvoiceApi.CreateInvoice(u.ctx).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()

	if errorXenditSdk != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", errorXenditSdk.Error())

		b, _ := json.Marshal(errorXenditSdk.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)

		errorLog := utils.WriteError(errorXenditSdk, http.StatusInternalServerError, "")
		return response, errorLog
	}

	response.ID = *resp.Id
	response.InvoiceURL = resp.InvoiceUrl
	response.ExternalID = resp.ExternalId

	return response, &models.ErrorLog{}
}
