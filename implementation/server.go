// implementation/server.go

// server.go contains all the manual config code that is used to implement the generated sysl
package implementation

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.service.anz/sysl/sysltemplate/gen/jsonplaceholder"
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

	serviceHandler := simple.NewServiceHandler(genCallbacks,
												&simpleServiceInterface,
												jsonplaceholder.NewClient(http.DefaultClient, "http://jsonplaceholder.typicode.com"))

	// Service Router
	serviceRouter := simple.NewServiceRouter(genCallbacks, serviceHandler)
	serviceRouter.WireRoutes(ctx, router)

	log.Println("Starting Server on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
	return nil
}

// GetFoobarList refers to the endpoint in our sysl file
func GetFoobarList(ctx context.Context, req *simple.GetFoobarListRequest, client simple.GetFoobarListClient) (*jsonplaceholder.TodosResponse, error) {

	// Here we can make a request on the client object which was generated from the call to "myDownstream" in the sysl model
	// We will get the id equal to one, which was generated from out {id} from /todos/{id<:int}
	ans, err := client.GetTodos(ctx, &jsonplaceholder.GetTodosRequest{ID: 1})
	fmt.Println(ans)
	return ans, err
}
