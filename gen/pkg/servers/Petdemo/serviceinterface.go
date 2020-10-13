// Code generated by sysl DO NOT EDIT.
package petdemo

import (
	"context"
	"time"

	"github.com/anz-bank/sysl-go/config"

	"github.com/anz-bank/sysl-go-demo/gen/pkg/servers/Petdemo/petstore"
)

// DefaultPetdemoImpl ...
type DefaultPetdemoImpl struct {
}

// NewDefaultPetdemoImpl for Petdemo
func NewDefaultPetdemoImpl() *DefaultPetdemoImpl {
	return &DefaultPetdemoImpl{}
}

// GetPetListClient provides access to all
// the clients used by the GetPetList method.
type GetPetListClient struct {
	PetstoreGetPetList func(
		ctx context.Context,
		req *petstore.GetPetListRequest,
	) (*petstore.Pet, error)
}

// ServiceInterface for Petdemo
type ServiceInterface struct {
	GetPetList func(ctx context.Context, req *GetPetListRequest, client GetPetListClient) (*Pet, error)
}

// DownstreamConfig for Petdemo
type DownstreamConfig struct {
	ContextTimeout time.Duration               `yaml:"contextTimeout"`
	Petstore       config.CommonDownstreamData `yaml:"petstore"`
}