package container

type Container interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

func NewContainer() Container {
	return &container{
		values: make(map[string]interface{}),
	}
}

type container struct {
	values map[string]interface{}
}

func (c *container) Get(key string) (interface{}, bool) {
	v, ok := c.values[key]
	return v, ok
}

func (c *container) Set(key string, value interface{}) {
	c.values[key] = value
}
