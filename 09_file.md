# File

https://pkg.go.dev/os

https://pkg.go.dev/io

```go
import (
  "os"
  "path/filepath"
)
```

```go
// CreateFile
// fullPath: "app/v1/file.txt"
func CreateFile(fullPath string) error {
  dir := filepath.Dir(fullPath)
  err := os.MkdirAll(dir, 0755)
  if err != nil {
    return err
  }
  file, err := os.Create(fullPath)
  if err != nil {
    return err
  }
  defer file.Close()
  return nil
}
```
