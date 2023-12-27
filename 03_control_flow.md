# Control flow

https://go.dev/ref/spec

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

```go
func BubbleSort(nums []int) []int {
    for i := 0; i < len(nums); i += 1 {
        for j := i + 1; j < len(nums); j += 1 {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    return nums
}
```
