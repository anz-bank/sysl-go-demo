package util

import (
	"context"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
)

//  NewString creates a string pointer.
func NewString(s string) *string {
	return &s
}

//  NewFloat64 creates a float pointer.
func NewFloat64(f float64) *float64 {
	return &f
}

// SetJSONResponseContentType sets the content type of the response to the appropriate application/json value.
func SetJSONResponseContentType(ctx context.Context) {
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
}

// NewRequestContext creates a new request context.
func NewRequestContext() context.Context {
	headers := http.Header{}
	ctx := context.Background()
	ctx = common.RequestHeaderToContext(ctx, headers)
	return ctx
}
