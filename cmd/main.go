package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/phaneendrayeluri/gokit101/pkg/endpoint"
	"github.com/phaneendrayeluri/gokit101/pkg/service"
	"github.com/phaneendrayeluri/gokit101/pkg/transport"
)

func main() {

	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters. The adapters, like
	// the HTTP handler or the gRPC server, are the bridge between Go kit and
	// the interfaces that the transports expect. Note that we're not binding
	// them to ports or anything yet; we'll do that next.
	var (
		service     = service.New(logger)
		endpoints   = endpoint.New(service, logger)
		httpHandler = transport.NewHTTPHandler(endpoints, logger)
	)

	logger.Log("accepting HTTP traffic on", ":8080")
	logger.Log("system exit reason", http.ListenAndServe(":8080", httpHandler))
}
