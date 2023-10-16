package main

import (
	"fmt"
	"time"
)

func main() {
	t1, err := time.Parse("2006-01-02 15:04:05 -07", "2023-01-01 17:00:00 +07")
	if err != nil {
		panic(err)
	}
	fmt.Println("🟢 t1:", t1, t1.UnixMicro())
}
