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

//Get a client from pool
func (rp *ResourcePool) GetConn() (pools.Resource, error) {
        resource, err := rp.pool.Get()
        if err != nil {
                return nil, err
        }
        return resource, nil
}

//Put a client back to pool
func (rp *ResourcePool) PutConn(conn pools.Resource) {
        rp.pool.Put(conn)
}

//Redis pool wrapper
type ResourcePool struct {
        pool *pools.ResourcePool
}

type dialAndConnect func (host, port string) (pools.Resource, error)

//We create pool using NewPool
func NewPool(connect dialAndConnect, settings PoolSettings) *ResourcePool {
        return &ResourcePool{pools.NewResourcePool( newRedisFactory(connect, settings.Host, settings.Port),settings.Capacity,settings.MaxCapacity,settings.Timeout)}
}

//Helper methods for creating a pool
func newRedisFactory(connect dialAndConnect, host string, port string) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(host, port)
        }
}

