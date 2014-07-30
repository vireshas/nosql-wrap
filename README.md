Mantle
======

Go wrapper for nosql dbs.

####Get the package:
        go get github.com/vireshas/mantle
        
####Code:
        package main

        import (
                "fmt"
                "github.com/vireshas/mantle"
                "time"
        )

        func main(){
                keyValue := map[string]interface{}{"key1":"val1", "key2":"val2", "key3":"val3"}

                orm := mantle.Orm{Driver: "redis"}
                connection := orm.Get()

                fmt.Println(connection.Set("key", "value2"))
                //output: true

                fmt.Println(connection.Get("key"))
                //output: value2

                fmt.Println(connection.MSet(keyValue))
                //output: true

                fmt.Println(connection.MGet("key3", "key2"))
                //output: [val3 val2]

                connection.Expire("key", 2)
                time.Sleep(2 * time.Second)
                fmt.Println(connection.Get("key"))
                //output: "" 
        }
