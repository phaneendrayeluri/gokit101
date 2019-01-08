package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/phaneendrayeluri/gokit101/pkg/service"
)

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	AddEndpoint endpoint.Endpoint
	SubEndpoint endpoint.Endpoint
	MulEndpoint endpoint.Endpoint
	DivEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Calculator, logger log.Logger) Set {
	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = MakeAddEndpoint(svc)
		addEndpoint = LoggingMiddleware(log.With(logger, "method", "Sum"))(addEndpoint)
	}
	var subEndpoint endpoint.Endpoint
	{
		subEndpoint = MakeSubEndpoint(svc)
		subEndpoint = LoggingMiddleware(log.With(logger, "method", "Sub"))(subEndpoint)
	}
	var mulEndpoint endpoint.Endpoint
	{
		mulEndpoint = MakeMulEndpoint(svc)
		mulEndpoint = LoggingMiddleware(log.With(logger, "method", "Mul"))(mulEndpoint)
	}
	var divEndpoint endpoint.Endpoint
	{
		divEndpoint = MakeDivEndpoint(svc)
		divEndpoint = LoggingMiddleware(log.With(logger, "method", "Div"))(divEndpoint)
	}

	return Set{
		AddEndpoint: addEndpoint,
		SubEndpoint: subEndpoint,
		MulEndpoint: mulEndpoint,
		DivEndpoint: divEndpoint,
	}
}

// APIRequest ...
type APIRequest []int

// APIResponse ...
type APIResponse struct {
	V   int
	Err error
}

// MakeAddEndpoint constructs a Sum endpoint wrapping the service.
func MakeAddEndpoint(s service.Calculator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(APIRequest)
		v, err := s.Add(ctx, req)
		return APIResponse{V: v, Err: err}, nil
	}
}

// MakeSubEndpoint constructs a Sum endpoint wrapping the service.
func MakeSubEndpoint(s service.Calculator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(APIRequest)
		v, err := s.Sub(ctx, req)
		return APIResponse{V: v, Err: err}, nil
	}
}

// MakeMulEndpoint constructs a Sum endpoint wrapping the service.
func MakeMulEndpoint(s service.Calculator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(APIRequest)
		v, err := s.Mul(ctx, req)
		return APIResponse{V: v, Err: err}, nil
	}
}

// MakeDivEndpoint constructs a Sum endpoint wrapping the service.
func MakeDivEndpoint(s service.Calculator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(APIRequest)
		v, err := s.Div(ctx, req)
		return APIResponse{V: v, Err: err}, nil
	}
}
