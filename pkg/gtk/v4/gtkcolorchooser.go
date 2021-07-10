// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_get_type()), F: marshalColorChooser},
	})
}

// ColorChooserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ColorChooserOverrider interface {
	ColorActivated(color *gdk.RGBA)
	// RGBA gets the currently-selected color.
	RGBA() gdk.RGBA
	// SetRGBA sets the color.
	SetRGBA(color *gdk.RGBA)
}

// ColorChooser: `GtkColorChooser` is an interface that is implemented by
// widgets for choosing colors.
//
// Depending on the situation, colors may be allowed to have alpha
// (translucency).
//
// In GTK, the main widgets that implement this interface are
// [class@Gtk.ColorChooserWidget], [class@Gtk.ColorChooserDialog] and
// [class@Gtk.ColorButton].
type ColorChooser interface {
	gextras.Objector

	// RGBA gets the currently-selected color.
	RGBA() gdk.RGBA
	// UseAlpha returns whether the color chooser shows the alpha channel.
	UseAlpha() bool
	// SetRGBA sets the color.
	SetRGBA(color *gdk.RGBA)
	// SetUseAlpha sets whether or not the color chooser should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)
}

// ColorChooserIface implements the ColorChooser interface.
type ColorChooserIface struct {
	*externglib.Object
}

var _ ColorChooser = (*ColorChooserIface)(nil)

func wrapColorChooser(obj *externglib.Object) ColorChooser {
	return &ColorChooserIface{
		Object: obj,
	}
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorChooser(obj), nil
}

// RGBA gets the currently-selected color.
func (chooser *ColorChooserIface) RGBA() gdk.RGBA {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.GdkRGBA          // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))

	C.gtk_color_chooser_get_rgba(_arg0, &_arg1)

	var _color gdk.RGBA // out

	_color = *(*gdk.RGBA)(unsafe.Pointer((&_arg1)))

	return _color
}

// UseAlpha returns whether the color chooser shows the alpha channel.
func (chooser *ColorChooserIface) UseAlpha() bool {
	var _arg0 *C.GtkColorChooser // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_color_chooser_get_use_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRGBA sets the color.
func (chooser *ColorChooserIface) SetRGBA(color *gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 *C.GdkRGBA         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(color))

	C.gtk_color_chooser_set_rgba(_arg0, _arg1)
}

// SetUseAlpha sets whether or not the color chooser should use the alpha
// channel.
func (chooser *ColorChooserIface) SetUseAlpha(useAlpha bool) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_chooser_set_use_alpha(_arg0, _arg1)
}
