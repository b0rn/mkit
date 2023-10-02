package datastore

import "github.com/b0rn/mkit/pkg/factorymanager"

type DataStore interface{}
type DataStoreFactory = factorymanager.Factory[DataStore]
type DataStoreManager = *factorymanager.FactoryManager[DataStore]

func NewManager() DataStoreManager {
	m := factorymanager.NewFactoryManager[DataStore]()
	return m
}
