package main

import (
        "github.com/garyburd/redigo/redis"
        "fmt"
        //"reflect"
)

type Redis struct{
        host string
        port string
        client redis.Conn
}

func (r *Redis) SetDefaults(){
        if r.host == "" { r.host = "localhost" }
        if r.port == "" { r.port = "6379" }
}

func (r *Redis) Connect(){
        //use default host and port
        r.SetDefaults()

        r.client, _ := redis.Dial("tcp", r.host + ":" + r.port)
}

func (r *Redis) Get(key string) string{
        value, err := redis.String(r.client.Do("GET", key))
        if err != nil {
                fmt.Println(err)
                return "Key not found!"
        }
        return value
}

func (r *Redis) Set(key string, value interface{}) bool{
        _, err := r.client.Do("SET", key, value)
        if err != nil {
                return false
        }
        return true
}

func main(){

        r := &Redis{}
        r.Connect()
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
