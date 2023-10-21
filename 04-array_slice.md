# Go 1.21’s ‘slices’ upgrades!

https://medium.com/@nlcostamagna/exploring-the-power-of-go-1-21-slices-package-6e017b2faec9

https://medium.com/@emreodabas_20110/quick-guide-go-1-21-features-80e302ec0110

```go
sl := []int{3, 2, 4, 7,3, 1, 2,4,6}
i, found := slices.BinarySearch(sl, 7)
fmt.Println(i) // 3
fmt.Println(found) // true
```

```go
sl := []int{3, 2, 4, 7, 3, 1, 2, 4 ,6}
slices.Sort(sl)
fmt.Println(sl) // [1 2 2 3 3 4 4 6 7]
```

```go
sl := []int{1, 2, 3, 3, 1, 2, 8, 1}
fmt.Println(slices.Contains(sl, 3)) // true
```
