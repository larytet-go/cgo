package gostring

import (
	"reflect"
	"unsafe"
)

// TODO Probably Ok for short strings, but I can do better
// for an average case
func StrLen(data unsafe.Pointer) int {
	p := uintptr(data)
	c := *(*uint8)(unsafe.Pointer(p))
	for c != 0 {
		p++
		c = *(*uint8)(unsafe.Pointer(p))
	}
	return int(p - uintptr(data))
}

// GoString typecast pointer to go.string
// The function does not allocate memory. Make sure
// that the reference data remains valid through the
// life cycle of the resulting sting
// TODO causes memoty leak? How does it work?
func GoString(data unsafe.Pointer) string {
	sh := &reflect.StringHeader{
		Data: uintptr(data),
		Len:  StrLen(data),
	}
	return *((*string)(unsafe.Pointer(sh)))
}
