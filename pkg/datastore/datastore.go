package datastore

import "github.com/b0rn/mkit/pkg/factorymanager"

// Data stores are unknown types. They can be an sql connection, a mongo connection, etc...
type DataStore interface{}
type DataStoreFactory = factorymanager.Factory[DataStore]
type DataStoreManager = *factorymanager.FactoryManager[DataStore]

func NewManager() DataStoreManager {
	m := factorymanager.NewFactoryManager[DataStore]()
	return m
}
