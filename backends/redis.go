package mantle

import (
        "github.com/youtube/vitess/go/pools"
        "github.com/garyburd/redigo/redis"
        "time"
)

const (
        PoolSize = 10
        DefaultHost = "localhost"
        DefaultPort = "6379"
)

type Redis struct {
        Host string
        Port string
        Capacity int
        pool *ResourcePool
}

func (r *Redis) SetDefaults() {
        if r.Host == "" { r.Host = DefaultHost }
        if r.Port == "" { r.Port = DefaultPort }
        if r.Capacity == 0 { r.Capacity = PoolSize }
        r.pool = NewPool( Connect, r.Host, r.Port, r.Capacity, r.Capacity, time.Minute )
}


/*
        Creating a redis pool
*/


//Wrapping redis connection
type RedisConn struct {
        redis.Conn
}

//Close a redis connection
func (r *RedisConn) Close() {
        _ = r.Conn.Close()
}

//Alias to SetDefaults
func (r *Redis) Configure() {
        r.SetDefaults()
}

func (r *Redis) GetClient() (*RedisConn, error) {
        connection, err := r.pool.GetConn()
        if err != nil{
                return nil, err
        }
        return connection.(*RedisConn), nil
}

func (r *Redis) PutClient(c *RedisConn) {
        r.pool.PutConn(c)
}

//This method actually connects to redis
func Connect(host string, port string) (pools.Resource, error) {
        cli, err := redis.Dial("tcp", host + ":" + port)
        if err != nil {
                return nil, err
        }
        return &RedisConn{cli}, nil
}

/*
        Wrappers on top of redis
*/

//Generic method to execute any redis call
func (r *Redis) Execute(cmd string, args ...interface{}) (interface{}, error) {
        client, err := r.GetClient()
        if err != nil {
                return nil, err
        }
        defer r.PutClient(client)
        return client.Do(cmd, args...)
}

func (r *Redis) Delete(keys ...interface{}) int {
        value, err := redis.Int(r.Execute("DEL", keys...))
        if err != nil { return -1 }
        return value
}

func (r *Redis) Get(key string) string {
        value, err := redis.String(r.Execute("GET", key))
        if err != nil { return "" }
        return value
}

func (r *Redis) Set(key string, value interface{}) bool {
        _, err := r.Execute("SET", key, value)
        if err != nil { return false }
        return true
}

func (r *Redis) MGet(keys ...interface{}) []string {
        values, err := redis.Strings(r.Execute("MGET", keys...))
        if err != nil { return []string{} }
        return values
}

func (r *Redis) MSet(mapOfKeyVal map[string]interface{}) bool {
        _, err := r.Execute("MSET", redis.Args{}.AddFlat(mapOfKeyVal)...)
        if err != nil { return false }
        return true
}

func (r *Redis) Expire(key string, duration int) bool {
        _, err := r.Execute("EXPIRE", key, duration)
        if err != nil { return false }
        return true
}

func (r *Redis) Setex(key string, duration int, val interface{}) bool {
        _, err := r.Execute("SETEX", key, duration, val)
        if err != nil { return false }
        return true
}

