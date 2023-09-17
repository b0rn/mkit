package Api

import (
	"github.com/b0rn/mkit/internal/factorywcmanager"
	"github.com/b0rn/mkit/pkg/container"
)

type Api interface {
	Serve() error
}
type ApiFactory factorywcmanager.Factory[Api]
type ApiManager struct {
	*factorywcmanager.FactoryManager[Api]
}

func NewManager(c container.Container) *ApiManager {
	m := factorywcmanager.NewManager[Api](c)
	return &ApiManager{m}
}

func (api *ApiManager) ServeAll() error {
	for _, v := range api.GetFactoryKeys() {
		a, ok := api.Get(v)
		if ok {
			if err := a.Serve(); err != nil {
				return err
			}
		}
	}
	return nil
}
