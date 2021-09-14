// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_layout_child_get_type()), F: marshalLayoutChilder},
	})
}

// LayoutChild: GtkLayoutChild is the base class for objects that are meant to
// hold layout properties.
//
// If a GtkLayoutManager has per-child properties, like their packing type, or
// the horizontal and vertical span, or the icon name, then the layout manager
// should use a GtkLayoutChild implementation to store those properties.
//
// A GtkLayoutChild instance is only ever valid while a widget is part of a
// layout.
type LayoutChild struct {
	*externglib.Object
}

// LayoutChilder describes LayoutChild's abstract methods.
type LayoutChilder interface {
	externglib.Objector

	// ChildWidget retrieves the GtkWidget associated to the given layout_child.
	ChildWidget() Widgetter
	// LayoutManager retrieves the GtkLayoutManager instance that created the
	// given layout_child.
	LayoutManager() LayoutManagerer
}

var _ LayoutChilder = (*LayoutChild)(nil)

func wrapLayoutChild(obj *externglib.Object) *LayoutChild {
	return &LayoutChild{
		Object: obj,
	}
}

func marshalLayoutChilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLayoutChild(obj), nil
}

// ChildWidget retrieves the GtkWidget associated to the given layout_child.
func (layoutChild *LayoutChild) ChildWidget() Widgetter {
	var _arg0 *C.GtkLayoutChild // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(layoutChild.Native()))

	_cret = C.gtk_layout_child_get_child_widget(_arg0)
	runtime.KeepAlive(layoutChild)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// LayoutManager retrieves the GtkLayoutManager instance that created the given
// layout_child.
func (layoutChild *LayoutChild) LayoutManager() LayoutManagerer {
	var _arg0 *C.GtkLayoutChild   // out
	var _cret *C.GtkLayoutManager // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(layoutChild.Native()))

	_cret = C.gtk_layout_child_get_layout_manager(_arg0)
	runtime.KeepAlive(layoutChild)

	var _layoutManager LayoutManagerer // out

	_layoutManager = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(LayoutManagerer)

	return _layoutManager
}
