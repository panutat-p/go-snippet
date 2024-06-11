# Profiling

https://go.dev/blog/pprof

## Install

Ubuntu
```sh
apt install -y graphviz
```

MacOS
```sh
brew install graphviz
```

## How to read a graph

https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph

## `net/http/pprof`

https://pkg.go.dev/net/http/pprof

https://medium.com/@ravikumarray92/profiling-in-go-with-pprof-e45656df033e

https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof

```go
package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"

    "github.com/labstack/echo/v4"
)

func main() {
    // profiler
    go http.ListenAndServe(":8081", nil)

    // application
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        fmt.Println("GET /")
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
```

Main menu
* Go to http://localhost:8081/debug/pprof

Download
```sh
wget -O heap.pprof http://localhost:8081/debug/pprof/heap
```

## `go tool`

CPU
```sh
go tool pprof http://localhost:8081/debug/pprof/profile?seconds=30
```

Heap
```sh
go tool pprof http://localhost:8081/debug/pprof/heap
```

Go-routines
```sh
go tool pprof http://localhost:8081/debug/pprof/goroutine
```

Threads
```sh
go tool pprof http://localhost:8081/debug/pprof/threadcreate
```

### `top`

* top: Shows the top functions sorted by number of samples in which the function was actively executing
* top10: Shows the top 10 functions sorted by number of samples in which the function was actively executing
* top10 -cum Shows the top 10 functions sorted by number of samples in which the function appeared

> The first column represents the number of samples in which the function was actively executing,
> i.e., it was the function currently running at the time the sample was taken.
> This gives you an idea of how much CPU time the function itself is consuming.

> The fourth column, on the other hand, represents the number of samples in which the function was present on the call stack.
> This includes the times when the function was actively executing (as represented in the first column)
> and the times when it was waiting for a function it called to return.
> This gives you an idea of the total time spent in the function and all the functions it called.

* list `main`: Print the annotated source code for the function specified
* tree: Print a tree of callers and callees
* peek `main`: Print a table of callers and callees of the function specified
* web: Generates a graph & opens it in a web browser
* png: Generate a graph as an image file
* help: Shows a list of available commands

Open GUI
```sh
go tool pprof --http=:8082 http://localhost:8081/debug/pprof/heap
```

* Go to http://localhost:8082/ui

## `runtime/pprof`

https://pkg.go.dev/runtime/pprof

```go
func main() {
	pkg.Create("main.pprof")
	defer pkg.Close()
	s := strings.Repeat("a", 200000)
	_ = Reverse(s)
	pkg.Write()
}

func Reverse(s string) string {
    if len(s) == 0 {
        return s
    }
    return Reverse(s[1:]) + string(s[0])
}
```

```go
import (
    "os"
    "runtime/pprof"
)

var hp *os.File

func Create(name string) {
    file, err := os.Create(name)
    if err != nil {
        panic(err)
    }
    hp = file
}

func Close() {
    hp.Close()
}

func Write() {
    pprof.WriteHeapProfile(hp)
}
```

Open command prompt
```sh
go tool pprof main.pprof
```
