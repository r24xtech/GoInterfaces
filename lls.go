package main

import (
  "fmt"
  "github.com/r24xtech/myutils"
)
/*
'next' goes to the back ==>
FRONT| x y z a b c d e f |BACK
<== goes to the front 'previous'
*/

type Node struct {
  value int
  next *Node
  previous *Node
}

type List struct {
  size uint64
  head *Node
  tail *Node
  keep *Node
}

type Moves interface {
  InsertBack()
  InsertFront()
  RemoveBack()
  RemoveFront()
  Print()
}

func (l *List) InsertBack() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.next = n
    n.previous = l.tail
    l.tail = n
  }
  l.size++
}

func (l *List) InsertFront() {
  n := new(Node)
  x := myutils.Get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.head.previous = n
    n.next = l.head
    l.head = n
  }
  l.size++
}

func (l *List) Print() {
  if l.size == 0 {
    fmt.Println("[...] the list is empty!")
  } else {
    l.keep = l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(l.head.value, " ")
      if l.head.next == nil {
        break
      }
      l.head = l.head.next
    }
    fmt.Print(" |back")
    fmt.Println("\n")
    l.head = l.keep
  }
}

func (l *List) RemoveBack() {
  if l.size == 0 {
    fmt.Println("Can't remove...empty list!")
  } else if l.size == 1 {
    l.tail = l.tail.previous
    l.size--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.size--
  }
}

func (l *List) RemoveFront() {
  if l.size == 0 {
    fmt.Println("Can't remove...empty list!")
  } else {
    l.head = l.head.next
    l.size--
  }
}

func InsertFrontAll(m ...Moves) {
  for _, v := range m {
    v.InsertFront()
  }
}

func InsertBackAll(m ...Moves) {
  for _, v := range m {
    v.InsertBack()
  }
}

func RemoveFrontAll(m ...Moves) {
  for _, v := range m {
    v.RemoveFront()
  }
}

func RemoveBackAll(m ...Moves) {
  for _, v := range m {
    v.RemoveBack()
  }
}

func PrintAll(m ...Moves) {
  for _, v := range m {
    v.Print()
  }
}

func main() {
  l1 := new(List)
  l2 := new(List)
  l3 := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print;")
  fmt.Println("[4] RemoveBack; [5] RemoveFront; [6] End;\n\n")

  myloop:for {
    option := myutils.Get_int("Choose option: ")
    switch option {
    case 1: InsertBackAll(l1, l2, l3)
    case 2: InsertFrontAll(l1, l2, l3)
    case 3: PrintAll(l1, l2, l3)
    case 4: RemoveBackAll(l1, l2, l3)
    case 5: RemoveFrontAll(l1, l2, l3)
    case 6: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
