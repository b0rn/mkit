package datastore

import (
	"github.com/b0rn/mkit/internal/factorywcmanager"
	"github.com/b0rn/mkit/pkg/container"
)

type DataStore interface{}
type DataStoreFactory = factorywcmanager.Factory[DataStore]
type DataStoreManager = *factorywcmanager.FactoryManager[DataStore]

func NewManager(c container.Container) DataStoreManager {
	m := factorywcmanager.NewManager[DataStore](c)
	return m
}
