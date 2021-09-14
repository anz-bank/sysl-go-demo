package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/anz-bank/sysl-go/common"

	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/petstore"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/pokeapi"
)

func TestPetDemo(t *testing.T) {
	t.Parallel()
	gatewayTester := petdemo.NewTestServer(t, context.Background(), createService, "")
	defer gatewayTester.Close()

	gatewayTester.Mocks.Petstore.GetPetList.
		MockResponse(http.StatusOK, map[string]string{"Content-Type": `application/json`}, petstore.Pet("gecko"))

	gatewayTester.Mocks.Pokeapi.GetPokemon.
		MockResponse(http.StatusOK, map[string]string{"Content-Type": `application/json`}, pokeapi.Pokemon{Name: common.NewString("charmeleon")})

	gatewayTester.GetPetList().
		ExpectResponseCode(http.StatusOK).
		ExpectResponseBody(petdemo.Pet{
			Breed:   "gecko",
			Pokemon: "charmeleon",
		}).
		Send()
}

func TestPetDemoFail(t *testing.T) {
	t.Parallel()
	gatewayTester := petdemo.NewTestServer(t, context.Background(), createService, "")
	defer gatewayTester.Close()

	gatewayTester.Mocks.Petstore.GetPetList.
		Timeout()

	gatewayTester.GetPetList().
		ExpectResponseCode(http.StatusServiceUnavailable).
		ExpectResponseBody(`{"status":{"code":"1013","description":"Downstream system is unavailable"}}`).
		Send()
}
