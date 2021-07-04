// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewport},
	})
}

// Viewport: the Viewport widget acts as an adaptor class, implementing
// scrollability for child widgets that lack their own scrolling capabilities.
// Use GtkViewport to scroll child widgets such as Grid, Box, and so on.
//
// If a widget has native scrolling abilities, such as TextView, TreeView or
// IconView, it can be added to a ScrolledWindow with gtk_container_add(). If a
// widget does not, you must first add the widget to a Viewport, then add the
// viewport to the scrolled window. gtk_container_add() does this automatically
// if a child that does not implement Scrollable is added to a ScrolledWindow,
// so you can ignore the presence of the viewport.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
type Viewport interface {
	Bin
	Scrollable

	BinWindow() gdk.Window

	GetHAdjustment() Adjustment

	ShadowType() ShadowType

	GetVAdjustment() Adjustment

	ViewWindow() gdk.Window

	SetHAdjustmentViewport(adjustment Adjustment)

	SetShadowTypeViewport(typ ShadowType)

	SetVAdjustmentViewport(adjustment Adjustment)
}

// viewport implements the Viewport class.
type viewport struct {
	Bin
}

// WrapViewport wraps a GObject to the right type. It is
// primarily used internally.
func WrapViewport(obj *externglib.Object) Viewport {
	return viewport{
		Bin: WrapBin(obj),
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapViewport(obj), nil
}

func NewViewport(hadjustment Adjustment, vadjustment Adjustment) Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport Viewport // out

	_viewport = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Viewport)

	return _viewport
}

func (v viewport) BinWindow() gdk.Window {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_bin_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (v viewport) GetHAdjustment() Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (v viewport) ShadowType() ShadowType {
	var _arg0 *C.GtkViewport  // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

func (v viewport) GetVAdjustment() Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (v viewport) ViewWindow() gdk.Window {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_view_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (v viewport) SetHAdjustmentViewport(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_hadjustment(_arg0, _arg1)
}

func (v viewport) SetShadowTypeViewport(typ ShadowType) {
	var _arg0 *C.GtkViewport  // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_viewport_set_shadow_type(_arg0, _arg1)
}

func (v viewport) SetVAdjustmentViewport(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_vadjustment(_arg0, _arg1)
}

func (b viewport) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b viewport) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b viewport) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b viewport) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b viewport) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b viewport) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b viewport) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b viewport) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b viewport) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b viewport) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (s viewport) Border() (Border, bool) {
	return WrapScrollable(gextras.InternObject(s)).Border()
}

func (s viewport) HAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).HAdjustment()
}

func (s viewport) HScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).HScrollPolicy()
}

func (s viewport) VAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).VAdjustment()
}

func (s viewport) VScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).VScrollPolicy()
}

func (s viewport) SetHAdjustment(hadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetHAdjustment(hadjustment)
}

func (s viewport) SetHScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetHScrollPolicy(policy)
}

func (s viewport) SetVAdjustment(vadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetVAdjustment(vadjustment)
}

func (s viewport) SetVScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetVScrollPolicy(policy)
}
