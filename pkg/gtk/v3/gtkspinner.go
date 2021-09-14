// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_spinner_get_type()), F: marshalSpinnerer},
	})
}

// Spinner widget displays an icon-size spinning animation. It is often used as
// an alternative to a ProgressBar for displaying indefinite activity, instead
// of actual progress.
//
// To start the animation, use gtk_spinner_start(), to stop it use
// gtk_spinner_stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner struct {
	Widget
}

func wrapSpinner(obj *externglib.Object) *Spinner {
	return &Spinner{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalSpinnerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSpinner(obj), nil
}

// NewSpinner returns a new spinner widget. Not yet started.
func NewSpinner() *Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner *Spinner // out

	_spinner = wrapSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _spinner
}

// Start starts the animation of the spinner.
func (spinner *Spinner) Start() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_start(_arg0)
	runtime.KeepAlive(spinner)
}

// Stop stops the animation of the spinner.
func (spinner *Spinner) Stop() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_stop(_arg0)
	runtime.KeepAlive(spinner)
}
