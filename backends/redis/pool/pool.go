package mantle

import (
        "github.com/youtube/vitess/go/pools"
        "github.com/garyburd/redigo/redis"
        "time"
)

//Wrapping redis connection
type RedisConn struct {
        redis.Conn
}

//Close a redis connection
func (r *RedisConn) Close() {
        _ = r.Conn.Close()
}

//Redis pool wrapper
type RedisPool struct {
        pool *pools.ResourcePool
}

//We create pool using NewPool
func NewPool(host string, port string, capacity int, maxCapacity int, idleTimout time.Duration) *RedisPool {
        return &RedisPool{pools.NewResourcePool(newRedisFactory(host, port), capacity, maxCapacity, idleTimout)}
}

//Get a client from pool
func (rp *RedisPool) GetConn() (*RedisConn, error) {
        resource, err := rp.pool.Get()

        if err != nil {
                return nil, err
        }
        return resource.(*RedisConn), nil
}

//Put a client back to pool
func (rp *RedisPool) PutConn(conn *RedisConn) {
        rp.pool.Put(conn)
}

//Helper methods for creating a pool
func newRedisFactory(host string, port string) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(host, port)
        }
}

//This method actually connects to redis
func connect(host string, port string) (*RedisConn, error) {
        cli, err := redis.Dial("tcp", host + ":" + port)
        if err != nil {
                return nil, err
        }
        return &RedisConn{cli}, nil
}
