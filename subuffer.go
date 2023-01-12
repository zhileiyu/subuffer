package subuffer

import (
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Buffer interface {
	Update(interface{})
}

type SuBuffer[T Buffer] struct {
	front T
	back  T
	once  sync.Once
	mux   sync.Mutex
}

func New[T Buffer](b T) *SuBuffer[T] {
	if reflect.ValueOf(b).Kind() != reflect.Ptr {
		//must use prt type to create subuffer
		return nil
	}
	return &SuBuffer[T]{
		front: createBuffer(b),
		back:  createBuffer(b),
	}
}

func (d *SuBuffer[T]) Buffer() T {
	return d.front
}

func (d *SuBuffer[T]) Update(data interface{}) {
	d.once.Do(func() {
		d.front.Update(data)
	})
	d.mux.Lock()
	defer d.mux.Unlock()
	d.back.Update(data)
	d.swapFrontAndBack()
}

func (d *SuBuffer[T]) swapFrontAndBack() {
	bPtr := reflect.ValueOf(d.back).UnsafePointer()
	bPtr0 := (*unsafe.Pointer)(unsafe.Pointer(&d.back))
	fPtr := reflect.ValueOf(d.front).UnsafePointer()
	fPtr0 := (*unsafe.Pointer)(unsafe.Pointer(&d.front))
	atomic.StorePointer(fPtr0, bPtr)
	atomic.StorePointer(bPtr0, fPtr)
}

func createBuffer[T Buffer](b T) T {
	bt := reflect.ValueOf(b).Elem().Type()
	return reflect.New(bt).Interface().(T)
}
