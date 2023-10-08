package api

import (
	"context"
	"testing"
)

type stubApi struct {
	calledServe            bool
	calledGracefulShutdown bool
}

func (s *stubApi) Serve(ctx context.Context) error {
	s.calledServe = true
	return nil
}

func (s *stubApi) GracefulShutdown(context.Context) error {
	s.calledGracefulShutdown = true
	return nil
}

func TestApiManager(t *testing.T) {
	ctx := context.Background()
	m := NewManager()
	keys := []string{"foo", "bar"}
	factory := func(ctx context.Context, cfg interface{}) (Api, error) {
		return &stubApi{}, nil
	}
	m.SetFactory(keys[0], factory)
	m.SetFactory(keys[1], factory)
	api1, _ := m.Build(ctx, keys[0], nil)
	api2, _ := m.Build(ctx, keys[1], nil)
	stubApi1 := api1.(*stubApi)
	stubApi2 := api2.(*stubApi)

	if err := m.ServeAll(ctx); err != nil {
		t.Errorf("expected no error when running ServeAll but got %v", err)
	}
	if !stubApi1.calledServe || !stubApi2.calledServe {
		t.Error("expected Serve functions to be called but some were never called")
	}

	if err := m.GracefulShutdown(ctx); err != nil {
		t.Errorf("expexted no error when running GracefulShutdown but got %v", err)
	}
	if !stubApi1.calledGracefulShutdown || !stubApi2.calledGracefulShutdown {
		t.Error("expected GracefulShutdown functions to be called but some were never called")
	}
}
