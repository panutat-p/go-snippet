package counter

import (
	"testing"
)

func TestCounter_increment_one(t *testing.T) {
    var count = NewCounter()
    c := count.Increment()
    if c != 1 {
        t.Error("Invalid counter")
    }
}

func TestCounter_increment_go_routines(t *testing.T) {
	var count = NewCounter()
	var c uint32

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🟢 c:", c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🔵 c:", c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c = count.Increment()
			fmt.Println("🟠 c:", c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(5*time.Second)
}
