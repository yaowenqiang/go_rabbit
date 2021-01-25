package main

import (
    "fmt"
    "github.com/yaowenqiang/go_rabbit/distributed/coordinator"
)

func main() {
    ql := coordinator.NewQueueListener()

    go ql.ListenForNewSource()

    var a  string

    fmt.Scanln(&a)
}
