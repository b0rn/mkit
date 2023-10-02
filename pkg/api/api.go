package api

import (
	"context"
	"errors"

	"github.com/b0rn/mkit/pkg/factorymanager"
)

type Api interface {
	Serve(ctx context.Context) error
	GracefulShutdown(ctx context.Context) error
}
type ApiFactory factorymanager.Factory[Api]
type ApiManager struct {
	*factorymanager.FactoryManager[Api]
}

func NewManager() *ApiManager {
	m := factorymanager.NewFactoryManager[Api]()
	return &ApiManager{m}
}

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
