// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProvider},
	})
}

// StyleProvider: gtkStyleProvider is an interface used to provide style
// information to a StyleContext. See gtk_style_context_add_provider() and
// gtk_style_context_add_provider_for_screen().
type StyleProvider interface {
	gextras.Objector

	// IconFactory looks up a widget style property as defined by @provider for
	// the widget represented by @path.
	IconFactory(path *WidgetPath) IconFactory
	// Style looks up a widget style property as defined by @provider for the
	// widget represented by @path.
	Style(path *WidgetPath) StyleProperties
}

// styleProvider implements the StyleProvider interface.
type styleProvider struct {
	gextras.Objector
}

var _ StyleProvider = (*styleProvider)(nil)

// WrapStyleProvider wraps a GObject to a type that implements
// interface StyleProvider. It is primarily used internally.
func WrapStyleProvider(obj *externglib.Object) StyleProvider {
	return styleProvider{
		Objector: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProvider(obj), nil
}

func (p styleProvider) IconFactory(path *WidgetPath) IconFactory {
	var _arg0 *C.GtkStyleProvider // out
	var _arg1 *C.GtkWidgetPath    // out
	var _cret *C.GtkIconFactory   // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_style_provider_get_icon_factory(_arg0, _arg1)

	var _iconFactory IconFactory // out

	_iconFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IconFactory)

	return _iconFactory
}

func (p styleProvider) Style(path *WidgetPath) StyleProperties {
	var _arg0 *C.GtkStyleProvider   // out
	var _arg1 *C.GtkWidgetPath      // out
	var _cret *C.GtkStyleProperties // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_style_provider_get_style(_arg0, _arg1)

	var _styleProperties StyleProperties // out

	_styleProperties = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(StyleProperties)

	return _styleProperties
}
