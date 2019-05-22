package cgo

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestNonzero(t *testing.T) {
	var data [4]uint8 = [4]uint8{0x30, 0x31, 0x32, 0x00}
	s := GoString(unsafe.Pointer(&data))
	if s != "012" {
		t.Errorf("Failed to convert '%s' instead of '%s'", s, "012")
	}
}

func TestZero(t *testing.T) {
	var data [1]uint8 = [1]uint8{0x00}
	s := GoString(unsafe.Pointer(&data))
	if s != "" {
		t.Errorf("Failed to convert empty string, got %s", s)
	}
}

func TestClone(t *testing.T) {
	type Data struct {
		x int
		y int
	}
	i1 := &Data{
		x: 1,
		y: 2,
	}
	d := Data{}
	i2 := CloneBytes(unsafe.Pointer(i1), int(unsafe.Sizeof(d)))
	data := ((*reflect.SliceHeader)(unsafe.Pointer(&i2))).Data
	i3 := (*Data)(unsafe.Pointer(data))
	if !reflect.DeepEqual(i1, i3) {
		t.Errorf("Failed to clone %v instead of %v", *i1, *i3)
	}
}
