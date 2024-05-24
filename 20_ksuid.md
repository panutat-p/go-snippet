# KSUID

https://github.com/segmentio/ksuid

```go
package main

import (
    "fmt"
    "github.com/segmentio/ksuid"
)

func main() {
    id := ksuid.New()
    fmt.Println(id.String())
}
```
