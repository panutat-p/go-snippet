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
sql, args = sb.Build()

var fruit Fruit
err = db.QueryRow(sql, args...).Scan(&fruit.Name, &fruit.Color, &fruit.Price)
if err != nil {
    panic(err)
}

fmt.Println("🟢 fruit:", fruit)
```
