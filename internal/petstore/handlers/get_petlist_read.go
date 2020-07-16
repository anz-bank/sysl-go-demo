package handlers

import (
	"context"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

// GetPetsListRead fetches a pet from the pet store.
func GetPetsListRead(ctx context.Context,
	getPetsListRequest *petstore.GetPetListRequest,
	client petstore.GetPetListClient) (*petstore.Pet, error) {
	util.SetJSONResponseContentType(ctx)

	// create the dummy pet
	dummyPet := petstore.Pet{
		ID:   1,
		Name: "dummy pet dog",
		Tag:  util.NewString("dog"),
	}

	// return the result
	return &dummyPet, nil
}
