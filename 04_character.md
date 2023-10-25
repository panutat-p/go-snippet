# Character

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
