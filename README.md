# DBL (Doubly Linked List)

![](https://github.com/gompact/dbl/workflows/Build/badge.svg)

dbl is a Doubly Linked List implementation in Go

## Usage

```go
package main

import "github.com/gompact/dbl"

func main() {
    // create empty list
    l := dbl.NewList()
    // append values to the list
    l.Append("first node")
    l.Append("second node")
    
    // remove first index
    ok := l.Remove(0)

    // Pop last element
    l.Append("third node")
    thirdNode := l.Pop()

    // remove all elements in the list
    l.RemoveAll()
    
    // get length of the list
    length := l.Length()
}
```
