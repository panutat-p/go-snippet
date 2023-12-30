# Basic

https://go.dev/ref/spec

## Primitives

* bool
* string
* int  int8  int16  int32  int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte (alias for uint8)
* rune (alias for int32)
* float32 float64
* complex64 complex128

## Operators

```go
func PrintDigit(num int) {
    if num < 0 {
        num *= -1
    }
    for num > 0 {
        digit := num % 10
        fmt.Println(digit)
        num /= 10
    }
}
```

## nil

[Billion dollar mistake?](https://www.reddit.com/r/golang/comments/18sncxt/go_nil_panic_and_the_billion_dollar_mistake)

https://github.com/uber-go/nilaway

## Auto zero values

* Go: init as zero value
* Kotlin: compilation error
* Rust: compilation error
* JavaScript: init as `undefined`

```go
type Profile struct {
    uuid uuid.UUID
    name string
}

s := Profile {
    // Oops, forgot my `uuid`, it's now `0`, so it's not really unique, is it?
    name: "My name",
}
```
