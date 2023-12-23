# os

```go
import "os"
```

## Exit

```go
os.exit(0)
```

## Environment variables

Run Go program with inline ENV
```shell
APP_NAME=quick-go go run main.go
```

Run Go program with Linux ENV
```shell
export APP_NAME=quick-go
go run main.go
```

```go
name := os.Getenv("APP_NAME")
fmt.Println(name)
```

## Signals

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

Graceful shuwdown for Echo web server

```go
e := echo.New()
e.GET("/", func(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, Echo!")
})

go func() {
  if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
    e.Logger.Fatal(err)
  }
}()

var stop = make(chan os.Signal, 1)
signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
<-stop
fmt.Println("🟡 Gracefully shutting down")
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := e.Shutdown(ctx); err != nil {
  fmt.Println("🔴 Failed to Shutdown Echo")
  fmt.Println(err)
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
    case <-done:
      fmt.Println("🟢 Done")
  }
}

func Run(done chan bool) {
  time.Sleep(time.Second * 5)
  done <- true
}
```

## Directory

```go
import "path/filepath"
```

Get current directory
```go
dir, err := os.Getwd()
if err != nil {
  panic(err)
}
fmt.Println("Current directory:", dir)
dirName := filepath.Base(dir)
fmt.Println("Directory name:", dirName)
parentDir := filepath.Dir(dir)
fmt.Println("Parent directory:", parentDir)
```

Check directory exist
```go
_, err := os.Stat("/root/go")
if os.IsNotExist(err) {
  fmt.Println("Directory does not exist")
} else {
  fmt.Println("Directory exists")
}
```

List all files in the directory including dot
```go
files, err := os.ReadDir("/root")
if err != nil {
  panic(err)
}
for _, file := range files {
  fmt.Println(file.Name())
}
```

Make a directory
* `0` octal (base-8) notation
* `7` owner
* `5` group
* `5` other

Creates a directory
```go
err := os.Mkdir("app", 0755)
if err != nil {
  panic(err)
}
```

Creates a directory and all necessary parent directories if they do not exist
```go
err := os.MkdirAll("app/v1", 0755)
if err != nil {
  panic(err)
}
```
