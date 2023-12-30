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

### nilaway

https://github.com/uber-go/nilaway

```shell
go install go.uber.org/nilaway/cmd/nilaway@latest
```

## Auto zero values

* Go: init as zero value
* Kotlin: compilation error
* Rust: compilation error
* JavaScript: init as `undefined`
* [The zero value design of Go was pretty intriguing](https://www.reddit.com/r/golang/comments/18sncxt/comment/kf9dha8/?utm_source=share&utm_medium=web2x&context=3)

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
