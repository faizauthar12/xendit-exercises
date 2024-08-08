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
	var (
		response *models.Response
	)

	w.Header().Set("Content-Type", "application/json")

	request, errorLog := c.XenditValidator.CreateInvoiceValidator(r)
	if errorLog != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(errorLog))
		return
	}

	invoice, errorLog := c.XenditUseCase.CreateInvoice(request)
	if errorLog != nil {
		response.StatusCode = errorLog.StatusCode
		response.Error = errorLog
		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(errorLog))
		return
	}

	response.Data = invoice

	w.WriteHeader(http.StatusOK)
	w.Write(utils.WriteResponseBody(response))
}
