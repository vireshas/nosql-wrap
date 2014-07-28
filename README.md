mantle
======

Wrapper for nosql dbs.

## Usage:
        package main

        import "fmt"
        import "github.com/vireshas/mantle"

        func main(){
                orm := mantle.Orm{}
                connection := orm.Get()
                connection.Connect()
                fmt.Println(connection.Get("key"))
        }
