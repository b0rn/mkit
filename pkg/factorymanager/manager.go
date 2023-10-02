package factorymanager

import (
	"context"
)

type Factory[T interface{}] func(ctx context.Context, params interface{}) (T, error)

type FactoryManager[T interface{}] struct {
	factories map[string]Factory[T]
	concretes map[string]T
}

func NewFactoryManager[T interface{}]() *FactoryManager[T] {
	return &FactoryManager[T]{
		factories: make(map[string]Factory[T]),
		concretes: make(map[string]T),
	}
}

func (m *FactoryManager[T]) GetFactoryKeys() []string {
	var keys []string
	for k := range m.factories {
		keys = append(keys, k)
	}
	return keys
}

func (m *FactoryManager[T]) GetFactory(key string) (Factory[T], bool) {
	f, ok := m.factories[key]
	return f, ok
}

func (m *FactoryManager[T]) SetFactory(key string, factory Factory[T]) {
	m.factories[key] = factory
}

func (m *FactoryManager[T]) Build(ctx context.Context, key string, params interface{}) (T, error) {
	obj, ok := m.concretes[key]
	if ok {
		return obj, nil
	}
	f, ok := m.GetFactory(key)
	if !ok {
		return obj, ErrFactoryNotFound
	}
	obj, err := f(ctx, params)
	if err != nil {
		return obj, err
	}
	m.concretes[key] = obj
	return obj, nil
}

func (m *FactoryManager[T]) BuildAll(ctx context.Context, params interface{}) []error {
	var errs []error
	for k, _ := range m.factories {
		_, err := m.Build(ctx, k, params)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (m *FactoryManager[T]) GetObject(key string) (T, bool) {
	o, ok := m.concretes[key]
	return o, ok
}
