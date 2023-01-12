package subuffer

import (
	"reflect"
	"sync/atomic"
)

type Buffer interface {
	Update(interface{})
	Clear()
}

type SuBuffer[T Buffer] struct {
	first  T
	second T
	now    *int32
}

func New[T Buffer](b T) *SuBuffer[T] {
	if reflect.ValueOf(b).Kind() != reflect.Ptr {
		//must use prt type to create subuffer
		return nil
	}
	var zero int32
	return &SuBuffer[T]{
		first:  createBuffer(b),
		second: createBuffer(b),
		now:    &zero,
	}
}

func (d *SuBuffer[T]) Buffer() T {
	if atomic.LoadInt32(d.now) == 0 {
		return d.first
	} else {
		return d.second
	}
}

func (d *SuBuffer[T]) Update(data interface{}) {
	if atomic.CompareAndSwapInt32(d.now, 0, 1) {
		d.second.Update(data)
	} else {
		d.first.Update(data)
		atomic.SwapInt32(d.now, 0)
	}
}

func createBuffer[T Buffer](b T) T {
	bt := reflect.ValueOf(b).Elem().Type()
	t := reflect.New(bt).Interface().(T)
	t.Clear()
	return t
}
