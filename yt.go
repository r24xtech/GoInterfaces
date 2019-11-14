package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  previous *Node
}

type Queue struct {
  size int
  current *Node
  last *Node
  keep *Node
}

type Moves interface {
  Push()
  Pop()
  Print()
}

func (q *Queue) Push() {
  x := myutils.Get_int("Enter value: ")
  qx := &Node{value: x, previous: nil}
  if q.size == 0 {
    q.current = qx
    q.last = q.current
  } else {
    q.last.previous = qx
    q.last = qx
  }
  q.size++
}

func (q *Queue) Pop() {
  if q.size == 0 {
    fmt.Println("Can't pop...the queue is empty!")
  } else {
    q.current = q.current.previous
    q.size--
  }
}

func (q *Queue) Print() {
  if q.size == 0 {
    fmt.Println("[...] the queue is empty!")
  } else {
    q.keep = q.current
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(q.current.value, " ")
      if q.current.previous == nil {
        break
      }
      q.current = q.current.previous
    }
    fmt.Print(" |back")
    fmt.Println("\n")
    q.current = q.keep
  }
}

func PushAll(m ...Moves){
  for _, v := range m {
    v.Push()
  }
}

func PopAll(m ...Moves){
  for _, v := range m {
    v.Pop()
  }
}

func PrintAll(m ...Moves){
  for _, v := range m {
    v.Print()
  }
}

func main() {
  q := new(Queue)

  fmt.Println("|| Welcome to the Queue ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: PushAll(q)
    case 2: PopAll(q)
    case 3: PrintAll(q)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
