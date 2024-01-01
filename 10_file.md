# File

https://pkg.go.dev/os

https://pkg.go.dev/io

```go
import (
    "os"
    "path/filepath"
)
```

```go
// CreateFile
// fullPath: "app/v1/file.txt"
func CreateFile(fullPath string) error {
    dir := filepath.Dir(fullPath)
    err := os.MkdirAll(dir, 0755)
    if err != nil {
        return err
    }
    file, err := os.Create(fullPath)
    if err != nil {
        return err
    }
    defer file.Close()
        return nil
}
```

## Write slice into files

```go
var (
    size = 30
)

func main() {
    fruits := GenerateFruits(100)

    var chunks []Fruit
    for i, fruit := range fruits {
        chunks = append(chunks, fruit)
        if (i+1)%size == 0 || i+1 == len(fruits) {
            fileName := fmt.Sprintf("fruits_%d.jsonl", i/size+1)
            WriteFruits(fileName, chunks)
            chunks = nil
        }
    }
}

func GenerateFruits(count int) []Fruit {
    var fruits []Fruit
    for i := 0; i < count; i++ {
        f := Fruit{
            ID:    i + 1,
            Name:  RandStringRunes(6),
            Price: rand.Intn(100),
        }
        fruits = append(fruits, f)
    }
    return fruits
}

func RandStringRunes(n int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func WriteFruits(fileName string, fruits []Fruit) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, fruit := range fruits {
        b, err := json.Marshal(fruit)
        if err != nil {
            return err
        }
        file.Write(b)
        file.WriteString("\n")
    }

    return nil
}

type Fruit struct {
    ID    int
    Name  string
    Price int
}
```
