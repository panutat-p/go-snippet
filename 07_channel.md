# Channel

## FIFO queue

```go
func main() {
    var (
        ch = make(chan int, 10)
        wg sync.WaitGroup
    )
    
    wg.Add(2)

    go func ()  {
        for i := 1; i < 11; i++ {
            ch <- i
            time.Sleep(100 * time.Millisecond)
        }
        fmt.Println("positive producer complete")
        wg.Done()
    }()

    go func ()  {
        for i := -1; i > -11; i-- {
            ch <- i
            time.Sleep(100 * time.Millisecond)
        }
        fmt.Println("negative producer complete")
        wg.Done()
    }()

    go func() {
        wg.Wait()
        fmt.Println("✅ done")
        close(ch)
    }()

    // Read from the channel until it's closed
    for i := range ch {
        fmt.Println(i)
    }
}
```
