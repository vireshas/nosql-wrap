package mantle

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/vireshas/minimal_vitess_pool/pools"
	"time"
)

var MemcachePoolSize = 10
var DefaultMemcacheIpAndHost = []string{"localhost:11211"}

func MConnect(Instance interface{}) (pools.Resource, error) {
	redisInstance := Instance.(*Memcache)
	hostNPorts := redisInstance.Settings.HostAndPorts

	if len(hostNPorts) > 1 {
		panic("we can only connect to 1 server at the moment")
	}

	fmt.Println("connecting to ", hostNPorts[0])
	cli := memcache.New(hostNPorts[0])

	return &MemConn{cli}, nil
}

type MemConn struct {
	*memcache.Client
}

func (r *MemConn) Close() {
}

type Memcache struct {
	Settings PoolSettings
	pool     *ResourcePool
}

func (r *Memcache) GetClient() *MemConn {
	connection, err := r.pool.GetConn(r.Settings.Timeout)
	if err != nil {
		panic(err)
	}
	return connection.(*MemConn)
}

func (r *Memcache) PutClient(c *MemConn) {
	r.pool.PutConn(c)
}

func (r *Memcache) SetDefaults() {
	if len(r.Settings.HostAndPorts) == 0 {
		r.Settings.HostAndPorts = DefaultMemcacheIpAndHost
	}
	if r.Settings.Capacity == 0 {
		r.Settings.Capacity = PoolSize
	}
	if r.Settings.MaxCapacity == 0 {
		r.Settings.MaxCapacity = PoolSize
	}
	r.Settings.Timeout = time.Minute
	r.pool = NewPool(MConnect, r, r.Settings)
}

func (r *Memcache) Configure(settings PoolSettings) {
	r.Settings = settings
	r.SetDefaults()
}

func (r *Memcache) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return "inside GEt", nil
}

func (r *Memcache) Delete(keys ...interface{}) int {
	return 1
}

func (r *Memcache) Get(key string) string {
	mc := r.GetClient()
	it, err := mc.Get(key)
	r.PutClient(mc)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting key %s", key)
		return errMsg
	}
	return string(it.Value)
}

func (r *Memcache) Set(key string, value interface{}) bool {
	mc := r.GetClient()
	newVal := value.(string)
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(newVal)})
	r.PutClient(mc)
	if err != nil {
		return false
	}
	return true
}

func (r *Memcache) MGet(keys ...interface{}) []string {
	return []string{"hello world"}
}

func (r *Memcache) MSet(mapOfKeyVal map[string]interface{}) bool {
	return true
}

func (r *Memcache) Expire(key string, duration int) bool {
	return true
}

func (r *Memcache) Setex(key string, duration int, val interface{}) bool {
	return true
}
