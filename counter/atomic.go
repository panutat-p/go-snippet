package counter

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// https://pkg.go.dev/sync/atomic
// https://gobyexample.com/atomic-counters

func main() {
    var ops uint64
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("ops:", ops)
}
