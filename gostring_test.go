package gostring

import (
	"testing"
	"unsafe"
)

func TestPadding(t *testing.T) {
	var data [4]uint8 = [4]uint8{0x30, 0x31, 0x32, 0x00}
	s := GoString(unsafe.Pointer(&data))
	if s != "012" {
		t.Errorf("Failed to convert %s instead of %s", s, "012")
	}
}
