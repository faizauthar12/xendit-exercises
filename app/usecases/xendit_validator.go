package usecases

import (
	"encoding/json"
	"net/http"
	"xendit-exercises/app/models"
	"xendit-exercises/app/requests"
	"xendit-exercises/app/utils"
)

type XenditValidatorInterface interface {
	CreateInvoiceValidator(r *http.Request) (*requests.XenditCreateInvoiceRequest, *models.ErrorLog)
	GetInvoicesValidator(r *http.Request) (*requests.XenditGetInvoiceRequest, *models.ErrorLog)
}

type XenditValidator struct{}

func InitXenditValidatorInterface() XenditValidatorInterface {
	return &XenditValidator{}
}

func (v *XenditValidator) CreateInvoiceValidator(r *http.Request) (*requests.XenditCreateInvoiceRequest, *models.ErrorLog) {
	var (
		request        *requests.XenditCreateInvoiceRequest
		decodedRequest *requests.XenditCreateInvoiceRequest
	)

	err := json.NewDecoder(r.Body).Decode(&decodedRequest)
	if err != nil {
		errorLog := utils.WriteError(err, http.StatusBadRequest, "")
		return request, errorLog
	}

	request = decodedRequest
	return request, &models.ErrorLog{}
}

func (v *XenditValidator) GetInvoicesValidator(r *http.Request) (*requests.XenditGetInvoiceRequest, *models.ErrorLog) {
	var (
		request        *requests.XenditGetInvoiceRequest
		decodedRequest *requests.XenditGetInvoiceRequest
	)

	err := json.NewDecoder(r.Body).Decode(&decodedRequest)
	if err != nil {
		errorLog := utils.WriteError(err, http.StatusBadRequest, "")
		return request, errorLog
	}

	request = decodedRequest
	return request, &models.ErrorLog{}
}
