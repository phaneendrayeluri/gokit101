package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
)

// Calculator ...
type Calculator interface {
	Add(ctx context.Context, input []int) (int, error)
	Sub(ctx context.Context, input []int) (int, error)
	Div(ctx context.Context, input []int) (int, error)
	Mul(ctx context.Context, input []int) (int, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger) Calculator {
	var svc Calculator
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService() Calculator {
	return basicService{}
}

type basicService struct{}

func (bs basicService) Add(ctx context.Context, input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("No Input")
	}
	var sum int
	for _, i := range input {
		sum += i
	}
	return sum, nil
}
func (bs basicService) Sub(ctx context.Context, input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("No Input")
	}
	var sum int
	for _, i := range input {
		sum -= i
	}
	return sum, nil
}
func (bs basicService) Div(ctx context.Context, input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("No Input")
	}
	var sum int
	for _, i := range input {
		sum *= i
	}
	return sum, nil
}
func (bs basicService) Mul(ctx context.Context, input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("No Input")
	}
	var sum int
	for _, i := range input {
		sum /= i
	}
	return sum, nil
}
