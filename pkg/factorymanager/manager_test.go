package factorymanager

import (
	"context"
	"errors"
	"testing"
)

func TestManager(t *testing.T) {
	m := NewFactoryManager[error]()
	keys := []string{"foo", "bar"}
	errFactory := errors.New("test error")
	keyErr := "err"
	calledBuild := false
	factory := func(ctx context.Context, params interface{}) (error, error) {
		calledBuild = true
		return errors.New("test"), nil
	}
	factoryErr := func(ctx context.Context, params interface{}) (error, error) {
		return nil, errFactory
	}
	m.SetFactory(keys[0], factory)
	m.SetFactory(keys[1], factory)

	fKeys := m.GetFactoryKeys()
	if len(fKeys) != len(keys) {
		t.Errorf("expected %v keys but got %v keys", len(keys), len(fKeys))
	}

	if _, ok := m.GetFactory(keys[0]); !ok {
		t.Errorf("expected key %s to be found", keys[0])
	}

	if err := m.BuildAll(context.Background(), nil); err != nil {
		t.Errorf("expected no error when running BuildAll but got %v", err)
	}
	if !calledBuild {
		t.Error("factory has not been called")
	}

	obj1, ok := m.GetObject(keys[0])
	if !ok {
		t.Errorf("could not get object with key %s", keys[0])
	}
	obj2, err := m.Build(context.Background(), keys[0], nil)
	if err != nil {
		t.Errorf("expected no error when running Build but got : %v", err)
	}
	if obj1 != obj2 {
		t.Errorf("expected %v but got %v", obj1, obj2)
	}

	m.SetFactory(keyErr, factoryErr)
	if _, e := m.Build(context.Background(), "invalid", nil); e != ErrFactoryNotFound {
		t.Error("must return an error when key is not found")
	}

	if _, e := m.Build(context.Background(), keyErr, nil); e != errFactory {
		t.Errorf("expected Build error to be %v but got %v", errFactory, e)
	}

	if errs := m.BuildAll(context.Background(), nil); len(errs) != 1 || errs[0] != errFactory {
		t.Errorf("expected BuildAll to only result in error %v but got %v", errFactory, errs)
	}
}
