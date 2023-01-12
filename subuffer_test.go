package subuffer

import (
	"testing"
)

type Name struct {
	Name string
}

func (n *Name) Update(data interface{}) {
	n.Name = data.(string)
}

func (n *Name) Clear() {}

func BenchmarkSuBuffer_Update(b *testing.B) {
	doubleBuf := New(&Name{})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doubleBuf.Update("test")
		}
	})
	b.ReportAllocs()
}

func BenchmarkSuBuffer_Buffer(b *testing.B) {
	doubleBuf := New(&Name{})
	doubleBuf.Update("test")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if doubleBuf.Buffer().Name != "test" {
				b.Fail()
			}
		}
	})
	b.ReportAllocs()
}
