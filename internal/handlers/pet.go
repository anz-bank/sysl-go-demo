package handlers

import (
	"context"

	petdemo "github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/petstore"
	"github.com/anz-bank/sysl-go-demo/internal/gen/pkg/servers/Petdemo/pokeapi"
	"github.com/anz-bank/sysl-go/common"
)

// GetRandomPetPicListRead reads random pic from downstream
func GetRandomPetPicListRead(ctx context.Context,
	getRandomPetPicListRequest *petdemo.GetPetListRequest,
	client petdemo.GetPetListClient) (*petdemo.Pet, error) {

	// Set response content type to JSON
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}

	reqPetstore := petstore.GetPetListRequest{}
	pet, err := client.PetstoreGetPetList(ctx, &reqPetstore)
	if err != nil {
		return nil, err
	}

	// Set response encoding type to gzip
	// This is required as only gzip encoding is currently supported by sysl-codegen
	headers = common.RequestHeaderFromContext(ctx)
	headers["Accept-Encoding"] = []string{"gzip"}

	reqPokemon := pokeapi.GetPokemonRequest{ID: int64(len(*pet))}
	pokemon, err := client.PokeapiGetPokemon(ctx, &reqPokemon)
	if err != nil {
		return nil, err
	}

	// return the result
	return &petdemo.Pet{
		Breed:   string(*pet),
		Pokemon: *pokemon.Name,
	}, nil
}
