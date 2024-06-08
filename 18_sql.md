# database/sql

https://pkg.go.dev/database/sql

https://github.com/go-sql-driver/mysql

## MySQL

```go
import (
    "database/sql"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:1234@/poc?parseTime=true")
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

    _, err = db.Exec("TRUNCATE TABLE fruit")
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("DELETE FROM fruit")
    if err != nil {
        panic(err)
    }
}
```
