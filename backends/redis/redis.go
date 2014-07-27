package mantle

import (
        "github.com/garyburd/redigo/redis"
)

type Redis struct{
        Host string
        Port string
        Client redis.Conn
}

func (r *Redis) SetDefaults(){
        if r.Host == "" { r.Host = "localHost" }
        if r.Port == "" { r.Port = "6379" }
}

func (r *Redis) Connect(){
        //use default Host and Port
        r.SetDefaults()

        cli, err := redis.Dial("tcp", r.Host + ":" + r.Port)
        if err != nil {
                cli = nil
        }
        r.Client = cli
}

func (r *Redis) Get(key string) string{
        value, err := redis.String(r.Client.Do("GET", key))
        if err != nil {
                //fmt.Println(err)
                return "Key not found!"
        }
        return value
}

func (r *Redis) Set(key string, value interface{}) bool{
        _, err := r.Client.Do("SET", key, value)
        if err != nil {
                return false
        }
        return true
}

/*
func main(){

        r := &Redis{}
        r.Connect()
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
