// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_invisible_get_type()), F: marshalInvisible},
	})
}

// Invisible: the Invisible widget is used internally in GTK+, and is probably
// not very useful for application developers.
//
// It is used for reliable pointer grabs and selection handling in the code for
// drag-and-drop.
type Invisible interface {
	gextras.Objector

	// Screen returns the Screen object associated with @invisible
	Screen() *gdk.ScreenClass
	// SetScreen sets the Screen where the Invisible object will be displayed.
	SetScreen(screen gdk.Screen)
}

// InvisibleClass implements the Invisible interface.
type InvisibleClass struct {
	*externglib.Object
	WidgetClass
	BuildableIface
}

var _ Invisible = (*InvisibleClass)(nil)

func wrapInvisible(obj *externglib.Object) Invisible {
	return &InvisibleClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
	}
}

func marshalInvisible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInvisible(obj), nil
}

// NewInvisible creates a new Invisible.
func NewInvisible() *InvisibleClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_invisible_new()

	var _invisible *InvisibleClass // out

	_invisible = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*InvisibleClass)

	return _invisible
}

// NewInvisibleForScreen creates a new Invisible object for a specified screen
func NewInvisibleForScreen(screen gdk.Screen) *InvisibleClass {
	var _arg1 *C.GdkScreen // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gtk_invisible_new_for_screen(_arg1)

	var _invisible *InvisibleClass // out

	_invisible = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*InvisibleClass)

	return _invisible
}

// Screen returns the Screen object associated with @invisible
func (invisible *InvisibleClass) Screen() *gdk.ScreenClass {
	var _arg0 *C.GtkInvisible // out
	var _cret *C.GdkScreen    // in

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(invisible.Native()))

	_cret = C.gtk_invisible_get_screen(_arg0)

	var _screen *gdk.ScreenClass // out

	_screen = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.ScreenClass)

	return _screen
}

// SetScreen sets the Screen where the Invisible object will be displayed.
func (invisible *InvisibleClass) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkInvisible // out
	var _arg1 *C.GdkScreen    // out

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(invisible.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_invisible_set_screen(_arg0, _arg1)
}
