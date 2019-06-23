package main

import (
	"flag"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nad2000/uoa-orcidhub/gen/restapi"
	"github.com/nad2000/uoa-orcidhub/gen/restapi/operations"

	log "github.com/sirupsen/logrus"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func handleUpdate(w http.ResponseWriter, r *http.Request) {
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.WithError(err).Fatalf("Failed to open the spec: %q", restapi.SwaggerJSON)
	}
	api := operations.NewOrcidhubIntegrationAPI(swaggerSpec)
	server := restapi.NewServer(api)
	flag.Parse()
	server.Port = *portFlag
	defer server.Shutdown()
	// TODO: Set Handle
	api.GetPingHandler = operations.GetPingHandlerFunc(
		func(params operations.GetPingParams) middleware.Responder {
			return operations.NewGetPingOK().WithPayload("hello!")
		})

	if err := server.Serve(); err != nil {
		log.WithError(err).Fatal("Failed to start the servers.")
	}
}
