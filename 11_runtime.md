# Runtime

## flag

```go
import "flag"
```

```go
var name string

flag.StringVar(&name, "name", "", "a string")
flag.Parse()

if name == "" {
    fmt.Println("🔴 'name' flag is required")
    os.Exit(2)
}

fmt.Println("name:", name)
```

## Graceful shutdown

```go
import (
    "os"
    "os/signal"
    "syscall"
)
```

Graceful shutdown for a server

```go
func main() {
    var stop = make(chan os.Signal, 1)
    signal.Notify(
        stop,
        os.Interrupt,
        syscall.SIGINT,
        syscall.SIGTERM,
    )
    go StartServer()
    <-stop
    fmt.Println("🟡 Gracefully shutting down")
}

func StartServer() {
    time.Sleep(1<<63 - 1)
}
```

Graceful shutdown for a script

```go
func main() {
    var (
        stop = make(chan os.Signal, 1)
        done = make(chan bool, 1)
    )
    signal.Notify(
        stop,
        os.Interrupt,
        syscall.SIGINT,
        syscall.SIGTERM,
    )
    go Run(done)
    select {
    case <-stop:
        fmt.Println("🟡 Gracefully shutting down")
    case isSucess := <-done:
        if isSucess {
            fmt.Println("🟢 Run finished")
        } else {
            fmt.Println("🔴 Run failed")
        }
    }
}

func Run(done chan bool) {
    defer Recover(done)
    time.Sleep(time.Second * 1)
    n := rand.Intn(3)
    fmt.Println("🔵 n:", n)
    if n == 0 {
        done <- true
    } else if n == 1 {
        done <- false
    } else {
        panic("n is not 1 or 2")
    }
}

func Recover(done chan bool) {
    fmt.Println("🔵 defer Recover")
    r := recover()
    if r != nil {
        fmt.Println("🟡 Panic recovered, err:", r)
    }
    done <- false
}
```
