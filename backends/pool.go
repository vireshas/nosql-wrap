package mantle

import (
        "github.com/youtube/vitess/go/pools"
        "time"
)

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
func NewPool(connect dialAndConnect, host string, port string, capacity int, maxCapacity int, idleTimout time.Duration) *ResourcePool {
        return &ResourcePool{pools.NewResourcePool(newRedisFactory(connect, host, port), capacity, maxCapacity, idleTimout)}
}
//Helper methods for creating a pool
func newRedisFactory(connect dialAndConnect, host string, port string) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(host, port)
        }
}

