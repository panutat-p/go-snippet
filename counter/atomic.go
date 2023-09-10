package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
	var quit chan bool
    var count = NewCounter()
	var c uint32

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🟢 c:", c)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🔵 c:", c)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🟠 c:", c)
			time.Sleep(1 * time.Second)
		}
	}()

	<- quit
}

type Counter struct {
    value atomic.Uint32
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() uint32 {
    return c.value.Add(1)
}
