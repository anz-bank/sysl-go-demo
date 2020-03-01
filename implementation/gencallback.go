// implementation/gencallback.go

// gencallback.go contains all the manual config code that is used to implement the generated sysl
package implementation

import (
	"context"
	"github.com/go-chi/chi"
	"github.service.anz/sysl/sysltemplate/gen/mydependency"
	"github.service.anz/sysl/sysltemplate/gen/simple"
	"github.service.anz/sysl/sysltemplate/myconfig"
	"log"
	"net/http"
)
var serverAddress = ":8080"
func LoadServices(ctx context.Context) error {
	router := chi.NewRouter()

	// simpleServiceInterface is the struct which is composed of our functions we wrote in `methods.go`

	// Struct embedding is used for the Service interface (yes, not interfaces)
	simpleServiceInterface := simple.ServiceInterface{
		GetFoobarList: GetFoobarList,
	}

	// Default callback behaviour
	genCallbacks := myconfig.DefaultCallback()

	// Service Handler
	// use the Example service that implements behaviour of the downstream service
	// Here we specify that the base path (serviceURL) of the url we're hitting; https://jsonplaceholder.typicode.com
	// Sometimes this will have endpoints attached https://example.com/v2 but should never have a trailing slash
	// Note that this ServiceURL is http, as our service is served over http
	serviceHandler := simple.NewServiceHandler(genCallbacks,
												&simpleServiceInterface,
												mydependency.NewClient(http.DefaultClient, "http://jsonplaceholder.typicode.com"))

	// Service Router
	serviceRouter := simple.NewServiceRouter(genCallbacks, serviceHandler)
	serviceRouter.WireRoutes(ctx, router)

	log.Println("Starting Server on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
	return nil
}
