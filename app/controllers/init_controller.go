package controllers

import (
	"context"
	"xendit-exercises/app/usecases"

	"github.com/xendit/xendit-go/v6"
)

func InitHTTPXenditController(ctx context.Context, xenditClient *xendit.APIClient) XenditControllerInterface {
	xenditUseCase := usecases.InitXenditUseCaseInterface(ctx, xenditClient)
	xenditValidator := usecases.InitXenditValidatorInterface()

	handler := InitXenditControllerInterface(ctx, xenditUseCase, xenditValidator)
	return handler
}
