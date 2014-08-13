package mantle

import (
        "github.com/vireshas/mantle/backends"
)

type Mantle interface {
        Get(key string) string
        Set(key string, value interface{}) bool
        Delete(keys ...interface{}) int
        Setex(key string, duration int, value interface{}) bool
        MGet(key ...interface{}) []string
        MSet(keyValMap map[string]interface{}) bool
        Expire(key string, duration int) bool
        Execute(cmd string, args ...interface{}) (interface{}, error)
}

type Orm struct {
        Driver string
        Host string
        Port string
        Capacity int
}

func (o *Orm) Get() Mantle {
        if o.Driver == "memcache" {
                return Mantle(&mantle.Redis{Host:o.Host, Port:o.Port, Capacity:o.Capacity})
        }else{
		redis := &mantle.Redis{Host:o.Host, Port:o.Port, Capacity:o.Capacity}
		redis.Configure()
                return Mantle(redis)
        }
}
