package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Calculator) Calculator

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Calculator) Calculator {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Calculator
}

func (lm loggingMiddleware) Add(ctx context.Context, input []int) (int, error) {
	lm.logger.Log("Function", "Add Invoked", "Input", input)
	defer lm.logger.Log("Function", "Exiting Add")
	return lm.next.Add(ctx, input)
}
func (lm loggingMiddleware) Sub(ctx context.Context, input []int) (int, error) {
	lm.logger.Log("Function", "Sub Invoked", "Input", input)
	defer lm.logger.Log("Function", "Exiting Sub")
	return lm.next.Sub(ctx, input)
}
func (lm loggingMiddleware) Div(ctx context.Context, input []int) (int, error) {
	lm.logger.Log("Function", "Div Invoked", "Input", input)
	defer lm.logger.Log("Function", "Exiting Div")
	return lm.next.Div(ctx, input)
}
func (lm loggingMiddleware) Mul(ctx context.Context, input []int) (int, error) {
	lm.logger.Log("Function", "Mul Invoked", "Input", input)
	defer lm.logger.Log("Function", "Exiting Mul")
	return lm.next.Mul(ctx, input)
}
