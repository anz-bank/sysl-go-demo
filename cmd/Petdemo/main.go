package main

import (
    "context"
    "log"

    "github.com/anz-bank/sysl-go/core"

    "github.com/anz-bank/sysl-go-demo/gen/pkg/servers/Petdemo"
)

func main() {
    type AppConfig struct {
        // Define app-level config fields here.
    }
    log.Fatal(petdemo.Serve(context.Background(),
        func(ctx context.Context, config AppConfig) (*petdemo.ServiceInterface, *core.Hooks, error) {
            // Perform one-time setup based on config here.
            return &petdemo.ServiceInterface{
                // Add handlers here.
            }, nil, nil
        },
    ))
}
