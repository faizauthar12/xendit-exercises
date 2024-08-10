package routes

import (
	"context"
	"net/http"
	"xendit-exercises/app/controllers"
	"xendit-exercises/app/models"
	"xendit-exercises/app/utils"

	"github.com/xendit/xendit-go/v6"
)

type ctrlGroup struct {
	xenditController controllers.XenditControllerInterface
}

func Routes(
	ctx context.Context,
	xenditClient *xendit.APIClient,
) http.Handler {

	ctrl := ctrlGroup{
		xenditController: controllers.InitHTTPXenditController(ctx, xenditClient),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		response := &models.Response{
			Data:       "pong",
			StatusCode: http.StatusOK,
		}

		w.WriteHeader(http.StatusOK)
		w.Write(utils.WriteResponseBody(response))
	})

	mux.HandleFunc("POST /invoices", ctrl.xenditController.CreateInvoice)
	mux.HandleFunc("GET /invoices", ctrl.xenditController.GetInvoices)

	return mux
}
