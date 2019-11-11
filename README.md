# GoInterfaces
Golang interfaces


1. **Interface** = set of methods required to implement such interface. It’s defined using keyword `interface`.
```go
type Shape interface {
  area() float64
}
```

2. **Interface type** = variable of interface type which can hold any value implementing particular interface.
```go
type Node struct {
  prev *Node
  next *Node
  key interface{}
}
// the key field can be of any type (string, int, float,...)
```

* Like a struct, an interface is created using the `type` keyword, followed by a name and the keyword `interface`. But instead of defining fields, we define a **method set**. A method set is a **list of methods** that a type must have in order to implement the inter‐face.

* If the struct has a pointer receiver on some of its methods, it is better to use it for the rest of the methods because it enables better consistency and predictability for the struct behaviors.
* `(s Stack)` vs. `(s *Stack)`: If you want to modify the data of a receiver from the method, the receiver must be a pointer.

```go
type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Shape interface {
  area() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
  return math.Pi * c.r *c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
```

Suppose we want to write a function that calculates the area of several shapes. 
