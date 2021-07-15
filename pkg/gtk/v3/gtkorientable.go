// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_orientable_get_type()), F: marshalOrientabler},
	})
}

// Orientable interface is implemented by all widgets that can be oriented
// horizontally or vertically. Historically, such widgets have been realized as
// subclasses of a common base class (e.g Box/HBox/VBox or Scale/HScale/VScale).
// Orientable is more flexible in that it allows the orientation to be changed
// at runtime, allowing the widgets to “flip”.
//
// Orientable was introduced in GTK+ 2.16.
type Orientable struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Orientable)(nil)

// Orientabler describes Orientable's abstract methods.
type Orientabler interface {
	// Orientation retrieves the orientation of the orientable.
	Orientation() Orientation
	// SetOrientation sets the orientation of the orientable.
	SetOrientation(orientation Orientation)
}

var _ Orientabler = (*Orientable)(nil)

func wrapOrientable(obj *externglib.Object) *Orientable {
	return &Orientable{
		Object: obj,
	}
}

func marshalOrientabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOrientable(obj), nil
}

// Orientation retrieves the orientation of the orientable.
func (orientable *Orientable) Orientation() Orientation {
	var _arg0 *C.GtkOrientable // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(orientable.Native()))

	_cret = C.gtk_orientable_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation of the orientable.
func (orientable *Orientable) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkOrientable // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(orientable.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_orientable_set_orientation(_arg0, _arg1)
}
