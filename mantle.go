package mantle

import (
        "github.com/vireshas/mantle/backends/redis"
	"fmt"
)

type Mantle interface{
	//Configure()
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
		redis := &mantle.Redis{Host : "", Port : ""}
		fmt.Printf("HELLO  %#s", redis)
		redis.Configure()
                return Mantle(&redis)
        }else{
                return Mantle(&mantle.Redis{Host : "", Port : ""})
        }
}
