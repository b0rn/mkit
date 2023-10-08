package api

import (
	"context"
	"errors"

	"github.com/b0rn/mkit/pkg/factorymanager"
)

type Api interface {
	// Serve must be non-blocking function that serves the API
	Serve(ctx context.Context) error
	// To gracefully shutdown the API
	GracefulShutdown(ctx context.Context) error
}
type ApiFactory factorymanager.Factory[Api]
type ApiManager struct {
	*factorymanager.FactoryManager[Api]
}

// Return a new ApiManager
func NewManager() *ApiManager {
	m := factorymanager.NewFactoryManager[Api]()
	return &ApiManager{m}
}

// Sequentially calls the Serve function of every built API
// and returns the first error produced by any of those calls.
func (api *ApiManager) ServeAll(ctx context.Context) error {
	for _, v := range api.GetFactoryKeys() {
		a, ok := api.GetObject(v)
		if ok {
			if err := a.Serve(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// Sequentially calls the GracefulShutdown function of every built API
// and returns all errors wrapped into a single error.
func (api *ApiManager) GracefulShutdown(ctx context.Context) error {
	var err error
	for _, v := range api.GetFactoryKeys() {
		a, ok := api.GetObject(v)
		if ok {
			e := a.GracefulShutdown(ctx)
			if e != nil {
				err = errors.Join(err, e)
			}
		}
	}
	return err
}
