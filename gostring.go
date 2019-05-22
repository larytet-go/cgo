package gostring

import (
	"reflect"
	"unsafe"
)

// TODO I can do better
func GoString(data unsafe.Pointer) string {
	p0 := uintptr(data)
	p := p0
	c := *(*uint8)(unsafe.Pointer(p))
	for c != 0 {
		p++
		c = *(*uint8)(unsafe.Pointer(p))
	}
	sh := &reflect.StringHeader{
		Data: uintptr(data),
		Len:  int(p - p0),
	}
	return *((*string)(unsafe.Pointer(sh)))
}
