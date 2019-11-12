package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)

type Node struct {
  value int
  nxprev *Node // next for stack, previous for queue
}

type Stack struct {
  size uint64
  top *Node
  keep *Node
}

type Queue struct {
  size uint64
  current *Node
  last *Node
  keep *Node
}

type dataStructure interface {
  Push()
  Pop()
  Print()
}

func (s *Stack) Push() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if s.size != 0 {
    n.nxprev = s.top
  }
  s.top = n
  s.size++
}

func (q *Queue) Push() {
  x := myutils.Get_int("Enter value: ")
  if q.size == 0 {
    q.current = &Node{value: x, nxprev: nil}
    q.last = q.current
    q.size++
  } else {
    qx := &Node{value: x, nxprev: nil}
    q.last.nxprev = qx
    q.last = qx
    q.size++
  }
}

func (s *Stack) Pop() {
  if s.size == 0 {
    fmt.Println("Can't pop...the stack is empty!")
  } else {
    s.top = s.top.nxprev
    s.size--
  }
}

func (q *Queue) Pop() {
  if q.size > 0 {
    q.current = q.current.nxprev
    q.size--
  } else {
    fmt.Println("Can't pop...the queue is empty!")
  }
}

func (s *Stack) Print() {
  if s.size == 0 {
    fmt.Println("[...] The stack is empty!")
  } else {
    fmt.Println()
    fmt.Print("[stack] top| ")
    s.keep = s.top
    for {
      fmt.Print(s.top.value, " ")
      if s.top.nxprev == nil {
        break
      }
      s.top = s.top.nxprev
    }
    s.top = s.keep
    fmt.Print(" |bottom")
    fmt.Println("\n")
  }
}

func (q *Queue) Print() {
  if q.size == 0 {
    fmt.Println("[...] The queue is empty!")
  } else {
    q.keep = q.current
    fmt.Println()
    fmt.Print("[queue] front| ")
    for {
      fmt.Print(q.current.value, " ")
      if q.current.nxprev == nil {
        break
      }
      q.current = q.current.nxprev
    }
    fmt.Print(" |back")
    q.current = q.keep
    fmt.Println("\n")
  }
}

func PushAll(ds ...dataStructure) {
  for _, v := range ds {
    v.Push()
  }
}

func PopAll(ds ...dataStructure) {
  for _, v := range ds {
    v.Pop()
  }
}

func PrintAll(ds ...dataStructure) {
  for _, v := range ds {
    v.Print()
  }
}

func main() {
  s1 := new(Stack)
  q1 := new(Queue)


  fmt.Println("|| Welcome ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: PushAll(s1, q1)
    case 2: PopAll(s1, q1)
    case 3: PrintAll(s1, q1)
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
