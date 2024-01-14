# Array & Slice

```go
var sl []int
fmt.Println(sl == nil) // true
```

```go
var sl = []int{}
fmt.Println(sl == nil) // false
```

```go
var sl = make([]int, 0, 0)
fmt.Println(sl == nil) // false
```

```go
var sl = new([]int)
fmt.Printf("%T\n", sl) // *[]int
fmt.Println(sl == nil) // false
```

## Pop

```go
func Pop(sl []int, idx int) []int {
    var ret = make([]int, 0, cap(sl)-1)
    ret = append(ret, sl[:idx]...)
    ret = append(ret, sl[idx+1:]...)
    return ret
}
```

```go
func Pop(sl []int, idx int) []int {
    var ret = make([]int, len(sl)-1, len(sl)-1)
    copy(ret, sl[:idx])
    copy(ret[idx:], sl[idx+1:])
    return ret
}
```

## Package sort

https://pkg.go.dev/sort

```go
nums := []int{9, 3, 2, 8, 1, 5, 7, 2}
sort.Ints(nums)   // ascending order
fmt.Println(nums) // [1 2 2 3 5 7 8 9]
```

### Less function `func(i, j int) bool`
* returns `true`: elements at index `i` should come before elements at index `j`
* returns `false`: elements at index `i` should come after or be equal to elements at index `j`

```go
nums := []uint32{9, 3, 2, 8, 1, 5, 7, 2}
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j] // ascending order
})
fmt.Println(nums) // [1 2 2 3 5 7 8 9]
```

```go
rsl := []rune("hello world!")
sort.Slice(rsl, func(i, j int) bool {
    return rsl[i] < rsl[j] // ascending order
})
fmt.Println(rsl) // [32 33 100 101 104 108 108 108 111 111 114 119]
```

```go
words := []string{"banana", "apple", "orange", "grape", "cherry"}
sort.Slice(words, func(i, j int) bool {
    return words[i] > words[j] // descending order
})
fmt.Println(words) // [orange grape cherry banana apple]
```

```go
// Sort positive numbers in ascending order
// Swap negative numbers to the back in original order
nums := []int{4, -3, 2, -1, -7}
sort.Slice(nums, func(i, j int) bool {
    if nums[i] < 0 {
        return false // preserve
    }
    if nums[j] < 0 {
        return true // swap
    }
        return nums[i] < nums[j]
})
fmt.Println(nums) // [2 4 -3 -1 -7]
```

## Package slices

https://pkg.go.dev/slices

https://costamagna.medium.com/exploring-the-power-of-go-1-21-slices-package-6e017b2faec9

```go
sl := []int{3, 2, 4, 7, 3, 1, 2, 4 ,6}
slices.Sort(sl)
fmt.Println(sl) // [1 2 2 3 3 4 4 6 7]
```

```go
sl := []int{1, 2, 3, 3, 1, 2, 8, 1}
fmt.Println(slices.Contains(sl, 3)) // true
```

```go
sl := []int{3, 2, 4, 7,3, 1, 2,4,6}
i, found := slices.BinarySearch(sl, 7)
fmt.Println(i) // 3
fmt.Println(found) // true
```

```go
letters := []string{"a", "b", "c", "d", "e"}
letters = slices.Delete(letters, 1, 2)
fmt.Println(letters) // [a c d e]
```
