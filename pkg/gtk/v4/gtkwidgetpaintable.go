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
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_widget_paintable_get_type()), F: marshalWidgetPaintabler},
	})
}

// WidgetPaintable: GtkWidgetPaintable is a GdkPaintable that displays the
// contents of a widget.
//
// GtkWidgetPaintable will also take care of the widget not being in a state
// where it can be drawn (like when it isn't shown) and just draw nothing or
// where it does not have a size (like when it is hidden) and report no size in
// that case.
//
// Of course, GtkWidgetPaintable allows you to monitor widgets for size changes
// by emitting the gdk.Paintable::invalidate-size signal whenever the size of
// the widget changes as well as for visual changes by emitting the
// gdk.Paintable::invalidate-contents signal whenever the widget changes.
//
// You can use a GtkWidgetPaintable everywhere a GdkPaintable is allowed,
// including using it on a GtkPicture (or one of its parents) that it was set on
// itself via gtk_picture_set_paintable(). The paintable will take care of
// recursion when this happens. If you do this however, ensure that the
// gtk.Picture:can-shrink property is set to TRUE or you might end up with an
// infinitely growing widget.
type WidgetPaintable struct {
	*externglib.Object

	gdk.Paintable
}

var _ gextras.Nativer = (*WidgetPaintable)(nil)

func wrapWidgetPaintable(obj *externglib.Object) *WidgetPaintable {
	return &WidgetPaintable{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalWidgetPaintabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidgetPaintable(obj), nil
}

// NewWidgetPaintable creates a new widget paintable observing the given widget.
func NewWidgetPaintable(widget Widgeter) *WidgetPaintable {
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPaintable // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_widget_paintable_new(_arg1)

	var _widgetPaintable *WidgetPaintable // out

	_widgetPaintable = wrapWidgetPaintable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _widgetPaintable
}

// Widget returns the widget that is observed or NULL if none.
func (self *WidgetPaintable) Widget() Widgeter {
	var _arg0 *C.GtkWidgetPaintable // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_widget_paintable_get_widget(_arg0)

	var _widget Widgeter // out

	_widget = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgeter)

	return _widget
}

// SetWidget sets the widget that should be observed.
func (self *WidgetPaintable) SetWidget(widget Widgeter) {
	var _arg0 *C.GtkWidgetPaintable // out
	var _arg1 *C.GtkWidget          // out

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_widget_paintable_set_widget(_arg0, _arg1)
}
