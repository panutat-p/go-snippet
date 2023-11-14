# Logger

## slog

https://pkg.go.dev/golang.org/x/exp/slog

https://github.com/golang-cz/devslog

```go
var logger *slog.Logger

func main() {
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		AddSource: true,
	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, options))
	logger.Info(
		"Value of count",
		slog.Int("count", 3),
	)
}
```

## Data redaction

```go
func Mask(s string) string {
	if len(s) <= 5 {
		return strings.Repeat("*", len(s))
	}

	return s[0:5] + strings.Repeat("*", len(s) - 5)
}
```
