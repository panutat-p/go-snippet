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
}
```

```go
_, err = db.Exec("TRUNCATE TABLE fruit")
if err != nil {
    panic(err)
}
```

```go
_, err = db.Exec("DELETE FROM fruit")
if err != nil {
    panic(err)
}
```

```go
rows, err := db.Query("SELECT name, color, price FROM fruit LIMIT 50")
if err != nil {
    panic(err)
}
defer rows.Close()

type Fruit struct {
    Name  string
    Color string
    Price decimal.Decimal // "github.com/shopspring/decimal"
}
var fruit Fruit
err = db.QueryRow("SELECT name, color, price FROM fruit WHERE id = ?", 1).Scan(&fruit.Name, &fruit.Color, &fruit.Price)
if err != nil {
    panic(err)
}

fmt.Println("🟢 fruit:", fruit)
```

```go
db, err := sql.Open("mysql", "root:1234@/poc")
if err != nil {
    panic(err)
}

db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)

rows, err := db.Query("SELECT name, color, price FROM fruit LIMIT 50")
if err != nil {
    panic(err)
}
defer rows.Close()

type Fruit struct {
    Name  string
    Color string
    Price decimal.Decimal // "github.com/shopspring/decimal"
}
var fruits []Fruit
for rows.Next() {
    var f Fruit
    err := rows.Scan(&f.Name, &f.Color, &f.Price)
    if err != nil {
        panic(err)
    }
    fruits = append(fruits, f)
}

err = rows.Err()
if err != nil {
    panic(err)
}

fmt.Println("🟢 fruits:", fruits)
```
