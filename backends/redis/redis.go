package mantle

import (
        "github.com/garyburd/redigo/redis"
        pool "github.com/vireshas/mantle/backends/redis/pool"
        "time"
	"fmt"
)

type Redis struct{
        Host string
        Port string
        Capacity int
        pool *pool.RedisPool
}

func (r *Redis) SetDefaults(){
        if r.Host == "" { r.Host = "localHost" }
        if r.Port == "" { r.Port = "6379" }
        if r.Capacity == 0 { r.Capacity = 10 }
        r.pool = pool.NewPool( r.Host, r.Port, r.Capacity, r.Capacity, time.Minute )
}

func (r *Redis) Configure(){
	fmt.Println("Configure")
       r.SetDefaults()
	fmt.Println("Configure TRUE")
}

func (r *Redis) GetClient() (*pool.RedisConn, error){
        return r.pool.GetConn()
}

func (r *Redis) PutClient(c *pool.RedisConn){
        r.pool.PutConn(c)
}

func (r *Redis) Get(key string) string{
        client, err := r.GetClient()
        if err != nil {
                return "Client Error"
        }
        defer r.PutClient(client)

        value, err := redis.String(client.Do("GET", key))
        if err != nil {
                return "Key not found!"
        }
        return value
}

func (r *Redis) Set(key string, value interface{}) bool{
        client, err := r.GetClient()
        if err != nil {
                return false 
        }
        defer r.PutClient(client)

        _, redis_err := client.Do("SET", key, value)
        if redis_err != nil {
                return false
        }
        return true
}


/*
func main(){

        r := &Redis{}
        r.Configure()
        fmt.Println(r.Get("colll1"))
        r.Set("colll1", 1)
        fmt.Println(r.Get("colll1"))
}


func (r *Redis) MGet(keys ..interface{}) map[interface{}]interface{}{

}

func (r *Redis) Get(k_v_map map[interface{}]interface{}) bool{

}

func (r *Redis) Expire(keys ...interface{}) bool{

}
*/
