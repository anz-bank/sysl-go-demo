package myconfig

import (
	"context"
	"github.com/go-chi/chi"
	"github.service.anz/sysl/server-lib/common"
	"github.service.anz/sysl/server-lib/validator"
	"net/http"
	"time"
)

func DefaultCallback()Callback{
	return Callback{
		UpstreamTimeout:   120 * time.Second,
		DownstreamTimeout: 120 * time.Second,
		RouterBasePath:    "/",
		UpstreamConfig:    nil,
	}
}

type Callback struct {
	UpstreamTimeout   time.Duration
	DownstreamTimeout time.Duration
	RouterBasePath    string
	UpstreamConfig    validator.Validator
}

type Config struct{}

func (c Config) Validate() error {
	return nil
}

func (g Callback) AddMiddleware(ctx context.Context, r chi.Router) {
}

func (g Callback) BasePath() string {
	return "/"
}

func (g Callback) Config() validator.Validator {
	return Config{}
}

func (g Callback) HandleError(ctx context.Context, w http.ResponseWriter, kind common.Kind, message string, cause error) {
	se := common.CreateError(ctx, kind, message, cause)

	httpError := common.HandleError(ctx, se)

	httpError.WriteError(ctx, w)
}

func (g Callback) DownstreamTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 120*time.Second)
}

