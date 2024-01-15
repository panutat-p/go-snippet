# Character

## Rune

```go
r := rune(43)
fmt.Println(r)
```

```go
sl := []rune("123 ABC abc")
fmt.Println(sl) // [49 50 51 32 65 66 67 32 97 98 99]
```

```go
for _, r := range "hello world" {
    s := string(r)
    fmt.Println(r, s)
}
```

```go
func IsLetterOrDigit(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}
```

```go
func CountDigits(input string) [10]int {
    var digits [10]int
    for _, r := range input {
        if unicode.IsDigit(r) {
            idx := r - '0'
            digits[idx]++
        }
    }
    return digits
}
```

```go
func CountAlphabet(input string) [26]int {
    var alphabets [26]int
    for _, r := range input {
        if unicode.IsLetter(r) {
            lower := unicode.ToLower(r)
            idx := lower - 'a'
            alphabets[idx]++
        }
    }
    return alphabets
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
