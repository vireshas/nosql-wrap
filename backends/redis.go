package mantle

import (
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"strconv"
	"strings"
	"time"
)

//cant make these guys const as []string is not allowed in consts
var PoolSize = 10
var DefaultIpAndHost = []string{"localhost:6379"}

//This method creates a redis connection
func Connect(Instance interface{}) (pools.Resource, error) {
	//we need to read ip and db from this struct
	redis_struct := Instance.(*Redis)
	hostNPorts := redis_struct.Settings.HostAndPorts
	db := redis_struct.db

	//panic is more than 1 ip is given
	if len(hostNPorts) > 1 {
		panic("we can only connect to 1 server at the moment")
	}

	//dial host and port
	hostNPort := strings.Split(hostNPorts[0], ":")
	cli, err := redis.Dial("tcp", hostNPort[0]+":"+hostNPort[1])
	if err != nil {
		panic(err)
	}

	//select a redis db
	_, err = cli.Do("SELECT", db)
	if err != nil {
		panic(err)
	}

	//typecast to redisconn
	return &RedisConn{cli}, nil
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
	connection, err := r.pool.GetConn(r.Settings.Timeout)
	if err != nil {
		return nil, err
	}
	return connection.(*RedisConn), nil
}

func (r *Redis) PutClient(c *RedisConn) {
	r.pool.PutConn(c)
}

type Redis struct {
	Settings PoolSettings
	pool     *ResourcePool
	db       int
}

func (r *Redis) SetDefaults() {
	if len(r.Settings.HostAndPorts) == 0 {
		r.Settings.HostAndPorts = DefaultIpAndHost
	}
	//this is poolsize
	if r.Settings.Capacity == 0 {
		r.Settings.Capacity = PoolSize
	}
	//maxcapacity of the pool
	if r.Settings.MaxCapacity == 0 {
		r.Settings.MaxCapacity = PoolSize
	}
	//pool timeout
	r.Settings.Timeout = time.Minute

	//select a particular db in redis
	db, ok := r.Settings.Options["db"]
	if !ok {
		db = "0"
	}
	select_db, err := strconv.Atoi(db)
	if err != nil {
		panic("From Redis: select db is not a valid string")
	}
	r.db = select_db

	//create a pool finally
	r.pool = NewPool(Connect, r, r.Settings)
}

//Alias to SetDefaults
func (r *Redis) Configure(settings PoolSettings) {
	r.Settings = settings
	r.SetDefaults()
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
