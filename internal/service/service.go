package service

import (
	"context"

	petdemo "github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo"
	"github.com/anz-bank/sysl-go/core"
)

type AppConfig struct {
	// Define app-level config fields here.
}

func CreateService(si *petdemo.ServiceInterface) func(_ context.Context, _ AppConfig) (*petdemo.ServiceInterface, *core.Hooks, error) {
	return func(_ context.Context, _ AppConfig) (*petdemo.ServiceInterface, *core.Hooks, error) {
		// Perform one-time setup based on config here.
		return si, nil, nil
	}
}
