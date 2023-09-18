## Load dot ENV file

https://github.com/joho/godotenv
```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
	panic(err)
  }

  port := os.Getenv("PORT")
}
```

## ENV to struct

https://github.com/caarlos0/env
```go
package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v9"
)

type config struct {
	ID           int            `env:"ID"`
	Port         string         `env:"PORT" envDefault:"8080"`
	IsProduction bool           `env:"PRODUCTION"`
	Hosts        []string       `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration  `env:"DURATION"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
```
