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
