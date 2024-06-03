# ID Generator

## UUID

https://github.com/google/uuid

```go
import (
    "fmt"

    "github.com/google/uuid"
)

func main() {
    id := uuid.New()
    fmt.Println(id.String())
}
```

## KSUID

https://github.com/segmentio/ksuid

```sh
go install github.com/segmentio/ksuid/cmd/ksuid@latest
```

```sh
ksuid -n 5
```

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
