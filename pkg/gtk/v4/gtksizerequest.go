// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <gtk/gtk.h>
import "C"

// DistributeNaturalAllocation distributes extra_space to child sizes by
// bringing smaller children up to natural size first.
//
// The remaining space will be added to the minimum_size member of the
// GtkRequestedSize struct. If all sizes reach their natural size then the
// remaining space is returned.
func DistributeNaturalAllocation(extraSpace int, sizes []RequestedSize) int {
	var _arg1 C.int               // out
	var _arg3 *C.GtkRequestedSize // out
	var _arg2 C.guint
	var _cret C.int // in

	_arg1 = C.int(extraSpace)
	_arg2 = (C.guint)(len(sizes))
	_arg3 = (*C.GtkRequestedSize)(C.malloc(C.ulong(len(sizes)) * C.ulong(C.sizeof_GtkRequestedSize)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GtkRequestedSize)(_arg3), len(sizes))
		for i := range sizes {
			out[i] = *(*C.GtkRequestedSize)(gextras.StructNative(unsafe.Pointer((&sizes[i]))))
		}
	}

	_cret = C.gtk_distribute_natural_allocation(_arg1, _arg2, _arg3)
	runtime.KeepAlive(extraSpace)
	runtime.KeepAlive(sizes)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
//
// An instance of this type is always passed by reference.
type RequestedSize struct {
	*requestedSize
}

// requestedSize is the struct that's finalized.
type requestedSize struct {
	native *C.GtkRequestedSize
}

// Data: client pointer
func (r *RequestedSize) Data() cgo.Handle {
	var v cgo.Handle // out
	v = (cgo.Handle)(unsafe.Pointer(r.native.data))
	return v
}

// MinimumSize: minimum size needed for allocation in a given orientation
func (r *RequestedSize) MinimumSize() int {
	var v int // out
	v = int(r.native.minimum_size)
	return v
}

// NaturalSize: natural size for allocation in a given orientation
func (r *RequestedSize) NaturalSize() int {
	var v int // out
	v = int(r.native.natural_size)
	return v
}
