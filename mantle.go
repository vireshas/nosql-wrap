package mantle

import (
        "string"
        "./backends/redis"
        "./backends/memcache"
)

type Mantle interface{
        Connect()
        Get(key interface{}) interface{}
        Set(key interface{}, value interface{}) bool
        MGet(key ...interface{}) map[interface{}]interface{}
        MSet(k_v_map map[interface{}]interface{}) bool
        Expire(keys ...interface{}) bool
}

type Orm struct{
        driver string
        host string
        port string
}

func (o *Orm) Get() interface{}{
        if o.driver == "memcache" {
                return &Mantle{&Memcache{o.host, o.port}}
        }else{
                return &Mantle{&Redis{o.host, o.port}}
        }
}


func main(){
        driver := &Orm{}
        driver.Get()
        /*
        connection = driver.Connect()
        connection.Get("key") //returns value
        connection.Set("key", "value") //returns true or false
        connection.MGet(["key1", "key2", "key3"]) // returns map of k v
        connection.MSet(["key"]"value") //returns true or false
        */
}

