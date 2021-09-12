// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gdk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_snapshot_get_type()), F: marshalSnapshotter},
	})
}

// Snapshot: base type for snapshot operations.
//
// The subclass of GdkSnapshot used by GTK is gtk.Snapshot.
type Snapshot struct {
	*externglib.Object
}

// Snapshotter describes Snapshot's abstract methods.
type Snapshotter interface {
	externglib.Objector

	privateSnapshot()
}

var _ Snapshotter = (*Snapshot)(nil)

func wrapSnapshot(obj *externglib.Object) *Snapshot {
	return &Snapshot{
		Object: obj,
	}
}

func marshalSnapshotter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSnapshot(obj), nil
}

func (*Snapshot) privateSnapshot() {}
