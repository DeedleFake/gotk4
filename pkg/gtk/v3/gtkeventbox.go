// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_event_box_get_type()), F: marshalEventBox},
	})
}

// EventBox: the EventBox widget is a subclass of Bin which also has its own
// window. It is useful since it allows you to catch events for widgets which do
// not have their own window.
type EventBox interface {
	Bin

	AboveChild() bool

	VisibleWindow() bool

	SetAboveChildEventBox(aboveChild bool)

	SetVisibleWindowEventBox(visibleWindow bool)
}

// eventBox implements the EventBox class.
type eventBox struct {
	Bin
}

// WrapEventBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventBox(obj *externglib.Object) EventBox {
	return eventBox{
		Bin: WrapBin(obj),
	}
}

func marshalEventBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventBox(obj), nil
}

func NewEventBox() EventBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_event_box_new()

	var _eventBox EventBox // out

	_eventBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EventBox)

	return _eventBox
}

func (e eventBox) AboveChild() bool {
	var _arg0 *C.GtkEventBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_event_box_get_above_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e eventBox) VisibleWindow() bool {
	var _arg0 *C.GtkEventBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_event_box_get_visible_window(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e eventBox) SetAboveChildEventBox(aboveChild bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))
	if aboveChild {
		_arg1 = C.TRUE
	}

	C.gtk_event_box_set_above_child(_arg0, _arg1)
}

func (e eventBox) SetVisibleWindowEventBox(visibleWindow bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))
	if visibleWindow {
		_arg1 = C.TRUE
	}

	C.gtk_event_box_set_visible_window(_arg0, _arg1)
}

func (b eventBox) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b eventBox) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b eventBox) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b eventBox) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b eventBox) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b eventBox) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b eventBox) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b eventBox) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b eventBox) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b eventBox) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
