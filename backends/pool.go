package mantle

import (
        "github.com/youtube/vitess/go/pools"
        "time"
)

//params required to create a pool
type PoolSettings struct {
        Host string
        Port string
        Capacity int
        MaxCapacity int
        Timeout time.Duration
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

type dialAndConnect func (host, port string) (pools.Resource, error)

/**
        We create pool using NewPool
        dialAndConnect is the callback which creates a connection
*/
func NewPool(connect dialAndConnect, settings PoolSettings) *ResourcePool {
        return &ResourcePool{pools.NewResourcePool( newRedisFactory(connect, settings.Host, settings.Port),settings.Capacity,settings.MaxCapacity,settings.Timeout)}
}

func newRedisFactory(connect dialAndConnect, host string, port string) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(host, port)
        }
}

