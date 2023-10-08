package service

import (
	"context"
	"testing"

	"github.com/b0rn/mkit/pkg/usecase"
)

type stubUseCase struct {
	calledGracefulShutdown bool
}

func (s *stubUseCase) GracefulShutdown(ctx context.Context) error {
	s.calledGracefulShutdown = true
	return nil
}

func TestGracefulShutdown(t *testing.T) {
	ctx := context.Background()
	svc := NewService()
	ucKeys := []string{"foo", "bar"}
	ucFactory := func(ctx context.Context, cfg interface{}) (usecase.UseCase, error) {
		return &stubUseCase{}, nil
	}
	svc.UseCaseManager.SetFactory(ucKeys[0], ucFactory)
	svc.UseCaseManager.SetFactory(ucKeys[1], ucFactory)
	uc1, _ := svc.UseCaseManager.Build(ctx, ucKeys[0], nil)
	stubUc1 := uc1.(*stubUseCase)
	uc2, _ := svc.UseCaseManager.Build(ctx, ucKeys[1], nil)
	stubUc2 := uc2.(*stubUseCase)
	if err := svc.GracefulShutdown(ctx); err != nil {
		t.Errorf("expected no error but got  : %v", err)
	}
	if !stubUc1.calledGracefulShutdown || !stubUc2.calledGracefulShutdown {
		t.Error("exepected GracefulShutdown functions to be called but some were never called")
	}
}
