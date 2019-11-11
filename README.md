# GoInterfaces
Golang interfaces


1. **Interface** = set of methods required to implement such interface. It’s defined using keyword `interface`.

2. **Interface type** = variable of interface type which can hold any value implementing particular interface.
```go
type Node struct {
  prev *Node
  next *Node
  key interface{}
}
// the key field can be of any type (string, int, float,...)
```
