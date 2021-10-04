package handlers

import (
	"context"
	"flag"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/core"
	"github.com/stretchr/testify/require"

	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/petstore"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/pokeapi"
	"github.com/anz-bank/sysl-go-demo/internal/service"
)

func petEnabledCreateService() func(_ context.Context, _ service.AppConfig) (*petdemo.ServiceInterface, *core.Hooks, error) {
	return service.CreateService(&petdemo.ServiceInterface{
		GetPetList: GetRandomPetPicListRead,
	})
}

func TestPetDemo(t *testing.T) {
	t.Parallel()
	gatewayTester := petdemo.NewTestServer(t, context.Background(), petEnabledCreateService(), "")
	defer gatewayTester.Close()

	gatewayTester.Mocks.Petstore.GetPetList.
		MockResponse(http.StatusOK, map[string]string{"Content-Type": `application/json`}, petstore.Pet("gecko"))

	gatewayTester.Mocks.Pokeapi.GetPokemon.
		ExpectURLParamID(5).
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
	gatewayTester := petdemo.NewTestServer(t, context.Background(), petEnabledCreateService(), "")
	defer gatewayTester.Close()

	gatewayTester.Mocks.Petstore.GetPetList.
		Timeout()

	gatewayTester.GetPetList().
		ExpectResponseCode(http.StatusServiceUnavailable).
		ExpectResponseBody(`{"status":{"code":"1013","description":"Downstream system is unavailable"}}`).
		Send()
}

func OnlyRunIfRequested(t *testing.T) {
	if m := flag.Lookup("test.run").Value.String(); m == "" || !regexp.MustCompile(m).MatchString(t.Name()) {
		t.Skip("skipping as execution was not requested explicitly using go test -run")
	}
}

func TestPetDemo_Integration(t *testing.T) {
	OnlyRunIfRequested(t)

	cfgBytes, err := os.ReadFile("../../config/config.yaml")
	require.NoError(t, err)

	t.Parallel()
	gatewayTester := petdemo.NewIntegrationTestServer(t, context.Background(), petEnabledCreateService(), cfgBytes)
	defer gatewayTester.Close()

	gatewayTester.GetPetList().
		ExpectResponseCode(http.StatusOK).
		TestResponseBody(func(t *testing.T, actual []byte) {
			// Test the result here (or use ExpectResponseBody)
			// The results for this test are random so we can't actually test them
			require.NotNil(t, actual)
		}).
		Send()
}
