package pool

import (
        "code.google.com/p/vitess/go/pools"
        "github.com/garyburd/redigo/redis"
        "time"
)

/////////////////////////////
type RedisConn struct {
        redis.Conn
}

func (r *RedisConn) Close() {
        _ = r.Conn.Close()
}

/////////////////////////////
type RedisPool struct {
        pool *pools.ResourcePool
}

func NewPool(r *Redis, capacity int, maxCapacity int, idleTimout time.Duration) *pools.ResourcePool {
        return &RedisPool{pools.NewResourcePool(newRedisFactory(r), capacity, maxCapacity, idleTimout)}
}

func (rp *RedisPool) GetConn() (*RedisConn, error) {
        resource, err := rp.pool.Get()

        if err != nil {
                return nil, err
        }
        return resource.(*RedisConn), nil
}

func (rp *RedisPool) PutConn(conn *RedisConn) {
        rp.pool.Put(conn)
}

/////////////////////////////
func newRedisFactory(r *Redis) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(r)
        }
}

func connect(r *Redis) (*RedisConn, error) {
        cli, err := redis.Dial("tcp", r.Host + ":" + r.Port)
        if err != nil {
                return (nil, err)
        }
        return (cli, nil)
}
