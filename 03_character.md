# Character

## Rune

```go
r := rune(43)
fmt.Println(r)
```

```go
sl := []rune("hello world")
fmt.Println(sl)
```

```go
for _, e := range "hello world" {
    fmt.Println(e, string(e))
}
```

```go
// INT convert '0' to 0, '1' to 1, ...
func INT(r rune) int {
    return int(r - '0')
}
```

```go
if unicode.IsLetter('A') {
    fmt.Println("🟢", 'A')
}
```

```go
if unicode.IsDigit('3') {
    fmt.Println("🟢", '3')
}
```

```go
input := "49672058310"
var digits [10]int
for _, r := range input {
    position := r - '0'
    digits[position]++
}
fmt.Println(digits) // [2 1 1 1 1 1 1 1 1 1]
```

## String

```go
s := "Hello World"
lower := strings.ToLower(s)
fmt.Println(lower)
```

```go
s := "Hello World"
sl := strings.Split(s, " ")
fmt.Println(sl)
```

```go
func Reverse(s string) string {
    sl := []rune(s)
    var ret []rune
    for i := len(sl) - 1; i > -1; i -= 1 {
        ret = append(ret, sl[i])
    }
    return string(ret)
}
```
