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
