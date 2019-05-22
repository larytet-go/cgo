package gostring

import (
  "unsafe"
  "reflect"
)

// TODO I can do better when
func goString(data unsafe.Pointer) string {
	chars := (*uint8)(data)
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