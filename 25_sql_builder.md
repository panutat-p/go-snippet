# go-sqlbuilder

https://github.com/huandu/go-sqlbuilder

## Query

```go
type Fruit struct {
    Name  string
    Color string
    Price decimal.Decimal // "github.com/shopspring/decimal"
}

sb = sqlbuilder.NewSelectBuilder()
sb.Select("name", "color", "price").From("fruit").Where(sb.Equal("id", 1))
command, args = sb.Build()

var fruit Fruit
err = db.QueryRow(command, args...).Scan(&fruit.Name, &fruit.Color, &fruit.Price)
if err != nil {
    panic(err)
}

fmt.Println("🟢 fruit:", fruit)
```

## Query with ORM

```go
type Fruit struct {
    Name  string           `db:"name"`
    Color string           `db:"color"`
    Price decimal.Decimal  `db:"price"`  // "github.com/shopspring/decimal"
}

sb := orm.SelectFrom("fruit")
sb.Where(
    sb.Equal("id", 1),
)
command, args = sb.Build()

rows, err = db.Query(command, args...).Scan(&fruit.Name, &fruit.Color, &fruit.Price)
if err != nil {
    panic(err)
}
defer rows.Close()

var fruit Fruit
err := rows.Scan(userStruct.Addr(&fruit)...)
if err != nil {
    panic(err)
}
```
