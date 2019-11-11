package main

import (
  "math"
  "fmt"
)

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
