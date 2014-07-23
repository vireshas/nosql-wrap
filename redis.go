package nosqlorm

import (
        "github.com/garyburd/redigo/redis"
)

type Redis struct{
        host string
        port string
}

func (r *Redis) Connect(){

}

func (r *Redis) Get(key interface{}) interface{}{

}

func (r *Redis) Set(key interface{}, value interface{}) bool{

}

func (r *Redis) MGet(keys ..interface{}) map[interface{}]interface{}{

}

func (r *Redis) Get(k_v_map map[interface{}]interface{}) bool{

}

func (r *Redis) Expire(keys ...interface{}) bool{

}
