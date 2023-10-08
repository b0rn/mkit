package dataservice

import (
	"context"

	"github.com/b0rn/mkit/pkg/factorymanager"
)

type DataService interface {
	GracefulShutdown(ctx context.Context) error
}
type DataServiceFactory = factorymanager.Factory[DataService]
type DataServiceManager = *factorymanager.FactoryManager[DataService]

func NewManager() DataServiceManager {
	m := factorymanager.NewFactoryManager[DataService]()
	return m
}
