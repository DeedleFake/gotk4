// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSnapshot = coreglib.Type(C.gdk_snapshot_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSnapshot, F: marshalSnapshot},
	})
}

// SnapshotOverrider contains methods that are overridable.
type SnapshotOverrider interface {
}

// Snapshot: base type for snapshot operations.
//
// The subclass of GdkSnapshot used by GTK is gtk.Snapshot.
type Snapshot struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Snapshot)(nil)
)

// Snapshotter describes types inherited from class Snapshot.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Snapshotter interface {
	coreglib.Objector
	baseSnapshot() *Snapshot
}

var _ Snapshotter = (*Snapshot)(nil)

func classInitSnapshotter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSnapshot(obj *coreglib.Object) *Snapshot {
	return &Snapshot{
		Object: obj,
	}
}

func marshalSnapshot(p uintptr) (interface{}, error) {
	return wrapSnapshot(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Snapshot) baseSnapshot() *Snapshot {
	return v
}

// BaseSnapshot returns the underlying base object.
func BaseSnapshot(obj Snapshotter) *Snapshot {
	return obj.baseSnapshot()
}
