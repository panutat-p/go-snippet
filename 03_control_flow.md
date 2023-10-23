# Control flow

https://go.dev/ref/spec

## Function
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
