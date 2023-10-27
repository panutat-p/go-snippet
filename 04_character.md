# Character

## Rune

```go
for _, e := range "hello world" {
	fmt.Println(e) // e is rune
}
```

```go
// INT convert '0' to 0, '1' to 1, ...
func INT(r rune) int {
	return int(r - '0')
}
```

## Unicode

```go
r := rune(97)
if unicode.IsLetter(r) {
	fmt.Println("r is a letter:", string(r))
}
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
