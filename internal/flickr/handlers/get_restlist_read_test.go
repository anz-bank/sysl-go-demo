package handlers

import (
	"testing"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/stretchr/testify/require"
)

func TestReadGetPetsList(t *testing.T) {
	req := flickr.GetRestListRequest{}
	client := flickr.GetRestListClient{}
	req.Method = util.NewString("flickr.photos.search")
	req.Tags = util.NewString("dog")

	// get the pet photo list
	response, err := GetRestListRead(util.NewRequestContext(), &req, client)

	// err check
	require.Nil(t, err)
	require.True(t, len(response.Photos) > 0)

	// verify one of the pet photos
	require.Equal(t, *response.Photos[1].ID, "I am JJ's dog Bingo")
	url := response.Photos[1].Urls.URL[0].Type
	require.Equal(t, *url, "http://www.example.com/dog/bingo")
}
