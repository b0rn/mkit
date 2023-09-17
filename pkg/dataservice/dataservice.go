package dataservice

import (
	"github.com/b0rn/mkit/internal/factorywcmanager"
	"github.com/b0rn/mkit/pkg/container"
)

type DataService interface{}
type DataServiceFactory = factorywcmanager.Factory[DataService]
type DataServiceManager = *factorywcmanager.FactoryManager[DataService]

func NewManager(c container.Container) DataServiceManager {
	m := factorywcmanager.NewManager[DataService](c)
	return m
}
