package set

import (
	"testing"
	"unsafe"
)

func TestSetType(t *testing.T) {
	if unsafe.Sizeof(Empty{}) != 0 { // internal implementation of golang test
		t.Errorf("the empty struct should be zero bytes")
	}
}
