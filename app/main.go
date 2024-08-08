package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"xendit-exercises/app/routes"
	"xendit-exercises/app/utils"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/v6"
)

type Configurations struct {
	XenditAPIKey string
}

var cfg Configurations

func init() {
	err := godotenv.Load()
	if err != nil {
		errorLog := utils.WriteError(err, http.StatusInternalServerError, "")
		log.Fatalf("Error loading .env file")
		log.Fatalf("Error: %v", errorLog)
	}

	cfg.XenditAPIKey = os.Getenv("XENDIT_API_KEY")
}

func main() {
	ctx := context.TODO()

	xenditClient := xendit.NewClient(cfg.XenditAPIKey)

	mux := routes.Routes(ctx, xenditClient)

	fmt.Println("Starting server on :8000")
	http.ListenAndServe("localhost:8000", mux)
}
