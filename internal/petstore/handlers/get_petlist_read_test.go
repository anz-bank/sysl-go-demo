package handlers

import (
	"testing"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/petstore"
	"github.com/stretchr/testify/require"
)

func TestReadGetPetsList(t *testing.T) {
	req := petstore.GetPetListRequest{}
	client := petstore.GetPetListClient{}

	// get the pet list
	response, err := GetPetsListRead(util.NewRequestContext(), &req, client)

	// verify the list count
	require.Nil(t, err)

	// verify pet
	require.Contains(t, response.Name, "dog")
	require.Contains(t, *(response.Tag), "dog")
}
