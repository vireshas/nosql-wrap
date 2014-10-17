package mantle

import (
	"github.com/youtube/vitess/go/pools"
	"time"
)

type PoolSettings struct {
	HostAndPorts []string
	Capacity     int
	MaxCapacity  int
	Options      map[string]string
	Timeout      time.Duration
}

func (rp *ResourcePool) GetConn(to time.Duration) (pools.Resource, error) {
	resource, err := rp.pool.Get(to)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (rp *ResourcePool) PutConn(conn pools.Resource) {
	rp.pool.Put(conn)
}

type ResourcePool struct {
	pool *pools.ResourcePool
}

//The connect callback that is passed by a backend
type dialAndConnect func(instance interface{}) (pools.Resource, error)

func NewPool(connect dialAndConnect, instance interface{}, settings PoolSettings) *ResourcePool {
	return &ResourcePool{
		pools.NewResourcePool(newRedisFactory(connect, instance),
			settings.Capacity, settings.MaxCapacity, settings.Timeout)}
}

func newRedisFactory(connect dialAndConnect, instance interface{}) pools.Factory {
	return func() (pools.Resource, error) {
		return connect(instance)
	}
}
