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

## Layout guidelines

https://go.dev/doc/modules/layout

### Package names

https://go.dev/blog/package-names

* Avoid meaningless: `util`, `common`, `misc`, ...
* Avoid single entry point: `api`, `types`, `interfaces`, ...
* Avoid unnecessary package name collisions

### Import

relative\
https://pkg.go.dev/cmd/go#hdr-Relative_import_paths

remote\
https://pkg.go.dev/cmd/go#hdr-Remote_import_paths

### Server layout

* Server projects typically won’t have packages for export
* It’s recommended to keep the Go packages implementing the server’s logic in the `internal` directory

```
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
```

### pkg directory

https://pkg.go.dev/github.com/geektime007/mgmt/pkg

https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal

