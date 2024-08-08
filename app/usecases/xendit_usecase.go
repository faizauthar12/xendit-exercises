package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"xendit-exercises/app/models"
	"xendit-exercises/app/requests"
	"xendit-exercises/app/responses"
	"xendit-exercises/app/utils"

	"github.com/google/uuid"
	"github.com/xendit/xendit-go/v6"
	"github.com/xendit/xendit-go/v6/invoice"
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

func (u *XenditUseCase) CreateInvoice(request *requests.XenditCreateInvoiceRequest) (*responses.XenditCreateInvoiceResponse, *models.ErrorLog) {
	var (
		response *responses.XenditCreateInvoiceResponse
	)
	// externalId will be used as invoice ID
	externalId := uuid.New().String()
	fmt.Println("External ID: ", externalId)

	// create Customer Address Object
	customerAddress := *invoice.NewAddressObject()
	customerAddress.SetStreetLine1(request.CustomerAddress)

	// create Customer Object
	customer := *invoice.NewCustomerObject()
	customer.SetId(request.CustomerUUID)
	customer.SetGivenNames(request.CustomerName)
	customer.SetEmail(request.CustomerEmail)
	customer.SetAddresses([]invoice.AddressObject{customerAddress})
	customer.SetPhoneNumber(request.CustomerPhoneNumber)

	// create invoice items
	invoiceItems := []invoice.InvoiceItem{}
	if len(request.InvoiceItems) > 0 {
		for _, item := range request.InvoiceItems {
			invoiceItem := invoice.InvoiceItem{
				Name:     item.Name,
				Price:    item.Price,
				Quantity: float32(item.Quantity),
				Url:      &item.Url,
			}

			invoiceItems = append(invoiceItems, invoiceItem)
		}
	}

	createInvoiceRequest := invoice.CreateInvoiceRequest{
		ExternalId:         externalId,
		Amount:             request.Amount,
		Description:        &request.Description,
		Customer:           &customer,
		SuccessRedirectUrl: u.SuccessRedirectURL,
		FailureRedirectUrl: u.FailureRedirectURL,
		Items:              invoiceItems,
	}

	resp, r, errorXenditSdk := u.XenditClient.InvoiceApi.CreateInvoice(u.ctx).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()

	if r.StatusCode != http.StatusOK {
		if errorXenditSdk != nil {
			errorLog := utils.WriteError(errorXenditSdk, http.StatusInternalServerError, "")
			return response, errorLog
		}

		errorLog := utils.WriteError(errors.New(r.Status), r.StatusCode, "")
		return response, errorLog
	}

	respJson, _ := json.Marshal(resp)

	err := json.Unmarshal(respJson, response)
	if err != nil {
		errorLog := utils.WriteError(err, http.StatusInternalServerError, "")
		return response, errorLog
	}

	return response, &models.ErrorLog{}
}
