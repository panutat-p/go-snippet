package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
    var count Counter
    quit := make(chan bool)

    for {
        go func() {
            for i := 0; i < 10; i++ {
                c := count.Increment()
                fmt.Println("c:", c)
                time.Sleep(time.Second)
            }
        }()
        time.Sleep(5*time.Second)
    }
    <-quit
}

type Counter struct {
    value atomic.Uint32
}

func (c Counter) Increment() uint32 {
    return c.value.Add(1)
}
