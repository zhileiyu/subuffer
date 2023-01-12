# subuffer
A lock-free concurrent doublebuffer with generic wirten in go. In chinese su means fast!

用go语言实现的泛型双缓冲buffer，并发安全，无锁，速度很快。

使用示例：

Example:

```go

package main

import (
	"fmt"
	"github.com/zhileiyu/subuffer"
)

type Name struct {
	Name string
}

func (n *Name) Update(data interface{}) {
	n.Name = data.(string)
}

func (n *Name) Show() {
	fmt.Println(n.Name)
}

func main() {
	doubleBuf := subuffer.New(&Name{})
	doubleBuf.Update("Alice")
	doubleBuf.Buffer().Show()
	doubleBuf.Update("Bob")
	doubleBuf.Buffer().Show()
}

```
