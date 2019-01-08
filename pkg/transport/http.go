package transport

import (
	"context"
	"encoding/json"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phaneendrayeluri/gokit101/pkg/endpoint"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	m := http.NewServeMux()
	m.Handle("/add", httptransport.NewServer(
		endpoints.AddEndpoint,
		decodeAPIRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/sub", httptransport.NewServer(
		endpoints.SubEndpoint,
		decodeAPIRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/mul", httptransport.NewServer(
		endpoints.MulEndpoint,
		decodeAPIRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/div", httptransport.NewServer(
		endpoints.DivEndpoint,
		decodeAPIRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return m
}

// decodeHTTPSumRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded sum request from the HTTP request body. Primarily useful in a
// server.
func decodeAPIRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.APIRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError) // We can have a function the responds with a code given an error to handle multiple error code scenarios
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"error"`
}
