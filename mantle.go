package mantle

import (
        "./backends/redis"
        //"./backends/memcache"
        "fmt"
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


/*
func main(){
        orm := &Orm{Host: "", Port: ""}
        driver := orm.Get()
        driver.Connect()
        driver.Set("key", "value")
        fmt.Println(driver.Get("key"))
        driver.Set("key", "value1")
        fmt.Println(driver.Get("key"))
        connection = driver.Connect()
        connection.Get("key") //returns value
        connection.Set("key", "value") //returns true or false
        connection.MGet(["key1", "key2", "key3"]) // returns map of k v
        connection.MSet(["key"]"value") //returns true or false
}
*/
