package mantle

import (
	"github.com/youtube/vitess/go/pools"
	"time"
)

type PoolSettings struct {
	IpAndHosts  []string
	Capacity    int
	MaxCapacity int
	Timeout     time.Duration
}

func (rp *ResourcePool) GetConn() (pools.Resource, error) {
	resource, err := rp.pool.Get()
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
type dialAndConnect func(ipHosts []string) (pools.Resource, error)

func NewPool(connect dialAndConnect, settings PoolSettings) *ResourcePool {
	return &ResourcePool{
		pools.NewResourcePool(newRedisFactory(connect, settings.IpAndHosts),
			settings.Capacity, settings.MaxCapacity, settings.Timeout)}
}

func newRedisFactory(connect dialAndConnect, IpAndHosts []string) pools.Factory {
	return func() (pools.Resource, error) {
		return connect(IpAndHosts)
	}
}
