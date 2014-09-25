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
		panic("we can only connect to 1 servem at the moment")
	}

	fmt.Println("connecting to ", hostNPorts[0])
	cli := memcache.New(hostNPorts[0])

	return &MemConn{cli}, nil
}

type MemConn struct {
	*memcache.Client
}

func (m *MemConn) Close() {
}

type Memcache struct {
	Settings PoolSettings
	pool     *ResourcePool
}

func (m *Memcache) GetClient() *MemConn {
	connection, err := m.pool.GetConn(m.Settings.Timeout)
	if err != nil {
		panic(err)
	}
	return connection.(*MemConn)
}

func (m *Memcache) PutClient(c *MemConn) {
	m.pool.PutConn(c)
}

func (m *Memcache) SetDefaults() {
	if len(m.Settings.HostAndPorts) == 0 {
		m.Settings.HostAndPorts = DefaultMemcacheIpAndHost
	}
	if m.Settings.Capacity == 0 {
		m.Settings.Capacity = PoolSize
	}
	if m.Settings.MaxCapacity == 0 {
		m.Settings.MaxCapacity = PoolSize
	}
	m.Settings.Timeout = time.Minute
	m.pool = NewPool(MConnect, m, m.Settings)
}

func (m *Memcache) Configure(settings PoolSettings) {
	m.Settings = settings
	m.SetDefaults()
}

func (m *Memcache) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return "inside GEt", nil
}

func (m *Memcache) Delete(keys ...interface{}) int {
	return 1
}

func (m *Memcache) Get(key string) string {
	mc := m.GetClient()
	it, erm := mc.Get(key)
	m.PutClient(mc)
	if erm != nil {
		errMsg := fmt.Sprintf("Errom getting key %s", key)
		return errMsg
	}
	return string(it.Value)
}

func (m *Memcache) Set(key string, value interface{}) bool {
	mc := m.GetClient()
	newVal := value.(string)
	erm := mc.Set(&memcache.Item{Key: key, Value: []byte(newVal)})
	m.PutClient(mc)
	if erm != nil {
		return false
	}
	return true
}

func (m *Memcache) MGet(keys ...interface{}) []string {
	return []string{"hello world"}
}

func (m *Memcache) MSet(mapOfKeyVal map[string]interface{}) bool {
	return true
}

func (m *Memcache) Expire(key string, duration int) bool {
	return true
}

func (m *Memcache) Setex(key string, duration int, val interface{}) bool {
	return true
}
