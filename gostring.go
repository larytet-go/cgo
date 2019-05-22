package gostring

import (
	"reflect"
	"unsafe"
)

// TODO I can do better
func GoString(data unsafe.Pointer) string {
	var count uintptr
	for {
		pC := uintptr(data) + count
		c := *(*uint8)(unsafe.Pointer(pC))
		if c == 0 {
			break
		}
		count++
	}
	sh := &reflect.StringHeader{
		Data: uintptr(data),
		Len:  int(count),
	}
	return *((*string)(unsafe.Pointer(sh)))
}
