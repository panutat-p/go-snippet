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
sb.
    Select("name", "color", "price").
    From("fruit").
    Where(sb.Equal("id", 1))
command, args = sb.Build()

var fruit Fruit
err = db.QueryRow(command, args...).Scan(&fruit.Name, &fruit.Color, &fruit.Price)
if err != nil {
    panic(err)
}

fmt.Println("🟢 fruit:", fruit)
```

```go
sb := sqlbuilder.NewSelectBuilder()
fromSb := sqlbuilder.NewSelectBuilder()
statusSb := sqlbuilder.NewSelectBuilder()
sb.Select("id")
sb.From(sb.BuilderAs(fromSb, "user"))
sb.From(sb.In("status", statusSb))
fromSb.Select("id").From("user").Where(fromSb.GreaterThan("level", 4))
statusSb.Select("status").From("config").Where(statusSb.Equal("state", 1))
command, args := sb.Build()
fmt.Println("👉", command)
fmt.Println("👉", args)
```

## Query with ORM

```go
type Fruit struct {
    Name  string           `db:"name" fieldtag:"pk"`
    Color string           `db:"color"`
    Price decimal.Decimal  `db:"price"`  // "github.com/shopspring/decimal"
}

var orm = sqlbuilder.NewStruct(new(Fruit))
sb := orm.SelectFrom("fruit")
sb.Where(
    sb.Equal("id", 1),
)
command, args := sb.Build()

rows, err := db.Query(command, args...)
if err != nil {
    panic(err)
}
defer rows.Close()

var fruit Fruit
err = rows.Scan(userStruct.Addr(&fruit)...)
if err != nil {
    panic(err)
}
```

```go
type Fruit struct {
    ID   string `db:"id" fieldtag:"pk"`
    Name string `db:"name"`
}
type Factory struct {
    ID      string `db:"id" fieldtag:"pk"`
    Name    string `db:"name"`
    Product string `db:"product"`
}
var ormFactory = sqlbuilder.NewStruct(new(Factory))
sb1 := ormFactory.SelectFrom("factory")
var ormFruit = sqlbuilder.NewStruct(new(Fruit))
sb := ormFruit.SelectFrom("fruit")
sb.JoinWithOption(
    sqlbuilder.LeftJoin,
    sb.BuilderAs(sb1, "factory"),
    "fruit.name = factory.product",
)
sb.Where(
    sb.GreaterEqualThan("fruit.id", "1"),
)
command, args := sb.Build()
fmt.Println("👉", command)
fmt.Println("👉", args)
```

## Insert

```go
ib := NewInsertBuilder()
ib.InsertInto("fruit")
ib.Cols("name", "color", "price")
ib.Values("apple", "red", 12)
ib.Values("banana", "yellow", 8)
command, args := ib.Build()
```

## Insert with ORM

```go
type Fruit struct {
    Name  string           `db:"name" fieldtag:"pk"`
    Color string           `db:"color"`
    Price decimal.Decimal  `db:"price"`  // "github.com/shopspring/decimal"
}

var orm = NewStruct(new(Fruit))
fruit := Fruit{
	Name:  "apple",
    Color: "red",
	Price: decimal.NewFromInt(12),
}

ib := orm.
    WithoutTag("pk").
    InsertInto("fruit", &fruit)
command, args := ib.Build()
```
