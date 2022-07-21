package api

import (
	"context"
	"net/http"
	"time"

	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

var (
	stdSrv *http.Server
)

type Config struct {
	Endpoint string `default:":8080"`
}

// MustServe serves the API endpoints.
// Be minded this function will block until application exit.
func MustServe(svcFactory *service.Factory) {
	var config Config
	viperutil.MustUnmarshalKey("api", &config)

	// TODO: add JSON-RPC support

	stdSrv = &http.Server{
		Addr:        config.Endpoint,
		ReadTimeout: 1 * time.Minute,
		Handler:     handlers.RecoveryHandler()(newRouter(svcFactory)),
	}

	if err := stdSrv.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithField("endpoint", config.Endpoint).
			WithError(err).
			Fatal("Failed to serve api endpoint")
	}
}

// Shutdown shutdowns the API endpoints with 5s timeout.
func Shutdown() error {
	if stdSrv == nil { // not started yet
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return stdSrv.Shutdown(ctx)
}
