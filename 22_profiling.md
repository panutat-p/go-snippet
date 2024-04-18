# Profiling

## `net/http/pprof`

https://medium.com/@ravikumarray92/profiling-in-go-with-pprof-e45656df033e

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

Open command prompt
```sh
go tool pprof http://localhost:8081/debug/pprof/heap
```
* top: Shows the top functions that are taking up the most memory or CPU (depending on the profile you're looking at)
* tree: Print a tree of callers and callees
* peek `main`: Print a table of callers and callees of the function specified
* web: Generates a graph & opens it in a web browser
* png: Generate a graph as an image file
* list `main`: Print the annotated source code for the function specified
* help: Shows a list of available commands

Open GUI
```sh
go tool pprof --http=:8082 http://localhost:8081/debug/pprof/heap
```

* Go to http://localhost:8082/ui

## `runtime/pprof`
