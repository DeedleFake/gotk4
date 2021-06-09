// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_popup_get_type()), F: marshalPopup},
	})
}

// Popup: a `GdkPopup` is a surface that is attached to another surface.
//
// The `GdkPopup` is positioned relative to its parent surface.
//
// `GdkPopup`s are typically used to implement menus and similar popups. They
// can be modal, which is indicated by the [property@GdkPopup:autohide]
// property.
type Popup interface {
	Surface

	// Autohide returns whether this popup is set to hide on outside clicks.
	Autohide() bool
	// Parent returns the parent surface of a popup.
	Parent() Surface
	// PositionX obtains the position of the popup relative to its parent.
	PositionX() int
	// PositionY obtains the position of the popup relative to its parent.
	PositionY() int
	// RectAnchor gets the current popup rectangle anchor.
	//
	// The value returned may change after calling [method@Gdk.Popup.present],
	// or after the [signal@Gdk.Surface::layout] signal is emitted.
	RectAnchor() Gravity
	// SurfaceAnchor gets the current popup surface anchor.
	//
	// The value returned may change after calling [method@Gdk.Popup.present],
	// or after the [signal@Gdk.Surface::layout] signal is emitted.
	SurfaceAnchor() Gravity
	// Present: present @popup after having processed the PopupLayout rules.
	//
	// If the popup was previously now showing, it will be showed, otherwise it
	// will change position according to @layout.
	//
	// After calling this function, the result should be handled in response to
	// the [signal@GdkSurface::layout] signal being emitted. The resulting popup
	// position can be queried using [method@Gdk.Popup.get_position_x],
	// [method@Gdk.Popup.get_position_y], and the resulting size will be sent as
	// parameters in the layout signal. Use [method@Gdk.Popup.get_rect_anchor]
	// and [method@Gdk.Popup.get_surface_anchor] to get the resulting anchors.
	//
	// Presenting may fail, for example if the @popup is set to autohide and is
	// immediately hidden upon being presented. If presenting failed, the
	// [signal@Gdk.Surface::layout] signal will not me emitted.
	Present(width int, height int, layout *PopupLayout) bool
}

// popup implements the Popup interface.
type popup struct {
	Surface
}

var _ Popup = (*popup)(nil)

// WrapPopup wraps a GObject to a type that implements interface
// Popup. It is primarily used internally.
func WrapPopup(obj *externglib.Object) Popup {
	return Popup{
		Surface: WrapSurface(obj),
	}
}

func marshalPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopup(obj), nil
}

// Autohide returns whether this popup is set to hide on outside clicks.
func (p popup) Autohide() bool {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean

	cret = C.gdk_popup_get_autohide(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Parent returns the parent surface of a popup.
func (p popup) Parent() Surface {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret *C.GdkSurface

	cret = C.gdk_popup_get_parent(_arg0)

	var _surface Surface

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Surface)

	return _surface
}

// PositionX obtains the position of the popup relative to its parent.
func (p popup) PositionX() int {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret C.int

	cret = C.gdk_popup_get_position_x(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// PositionY obtains the position of the popup relative to its parent.
func (p popup) PositionY() int {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret C.int

	cret = C.gdk_popup_get_position_y(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// RectAnchor gets the current popup rectangle anchor.
//
// The value returned may change after calling [method@Gdk.Popup.present],
// or after the [signal@Gdk.Surface::layout] signal is emitted.
func (p popup) RectAnchor() Gravity {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret C.GdkGravity

	cret = C.gdk_popup_get_rect_anchor(_arg0)

	var _gravity Gravity

	_gravity = Gravity(_cret)

	return _gravity
}

// SurfaceAnchor gets the current popup surface anchor.
//
// The value returned may change after calling [method@Gdk.Popup.present],
// or after the [signal@Gdk.Surface::layout] signal is emitted.
func (p popup) SurfaceAnchor() Gravity {
	var _arg0 *C.GdkPopup

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var _cret C.GdkGravity

	cret = C.gdk_popup_get_surface_anchor(_arg0)

	var _gravity Gravity

	_gravity = Gravity(_cret)

	return _gravity
}

// Present: present @popup after having processed the PopupLayout rules.
//
// If the popup was previously now showing, it will be showed, otherwise it
// will change position according to @layout.
//
// After calling this function, the result should be handled in response to
// the [signal@GdkSurface::layout] signal being emitted. The resulting popup
// position can be queried using [method@Gdk.Popup.get_position_x],
// [method@Gdk.Popup.get_position_y], and the resulting size will be sent as
// parameters in the layout signal. Use [method@Gdk.Popup.get_rect_anchor]
// and [method@Gdk.Popup.get_surface_anchor] to get the resulting anchors.
//
// Presenting may fail, for example if the @popup is set to autohide and is
// immediately hidden upon being presented. If presenting failed, the
// [signal@Gdk.Surface::layout] signal will not me emitted.
func (p popup) Present(width int, height int, layout *PopupLayout) bool {
	var _arg0 *C.GdkPopup
	var _arg1 C.int
	var _arg2 C.int
	var _arg3 *C.GdkPopupLayout

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)
	_arg3 = (*C.GdkPopupLayout)(unsafe.Pointer(layout.Native()))

	var _cret C.gboolean

	cret = C.gdk_popup_present(_arg0, _arg1, _arg2, _arg3)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}
