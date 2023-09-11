package counter

import (
	"testing"
)

func TestCounter_increment_one(t *testing.T) {
    var count = NewCounter()
    c = count.Increment()
    if c != 1 {
        t.Error("Invalid counter")
    }
}
