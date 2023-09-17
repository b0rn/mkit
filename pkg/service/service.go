package service

import (
	"github.com/b0rn/mkit/pkg/config"
	"github.com/b0rn/mkit/pkg/container"
	"github.com/b0rn/mkit/pkg/dataservice"
	"github.com/b0rn/mkit/pkg/datastore"
	"github.com/b0rn/mkit/pkg/log"
	"github.com/b0rn/mkit/pkg/usecase"
	"github.com/rs/zerolog"
)

type Service struct {
	Config             interface{}
	Container          container.Container
	DataStoreManager   datastore.DataStoreManager
	DataServiceManager dataservice.DataServiceManager
	UsecaseManager     usecase.UseCaseManager
}

func NewService() *Service {
	container := container.NewContainer()
	return &Service{
		Container:          container,
		DataStoreManager:   datastore.NewManager(container),
		DataServiceManager: dataservice.NewManager(container),
		UsecaseManager:     usecase.NewManager(container),
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

func (s *Service) EnableLogger(cfg log.Config, errorStackMarshaller log.ErrorStackMarshaler) zerolog.Logger {
	return log.Init(cfg, errorStackMarshaller)
}
