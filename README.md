Mantle
======

Go wrapper for nosql dbs.

####Get the package:
        go get github.com/vireshas/mantle
        
####Code:
        package main

        import "fmt"
        import "github.com/vireshas/mantle"

        func main(){
                orm := mantle.Orm{Driver: "redis"}
                connection := orm.Get()
                fmt.Println(connection.Set("key", "value2"))
                fmt.Println(connection.Get("key"))
                fmt.Println(connection.MGet("a", "b"))
        }
