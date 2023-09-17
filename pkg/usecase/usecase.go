package usecase

import (
	"github.com/b0rn/mkit/internal/factorywcmanager"
	"github.com/b0rn/mkit/pkg/container"
)

type UseCase interface{}
type UseCaseFactory factorywcmanager.Factory[UseCase]
type UseCaseManager *factorywcmanager.FactoryManager[UseCase]

func NewManager(c container.Container) UseCaseManager {
	m := factorywcmanager.NewManager[UseCase](c)
	return m
}
