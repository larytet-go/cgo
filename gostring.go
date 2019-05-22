package gostring

import (
	"reflect"
	"unsafe"
)

func StrLen(data unsafe.Pointer) int {
	p := uintptr(data)
	c := *(*uint8)(unsafe.Pointer(p))
	for c != 0 {
		p++
		c = *(*uint8)(unsafe.Pointer(p))
	}
	return int(p - uintptr(data))
}

// TODO I can do better
func GoString(data unsafe.Pointer) string {
	sh := &reflect.StringHeader{
		Data: uintptr(data),
		Len:  StrLen(data),
	}
	return *((*string)(unsafe.Pointer(sh)))
}
