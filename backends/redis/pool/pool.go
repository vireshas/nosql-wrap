package mantle

import (
        "github.com/youtube/vitess/go/pools"
        "github.com/garyburd/redigo/redis"
        "time"
	"fmt"
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

func NewPool(host string, port string, capacity int, maxCapacity int, idleTimout time.Duration) *RedisPool {
        return &RedisPool{pools.NewResourcePool(newRedisFactory(host, port), capacity, maxCapacity, idleTimout)}
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
func newRedisFactory(host string, port string) pools.Factory {
        return func() (pools.Resource, error) {
                return connect(host, port)
        }
}

func connect(host string, port string) (*RedisConn, error) {
        cli, err := redis.Dial("tcp", host + ":" + port)
        if err != nil {
		fmt.Println("Error CONNECTING TO DB")
                return nil, err
        }
		fmt.Println("SUCCESSFUL")
        return &RedisConn{cli}, nil
}
