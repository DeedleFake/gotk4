// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_simple_get_type()), F: marshalIMContextSimple},
	})
}

// IMContextSimple: `GtkIMContextSimple` is an input method supporting
// table-based input methods.
//
// `GtkIMContextSimple` has a built-in table of compose sequences that is
// derived from the X11 Compose files.
//
// `GtkIMContextSimple` reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-4.0/Compose, ~/.XCompose,
// /usr/share/X11/locale/$locale/Compose (for locales that have a nontrivial
// Compose file). The syntax of these files is described in the Compose(5)
// manual page.
//
//
// Unicode characters
//
// `GtkIMContextSimple` also supports numeric entry of Unicode characters by
// typing <kbd>Ctrl</kbd>-<kbd>Shift</kbd>-<kbd>u</kbd>, followed by a
// hexadecimal Unicode codepoint.
//
// For example,
//
//    Ctrl-Shift-u 1 2 3 Enter
//
// yields U+0123 LATIN SMALL LETTER G WITH CEDILLA, i.e. ģ.
type IMContextSimple interface {
	IMContext

	// AddComposeFile adds an additional table from the X11 compose file.
	AddComposeFile(composeFile string)
}

// imContextSimple implements the IMContextSimple interface.
type imContextSimple struct {
	IMContext
}

var _ IMContextSimple = (*imContextSimple)(nil)

// WrapIMContextSimple wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMContextSimple(obj *externglib.Object) IMContextSimple {
	return IMContextSimple{
		IMContext: WrapIMContext(obj),
	}
}

func marshalIMContextSimple(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMContextSimple(obj), nil
}

// NewIMContextSimple constructs a class IMContextSimple.
func NewIMContextSimple() IMContextSimple {
	var _cret C.GtkIMContextSimple

	cret = C.gtk_im_context_simple_new()

	var _imContextSimple IMContextSimple

	_imContextSimple = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(IMContextSimple)

	return _imContextSimple
}

// AddComposeFile adds an additional table from the X11 compose file.
func (c imContextSimple) AddComposeFile(composeFile string) {
	var _arg0 *C.GtkIMContextSimple
	var _arg1 *C.char

	_arg0 = (*C.GtkIMContextSimple)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(composeFile))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_context_simple_add_compose_file(_arg0, _arg1)
}
