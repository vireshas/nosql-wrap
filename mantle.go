package nosqlorm

type Mantle interface{
        Connect()
        Get(key interface{}) interface{}
        Set(key interface{}, value interface{}) bool
        MGet(key ...interface{}) map[interface{}]interface{}
        MSet(k_v_map map[interface{}]interface{}) bool
        Expire(keys ...interface{}) bool
}


/*

func main(){
        driver := Mantle{Redis{host:"localhost", port:"6379"}}
        connection = driver.Connect()
        connection.Get("key") //returns value
        connection.Set("key", "value") //returns true or false
        connection.MGet(["key1", "key2", "key3"]) // returns map of k v
        connection.MSet(["key"]"value") //returns true or false
}


*/
