# database/sql

https://github.com/jmoiron/sqlx

## MySQL

```sh
go get github.com/jmoiron/sqlx
```

```go
import (
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

func main() {
    db, err := sqlx.Connect("mysql", "root:1234@/poc?parseTime=True")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    db.SetConnMaxLifetime(3 * time.Minute) // should be less than 5 minutes
    db.SetMaxOpenConns(10) // limit the number of connection
    db.SetMaxIdleConns(10) // should be equal to max_open_conns

    err = db.Ping()
    if err != nil {
        panic(err)
    }
}
```
