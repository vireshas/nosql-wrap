package mantle

import (
	"github.com/goibibo/mantle/backends"
)

type Mantle interface {
	Get(key string) string
	Set(key string, value interface{}) bool
	Delete(keys ...interface{}) int
	Setex(key string, duration int, value interface{}) bool
	MGet(keys ...interface{}) []string
	MSet(keyValMap map[string]interface{}) bool
	Expire(key string, duration int) bool
	Execute(cmd string, args ...interface{}) (interface{}, error)
}

//this struct is exported
type Orm struct {
	//redis|memcache|cassandra
	Driver string
	//arrays of ip:port,ip:port
	HostAndPorts []string
	//pool size
	Capacity int
	//any other options thats needed for creating a connection
	Options map[string]string
}

func (o *Orm) New() Mantle {
	poolSettings := mantle.PoolSettings{
		HostAndPorts: o.HostAndPorts,
		Capacity:     o.Capacity,
		MaxCapacity:  o.Capacity,
		Options:      o.Options}

	if o.Driver == "memcache" {
		return RedisConns(poolSettings)
	} else {
		return RedisConns(poolSettings)
	}
}

func RedisConns(settings mantle.PoolSettings) *mantle.Redis {
	redis := &mantle.Redis{}
	redis.Configure(settings)
	return redis
}
