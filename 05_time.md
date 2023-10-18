# Time

https://pkg.go.dev/time

```go
import "time"
```
## DateTime
```go
t1, err := time.Parse(time.DateTime, "2023-01-01 17:00:00")
if err != nil {
  panic(err)
}
fmt.Println("🟢 t1:", t1, t1.Unix())
```

## DateTime with timezone string
```go
t1, err := time.Parse("2006-01-02 15:04:05 -07", "2023-01-01 17:00:00 +07")
if err != nil {
  panic(err)
}
fmt.Println("🟢 t1:", t1, t1.Unix())
```

## DateTime with timezone location
```go
loc, err := time.LoadLocation("Asia/Bangkok")
if err != nil {
  panic(err)
}
t1, err := time.ParseInLocation(time.DateTime, "2023-01-01 17:00:00", loc)
if err != nil {
  panic(err)
}
fmt.Println("🟢 t1:", t1, t1.Unix())
t2, err := time.ParseInLocation(time.DateTime, "2023-01-01 17:00:00", time.UTC)
if err != nil {
  panic(err)
}
fmt.Println("🔵 t2:", t2, t2.Unix())
```

## RFC3339
```go
t1, err := time.Parse(time.RFC3339, "2023-01-01T17:00:00+07:00")
if err != nil {
  panic(err)
}
fmt.Println("🟢 t1:", t1, t1.UnixMicro())
```
