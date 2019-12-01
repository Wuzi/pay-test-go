// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/Wuzi/pay-test-go/restapi/operations"
	"github.com/Wuzi/pay-test-go/restapi/operations/city"
	"github.com/Wuzi/pay-test-go/restapi/operations/weather"
)

//go:generate swagger generate server --target ../../pay-test-go --name Paytest --spec ../swagger.yml

func configureFlags(api *operations.PaytestAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PaytestAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CityGetCitiesHandler == nil {
		api.CityGetCitiesHandler = city.GetCitiesHandlerFunc(func(params city.GetCitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation city.GetCities has not yet been implemented")
		})
	}
	if api.CityGetCitiesIDHandler == nil {
		api.CityGetCitiesIDHandler = city.GetCitiesIDHandlerFunc(func(params city.GetCitiesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation city.GetCitiesID has not yet been implemented")
		})
	}
	if api.WeatherGetCitiesIDWeathersHandler == nil {
		api.WeatherGetCitiesIDWeathersHandler = weather.GetCitiesIDWeathersHandlerFunc(func(params weather.GetCitiesIDWeathersParams) middleware.Responder {
			return middleware.NotImplemented("operation weather.GetCitiesIDWeathers has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
