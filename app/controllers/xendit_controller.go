package controllers

import (
	"context"
	"net/http"
	"xendit-exercises/app/models"
	"xendit-exercises/app/usecases"
	"xendit-exercises/app/utils"
)

type XenditControllerInterface interface {
	CreateInvoice(w http.ResponseWriter, r *http.Request)
	GetInvoices(w http.ResponseWriter, r *http.Request)
}

type XenditController struct {
	ctx             context.Context
	XenditUseCase   usecases.XenditUseCaseInterface
	XenditValidator usecases.XenditValidatorInterface
}

func InitXenditControllerInterface(
	ctx context.Context,
	XenditUseCase usecases.XenditUseCaseInterface,
	XenditValidator usecases.XenditValidatorInterface,
) XenditControllerInterface {
	return &XenditController{
		ctx:             ctx,
		XenditUseCase:   XenditUseCase,
		XenditValidator: XenditValidator,
	}
}

func (c *XenditController) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	request, errorLog := c.XenditValidator.CreateInvoiceValidator(r)
	if errorLog.Error != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(errorLog))
		return
	}

	invoice, errorLog := c.XenditUseCase.CreateInvoice(request)
	if errorLog.Error != nil {
		response.Error = errorLog
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(response))
		return
	}

	response.Data = invoice
	response.StatusCode = http.StatusOK

	w.WriteHeader(http.StatusOK)
	w.Write(utils.WriteResponseBody(response))
}

func (c *XenditController) GetInvoices(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	request, errorLog := c.XenditValidator.GetInvoicesValidator(r)
	if errorLog.Error != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(errorLog))
		return
	}

	invoices, errorLog := c.XenditUseCase.GetInvoices(request)
	if errorLog.Error != nil {
		response.Error = errorLog
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(response))
		return
	}

	response.Data = invoices
	response.StatusCode = http.StatusOK

	w.WriteHeader(http.StatusOK)
	w.Write(utils.WriteResponseBody(response))
}
