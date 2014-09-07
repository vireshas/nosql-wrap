package mantle

import (
	"github.com/vireshas/mantle/backends"
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

type Orm struct {
	Driver       string
	HostAndPorts []string
	Capacity     int
}

func (o *Orm) Get() Mantle {
	if o.Driver == "memcache" {
		redis := &mantle.Redis{
			Settings: mantle.PoolSettings{
				HostAndPorts: o.HostAndPorts,
				Capacity:     o.Capacity,
				MaxCapacity:  o.Capacity}}
		redis.Configure()
		return Mantle(redis)
	} else {
		redis := &mantle.Redis{
			Settings: mantle.PoolSettings{
				HostAndPorts: o.HostAndPorts,
				Capacity:     o.Capacity,
				MaxCapacity:  o.Capacity}}
		redis.Configure()
		return Mantle(redis)
	}
}
