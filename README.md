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

```go
// bad example
func totalArea(circles []Circle, rectangles []Rectangle) float64 {
  var total float64
  for _, c := range circles {
    total += c.area()
  }
  for _, r := range rectangles {
    total += r.area()
  }
  return total
}
```
This works, but it has a  major  issue — whenever we define a new shape, we have to change our function to handle it (a third parameter for Polygons, a fourth for Squares, etc.).

This is the problem interfaces are designed to solve.

Because both of our shapes have an area method, they both implement the Shape interface and we can change our function to this:

```go
func totalArea(shapes ...Shape) []float64 {
  var areas []float64
  for _, s := range shapes {
    areas = append(areas, s.area())
  }
  return areas
}

func main() {
  c := Circle{0, 0, 5}
  r := Rectangle{0, 0, 10, 10}

  a := totalArea(&c, &r)
  for i := 0; i < len(a); i++ {
    fmt.Println("Output: ", a[i])
  }
}
// Output: 78.53981633974483
// Output: 100
```
