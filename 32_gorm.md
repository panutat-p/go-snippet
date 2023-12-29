# Gorm

https://gorm.io

```shell
go get gorm.io/gorm
```

```go
import (
    "database/sql"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)
```

## Model

```go
type Fruit struct {
    ID    uint64 `gorm:"column:id"`
    Name  string `gorm:"column:name"`
    Price int    `gorm:"price"`
    Tags  string `gorm:"tags"`
}

func (f *Fruit) TableName() string {
    return "fruits"
}
```

## Connect

```go
func Connect(host, port, username, password, dbName string) (*gorm.DB, *sql.DB) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?parseTime=True&charset=utf8",
        username,
        password,
        host,
        port,
        dbName,
    )
    g, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        SkipDefaultTransaction: true,
    })
    if err != nil {
        panic(err)
    }
    db, err := g.DB()
    if err != nil {
        panic(err)
    }
    return g, db
}
```
