package counter

type Counter struct {
    value atomic.Uint32
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() uint32 {
    return c.value.Add(1)
}
