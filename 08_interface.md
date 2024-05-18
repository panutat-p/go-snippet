# Interface

https://go.dev/tour/methods/9

## PGO

https://tip.golang.org/doc/pgo

Go 1.22, interface method calls are better optimized.

## Stringer

```go
type Fruit struct {
    Name string
    Price int
}

func NewFruit(name string, price int) Fruit {
    return Fruit{
        Name: name,
        Price: price,
    }
}

func (f Fruit) String() string{
    return "Fruit.String()"
}

func (f Fruit) GoString() string{
    return "Fruit.GoString()"
}
```

## Value receivers and Pointer receivers

* `String()` and `GoString()` will be work with `Fruit` not `*Fruit`
* A method set of type `T` consists all methods with a receiver type `T`
* A method set of a pointer `*T` consists all methods with a receiver type `*T`

> Go's design choice to distinguish between value receivers and pointer receivers
> is rooted in its philosophy of explicitness and control over memory and performance characteristics.

> Control over Mutability: Methods with pointer receivers can modify the receiver.
> If a method needs to modify the receiver or the receiver is a large struct, a pointer receiver is more efficient.
> Value receivers get a copy of the value, and cannot modify the original value.

> Control over Memory and Performance: With value receivers,
> Go creates a copy of the value for each method call, which can be costly if the struct is large.
> With pointer receivers, the method operates on the original value directly, avoiding the cost of copying.

> Consistency: If some of the methods of the struct need to be pointer receivers,
> the Go convention is to make all methods for that struct have pointer receivers, for consistency.

> While this design choice might seem inconvenient at times,
> it provides explicit control over how your program uses memory,
> which can be crucial for performance in systems programming.
> It also helps in understanding whether a method call can have side effects by modifying the receiver.
