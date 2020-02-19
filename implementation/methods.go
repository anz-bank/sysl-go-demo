// implementation/methods.go

// methods.go contains the actual HTTP functions that control the behaviour for the endpoints defined in the sysl file
package implementation

import (
	"context"

	// These are the imports for our generated code
	"github.service.anz/sysl/sysltemplate/gen/mydependency"
	"github.service.anz/sysl/sysltemplate/gen/simple"
)

// GetFoobarList refers to the endpoint in our sysl file
func GetFoobarList(ctx context.Context, req *simple.GetFoobarListRequest, client simple.GetFoobarListClient) (*simple.Str, error) {

	// Here we can make a request on the client object which was generated from the call to "myDownstream" in the sysl model
	// We will get the id equal to one, which was generated from out {id} from /todos/{id<:int}
	ans, err := client.GetTodos(ctx, &mydependency.GetTodosRequest{ID: 1})
	foo := simple.Str(ans.Title)

	return &foo, err
}
