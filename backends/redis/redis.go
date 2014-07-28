package main

import (
        "github.com/garyburd/redigo/redis"
        "github.com/vireshas/mantle/backends/redis/pool"
        "time"
)

type Redis struct{
        Host string
        Port string
        Capacity int
        pool redis_pool.RedisPool
}

func (r *Redis) SetDefaults(){
        if r.Host == "" { r.Host = "localHost" }
        if r.Port == "" { r.Port = "6379" }
        if r.Capacity == 0 { r.Capacity = 10 }
        if r.pool == nil { pool.NewPool( r, r.Capacity, r.Capacity, time.Minute ) }
}

func (r *Redis) Configure(){
       r.SetDefaults()
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
        defer r.pool.PutClient(client)

        value, err := redis.String(client.Do("GET", key))
        if err != nil {
                return "Key not found!"
        }
        return value
}

func (r *Redis) Set(key string, value interface{}) bool{
        client, err := r.pool.GetClient()
        if err != nil {
                return "Client Error"
        }
        defer r.pool.PutClient(client)

        _, err := client.Do("SET", key, value)
        if err != nil {
                return false
        }
        return true
}


func main(){

        r := &Redis{}
        r.Configure()
        fmt.Println(r.Get("colll1"))
        r.Set("colll1", 1)
        fmt.Println(r.Get("colll1"))
}

/*

func (r *Redis) MGet(keys ..interface{}) map[interface{}]interface{}{

}

func (r *Redis) Get(k_v_map map[interface{}]interface{}) bool{

}

func (r *Redis) Expire(keys ...interface{}) bool{

}
*/
