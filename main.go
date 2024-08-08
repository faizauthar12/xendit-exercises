package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/v6"
	"github.com/xendit/xendit-go/v6/invoice"
)

type Configurations struct {
	XenditAPIKey string
}

var cfg Configurations

func writeError(err error) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		file = file[strings.LastIndex(file, "/")+1:]
		funcName := runtime.FuncForPC(pc).Name()

		log.Printf("Line: %d", line)
		log.Printf("File: %s", file)
		log.Printf("Function: %s", funcName)
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		writeError(err)
	}

	cfg.XenditAPIKey = os.Getenv("XENDIT_API_KEY")
}

func main() {
	ctx := context.TODO()

	xenditClient := xendit.NewClient(cfg.XenditAPIKey)

	// externalId will be used as invoice ID
	externalId := uuid.New().String()
	fmt.Println("External ID: ", externalId)

	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(externalId, float64(123))

	resp, r, err := xenditClient.InvoiceApi.CreateInvoice(ctx).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)

		writeError(err)
	}

	// response from `CreateInvoice`: Invoice
	respJson, _ := json.Marshal(resp)
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoice`: %v\n", string(respJson))
}
