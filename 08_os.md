# os

## Environment variables

```shell
APP_NAME=quick-go go run main.go
```

```shell
export APP_NAME=quick-go
go run main.go
```

```go
name := os.Getenv("APP_NAME")
fmt.Println(name)
```

## Directory

```go
import (
  "os"
  "path/filepath"
)
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
