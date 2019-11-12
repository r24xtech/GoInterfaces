package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Stack struct {
  slice []int
}

type allStacks interface {
  Push() *Stack
}

func (s *Stack) Push() *Stack{
  x := myutils.Get_int("Enter value: ")
  s.slice = append(s.slice, x)
  return s
}

func addStacks(as ...allStacks) []*Stack {
  var all []*Stack
  for _, v := range as {
    all = append(all, v.Push())
  }
  return all
}

func main() {
  s1 := new(Stack)
  s2 := new(Stack)
  s3 := new(Stack)
  var ds []*Stack

  T := myutils.Get_int("Choose size of stacks: ")
  for t := 0; t < T; t++ {
    ds = addStacks(s1, s2, s3)
    fmt.Println()
  }

  for i := 0; i < len(ds); i++ {
    fmt.Println(ds[i].slice)
  }
}
