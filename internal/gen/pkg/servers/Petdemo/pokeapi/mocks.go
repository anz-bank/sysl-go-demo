// Code generated by sysl DO NOT EDIT.
package pokeapi

import (
	"encoding/json"

	"github.com/anz-bank/sysl-go/testutil/e2e"
)

type DownstreamMocks struct {
	GetPokemon *GetPokemonMock
}

func NewDownstreamMocks(tester *e2e.Tester) *DownstreamMocks {
	return &DownstreamMocks{
		GetPokemon: NewGetPokemonMock(tester),
	}
}

type GetPokemonMock struct {
	e     e2e.Endpoint
	tests []e2e.Tests
}

func NewGetPokemonMock(tester *e2e.Tester) *GetPokemonMock {
	d := tester.NewDownstream("pokeapi", "GET", "/pokemon/{id}")

	return &GetPokemonMock{e: d}
}

func (d *GetPokemonMock) ExpectHeaders(headers map[string]string) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectHeaders(headers))

	return d
}

func (d *GetPokemonMock) ExpectHeadersExist(headers []string) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectHeadersExist(headers))

	return d
}

func (d *GetPokemonMock) ExpectHeadersDoNotExist(headers []string) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectHeadersDoNotExist(headers))

	return d
}

func (d *GetPokemonMock) ExpectHeadersExistExactly(headers []string) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectHeadersExistExactly(headers))

	return d
}

func (d *GetPokemonMock) ExpectBodyPlain(body []byte) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectBody(body))

	return d
}

func (d *GetPokemonMock) ExpectQueryParams(query map[string][]string) *GetPokemonMock {
	d.tests = append(d.tests, e2e.ExpectQueryParams(query))

	return d
}

func (d *GetPokemonMock) Expect(test e2e.Tests) *GetPokemonMock {
	d.tests = append(d.tests, test)

	return d
}

func (d *GetPokemonMock) MockResponse(returnCode int, returnHeaders map[string]string, returnBody interface{}) {
	var bodyBytes []byte
	switch returnBody := returnBody.(type) {
	case []byte:
		bodyBytes = returnBody
	case string:
		bodyBytes = ([]byte)(returnBody)
	default:
		bodyBytes, _ = json.Marshal(returnBody)
	}
	d.tests = append(d.tests, e2e.Response(returnCode, returnHeaders, bodyBytes))
	d.e.Expect(d.tests...)
	d.tests = nil
}

func (d *GetPokemonMock) MockResponsePlain(returnCode int, returnHeaders map[string]string, returnBody []byte) {
	d.tests = append(d.tests, e2e.Response(returnCode, returnHeaders, returnBody))
	d.e.Expect(d.tests...)
	d.tests = nil
}

func (d *GetPokemonMock) Timeout() {
	d.tests = append(d.tests, e2e.ForceDownstreamTimeout())
	d.e.Expect(d.tests...)
	d.tests = nil
}
