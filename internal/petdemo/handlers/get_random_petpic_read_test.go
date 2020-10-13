package handlers

import (
	"context"
	"testing"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
	"github.com/stretchr/testify/require"
)

func TestGetRandomPetPicListRead(t *testing.T) {
	req := petdemo.GetRandomPetPicListRequest{}
	client := petdemo.GetRandomPetPicListClient{
		GetPetList: func(ctx context.Context, req *petstore.GetPetListRequest) (*petstore.Pet, error) {
			util.SetJSONResponseContentType(ctx)
			dummyPet := petstore.Pet{
				ID:   1,
				Name: "dummy pet dog",
				Tag:  util.NewString("dog"),
			}
			return &dummyPet, nil
		},
		GetRestList: func(ctx context.Context, req *flickr.GetRestListRequest) (*flickr.PhotoResource, error) {
			return &flickr.PhotoResource{
				Page:    util.NewFloat64(1),
				Pages:   util.NewFloat64(1),
				Perpage: util.NewFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: util.NewString("I am peter rabbit"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("rabbit"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/rabbit/pr"),
								},
							},
						},
					},
				},
				Total: util.NewFloat64(2),
			}, nil
		},
	}

	// get the random pet
	response, err := GetRandomPetPicListRead(util.NewRequestContext(), &req, client)

	// check err
	require.Nil(t, err)

	// verify pet
	require.Contains(t, response.Name, "rabbit")
}
