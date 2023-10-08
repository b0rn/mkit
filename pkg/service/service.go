package service

import (
	"context"
	"errors"

	"github.com/b0rn/mkit/pkg/api"
	"github.com/b0rn/mkit/pkg/config"
	"github.com/b0rn/mkit/pkg/container"
	"github.com/b0rn/mkit/pkg/dataservice"
	"github.com/b0rn/mkit/pkg/datastore"
	"github.com/b0rn/mkit/pkg/mlog"
	"github.com/b0rn/mkit/pkg/usecase"
	"github.com/rs/zerolog"
)

type Service struct {
	Config             interface{}
	Container          container.Container
	DataStoreManager   datastore.DataStoreManager
	DataServiceManager dataservice.DataServiceManager
	UseCaseManager     usecase.UseCaseManager
	ApiManager         api.ApiManager
}

func NewService() *Service {
	container := container.NewContainer()
	return &Service{
		Container:          container,
		DataStoreManager:   datastore.NewManager(),
		DataServiceManager: dataservice.NewManager(),
		UseCaseManager:     usecase.NewManager(),
		ApiManager:         *api.NewManager(),
	}
}

func (s *Service) LoadEnvVars(filetype string, filepath string) error {
	return config.LoadEnvVars(filetype, filepath)
}

func (s *Service) BuildConfig(filetype string, filepath string, cfg interface{}) error {
	err := config.BuildConfig(filetype, filepath, cfg)
	s.Config = cfg
	return err
}

func (s *Service) EnableLogger(cfg mlog.Config, errorStackMarshaller mlog.ErrorStackMarshaler) zerolog.Logger {
	return mlog.Init(cfg, errorStackMarshaller)
}

func (s *Service) GracefulShutdown(ctx context.Context) error {
	var err error
	for _, v := range s.UseCaseManager.GetFactoryKeys() {
		if u, ok := s.UseCaseManager.GetObject(v); ok {
			if e := u.GracefulShutdown(ctx); e != nil {
				err = errors.Join(err, e)
			}
		}
	}
	if e := s.ApiManager.GracefulShutdown(ctx); e != nil {
		err = errors.Join(err, e)
	}
	return err
}
