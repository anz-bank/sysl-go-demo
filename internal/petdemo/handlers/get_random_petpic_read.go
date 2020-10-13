package handlers

import (
	"context"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
)

// GetRandomPetPicListRead reads random pic from downstream
func GetRandomPetPicListRead(ctx context.Context,
	getRandomPetPicListRequest *petdemo.GetRandomPetPicListRequest,
	client petdemo.GetRandomPetPicListClient) (*petdemo.PetResponse, error) {
	util.SetJSONResponseContentType(ctx)
	reqPetstore := petstore.GetPetListRequest{}
	// get the pet list
	pet, err := client.GetPetList(ctx, &reqPetstore)
	if err != nil {
		return nil, err
	}
	reqFlickr := flickr.GetRestListRequest{
		Method: util.NewString("flickr.photos.search"),
		Tags:   util.NewString(*pet.Tag),
	}
	photo, err := client.GetRestList(ctx, &reqFlickr)
	if err != nil {
		return nil, err
	}

	// return the result
	return &petdemo.PetResponse{
		Name: *(photo.Photos[0].ID),
		URI:  *photo.Photos[0].Urls.URL[0].Type,
	}, nil
}
