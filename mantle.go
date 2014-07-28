package mantle

import (
        "github.com/vireshas/mantle/backends/redis"
)

type Mantle interface{
        Connect()
        Get(key string) string
        Set(key string, value interface{}) bool
        //MGet(key ...interface{}) map[interface{}]interface{}
        //MSet(k_v_map map[interface{}]interface{}) bool
        //Expire(keys ...interface{}) bool
}


type Orm struct{
        driver string
        Host string
        Port string
}

func (o *Orm) Get() Mantle{
        if o.driver == "" {
                return Mantle(&mantle.Redis{Host : "", Port : ""})
        }else{
                return Mantle(&mantle.Redis{Host : "", Port : ""})
        }
}
