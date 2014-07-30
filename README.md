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
                orm := mantle.Orm{Driver: "redis"} //mantle.Orm{} <- defaults to redis
                connection := orm.Get()
                fmt.Println(connection.Get("key"))
        } 
