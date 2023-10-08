package usecase

import (
	"context"

	"github.com/b0rn/mkit/pkg/factorymanager"
)

type UseCase interface {
	GracefulShutdown(ctx context.Context) error
}
type UseCaseFactory = factorymanager.Factory[UseCase]
type UseCaseManager = *factorymanager.FactoryManager[UseCase]

func NewManager() UseCaseManager {
	m := factorymanager.NewFactoryManager[UseCase]()
	return m
}
