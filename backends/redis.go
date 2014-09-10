package mantle

import (
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"strings"
	"time"
)

//cant make these guys const as []string is not allowed in consts
var PoolSize = 10
var DefaultIpAndHost = []string{"localhost:6379"}

type Redis struct {
	Settings PoolSettings
	pool     *ResourcePool
}

func (r *Redis) SetDefaults() {
	if len(r.Settings.IpAndHosts) == 0 {
		r.Settings.IpAndHosts = DefaultIpAndHost
	}
	if r.Settings.Capacity == 0 {
		r.Settings.Capacity = PoolSize
	}
	if r.Settings.MaxCapacity == 0 {
		r.Settings.MaxCapacity = PoolSize
	}
	r.Settings.Timeout = time.Minute
	r.pool = NewPool(Connect, r.Settings)
}

//Alias to SetDefaults
func (r *Redis) Configure() {
	r.SetDefaults()
}

//Wrapping redis connection
type RedisConn struct {
	redis.Conn
}

//Close a redis connection
func (r *RedisConn) Close() {
	_ = r.Conn.Close()
}

func (r *Redis) GetClient() (*RedisConn, error) {
	connection, err := r.pool.GetConn()
	if err != nil {
		return nil, err
	}
	return connection.(*RedisConn), nil
}

func (r *Redis) PutClient(c *RedisConn) {
	r.pool.PutConn(c)
}

//This method creates a redis connection
func Connect(IpAndHost []string) (pools.Resource, error) {
	if len(IpAndHost) > 1 {
		panic("we can only connect to 1 server at the moment")
	}
	hostNPort := strings.Split(IpAndHost[0], ":")
	cli, err := redis.Dial("tcp", hostNPort[0]+":"+hostNPort[1])
	if err != nil {
		panic(err)
	}
	return &RedisConn{cli}, nil
}

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
	if err != nil {
		return -1
	}
	return value
}

func (r *Redis) Get(key string) string {
	value, err := redis.String(r.Execute("GET", key))
	if err != nil {
		return ""
	}
	return value
}

func (r *Redis) Set(key string, value interface{}) bool {
	_, err := r.Execute("SET", key, value)
	if err != nil {
		return false
	}
	return true
}

func (r *Redis) MGet(keys ...interface{}) []string {
	values, err := redis.Strings(r.Execute("MGET", keys...))
	if err != nil {
		return []string{}
	}
	return values
}

func (r *Redis) MSet(mapOfKeyVal map[string]interface{}) bool {
	_, err := r.Execute("MSET", redis.Args{}.AddFlat(mapOfKeyVal)...)
	if err != nil {
		return false
	}
	return true
}

func (r *Redis) Expire(key string, duration int) bool {
	_, err := r.Execute("EXPIRE", key, duration)
	if err != nil {
		return false
	}
	return true
}

func (r *Redis) Setex(key string, duration int, val interface{}) bool {
	_, err := r.Execute("SETEX", key, duration, val)
	if err != nil {
		return false
	}
	return true
}
