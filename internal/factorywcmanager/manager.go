package factorywcmanager

import (
	"context"

	"github.com/b0rn/mkit/pkg/container"
	fm "github.com/b0rn/mkit/pkg/factorymanager"
)

type Factory[T interface{}] func(ctx context.Context, c container.Container, config interface{}) (T, error)

type FactoryManager[T interface{}] struct {
	container container.Container
	manager   *fm.FactoryManager[T]
}

type factoryParams struct {
	container container.Container
	config    interface{}
}

func NewManager[T interface{}](c container.Container) *FactoryManager[T] {
	return &FactoryManager[T]{
		container: c,
		manager:   fm.NewFactoryManager[T](),
	}
}

func (m *FactoryManager[T]) SetFactory(key string, factory Factory[T]) {
	f := func(ctx context.Context, params interface{}) (T, error) {
		p := params.(factoryParams)
		return factory(ctx, p.container, p.config)
	}
	m.manager.SetFactory(key, f)
}

func (m *FactoryManager[T]) GetFactory(key string) (T, bool) {
	return m.GetFactory(key)
}

func (m *FactoryManager[T]) GetFactoryKeys() []string {
	return m.GetFactoryKeys()
}

func (m *FactoryManager[T]) Build(ctx context.Context, key string, config interface{}) (*T, error) {
	return m.manager.Build(ctx, key, factoryParams{
		container: m.container,
		config:    config,
	})
}

func (m *FactoryManager[T]) BuildAll(ctx context.Context, config interface{}) []error {
	keys := m.manager.GetFactoryKeys()
	var errs []error
	for _, k := range keys {
		_, err := m.Build(ctx, k, config)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (m *FactoryManager[T]) Get(key string) (T, bool) {
	return m.manager.GetObject(key)
}
