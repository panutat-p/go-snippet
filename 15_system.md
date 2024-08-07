# System

## os

```go
import "os"
```

```go
os.exit(0)
```

Run Go program with inline ENV
```sh
APP_NAME=quick-go go run main.go
```

Run Go program with Linux ENV
```sh
export APP_NAME=quick-go
go run main.go
```

```go
name := os.Getenv("APP_NAME")
fmt.Println(name)
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
