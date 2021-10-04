package main

import (
	"context"

	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo"
	"github.com/anz-bank/sysl-go-demo/internal/handlers"
	"github.com/anz-bank/sysl-go-demo/internal/service"
)

func main() {
	petdemo.Serve(context.Background(), service.CreateService(&petdemo.ServiceInterface{
		// Add handlers here.
		GetPetList: handlers.GetRandomPetPicListRead,
	}))
}
