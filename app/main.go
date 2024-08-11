package main

import (
	"context"
	"flag"
	"fmt"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
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

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "local")
	} else if args[0] != "local" && args[0] != "ngrok" {
		args[0] = "local"
	}

	mux := routes.Routes(ctx, xenditClient)

	switch args[0] {
	case "local":
		startLocal(mux)
	case "ngrok":
		startNgrok(ctx, mux)
	default:
		startLocal(mux)
	}
}

func startLocal(
	mux http.Handler,
) {

	fmt.Println("Starting server on :8000")
	http.ListenAndServe(":8000", mux)
}

func startNgrok(
	ctx context.Context,
	mux http.Handler,
) {

	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Starting server on ", listener.URL())
	http.Serve(listener, mux)
}
