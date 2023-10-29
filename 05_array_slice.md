# Array & Slice

## Pop

```go
func Pop(sl []int, idx int) []int {
	var (
		ret = make([]int, len(sl)-1, cap(sl)-1)
	)
	copy(ret, sl[:idx])         // copy first half
	copy(ret[idx:], sl[idx+1:]) // copy after the index to the end
	return ret
}
```

## Pakcage sort

```go
nums := []int{9, 3, 2, 8, 1, 5, 7, 2}
sort.Ints(nums)
fmt.Println(nums) // [1 2 2 3 5 7 8 9]
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
