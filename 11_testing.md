# testing

https://pkg.go.dev/testing

```go
import "testing"
```

## Simple

```go
func TestCalculator(t *testing.T) {
    t.Run("sum 4", func(t *testing.T) {
        got := Sum(2, 2)
        want := 4
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    })
    t.Run("sum 6", func(t *testing.T) {
        got := Sum(4, 2)
        want := 6
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    })
}
```

## Table

https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721

```go
func TestCalculator(t *testing.T) {
    tests := []struct {
        name string
        a    int
        b    int
        want int
    }{
        {"Sum 2", 2, 2, 4},
        {"Sum 3", 3, 3, 6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Sum(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("got %d want %d", got, tt.want)
            }
        })
    }
}
```

## Table parallel

⚠️ https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721

```go
func TestCalculator(t *testing.T) {
    tests := []struct {
        name string
        a    int
        b    int
        want int
    }{
        {"Sum", 2, 2, 4},
        {"Sum", 3, 3, 6},
        // Add more test cases here
    }

    for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := Sum(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("got %d want %d", got, tt.want)
            }
        })
    }
}
```
