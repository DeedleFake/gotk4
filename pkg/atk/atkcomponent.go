// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_scroll_type_get_type()), F: marshalScrollType},
		{T: externglib.Type(C.atk_component_get_type()), F: marshalComponent},
		{T: externglib.Type(C.atk_rectangle_get_type()), F: marshalRectangle},
	})
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType int

const (
	// TopLeft: scroll the object vertically and horizontally to bring its top
	// left corner to the top left corner of the window.
	ScrollTypeTopLeft ScrollType = iota
	// BottomRight: scroll the object vertically and horizontally to bring its
	// bottom right corner to the bottom right corner of the window.
	ScrollTypeBottomRight
	// TopEdge: scroll the object vertically to bring its top edge to the top
	// edge of the window.
	ScrollTypeTopEdge
	// BottomEdge: scroll the object vertically to bring its bottom edge to the
	// bottom edge of the window.
	ScrollTypeBottomEdge
	// LeftEdge: scroll the object vertically and horizontally to bring its left
	// edge to the left edge of the window.
	ScrollTypeLeftEdge
	// RightEdge: scroll the object vertically and horizontally to bring its
	// right edge to the right edge of the window.
	ScrollTypeRightEdge
	// Anywhere: scroll the object vertically and horizontally so that as much
	// as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollTypeAnywhere
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ComponentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ComponentOverrider interface {
	BoundsChanged(bounds *Rectangle)
	// Alpha returns the alpha value (i.e. the opacity) for this @component, on
	// a scale from 0 (fully transparent) to 1.0 (fully opaque).
	Alpha() float64
	// Layer gets the layer of the component.
	Layer() Layer
	// MDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	MDIZOrder() int
	// Size gets the size of the @component in terms of width and height.
	//
	// If the size can not be obtained (e.g. a non-embedded plug or missing
	// support), width and height are set to -1.
	//
	// Deprecated: Since 2.12. Use atk_component_get_extents() instead.
	Size() (width int, height int)
	// GrabFocus grabs focus for this @component.
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by @handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: If you need to track when an object gains or lose the focus,
	// use the Object::state-change "focused" notification instead.
	RemoveFocusHandler(handlerId uint)
	// SetSize: set the size of the @component in terms of width and height.
	SetSize(width int, height int) bool
}

// Component should be implemented by most if not all UI elements with an actual
// on-screen presence, i.e. components which can be said to have a
// screen-coordinate bounding box. Virtually all widgets will need to have
// Component implementations provided for their corresponding Object class. In
// short, only UI elements which are *not* GUI elements will omit this ATK
// interface.
//
// A possible exception might be textual information with a transparent
// background, in which case text glyph bounding box information is provided by
// Text.
type Component interface {
	gextras.Objector

	// Alpha returns the alpha value (i.e. the opacity) for this @component, on
	// a scale from 0 (fully transparent) to 1.0 (fully opaque).
	Alpha() float64
	// Layer gets the layer of the component.
	Layer() Layer
	// MDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	MDIZOrder() int
	// Size gets the size of the @component in terms of width and height.
	//
	// If the size can not be obtained (e.g. a non-embedded plug or missing
	// support), width and height are set to -1.
	//
	// Deprecated: Since 2.12. Use atk_component_get_extents() instead.
	Size() (width int, height int)
	// GrabFocus grabs focus for this @component.
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by @handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: If you need to track when an object gains or lose the focus,
	// use the Object::state-change "focused" notification instead.
	RemoveFocusHandler(handlerId uint)
	// SetSize: set the size of the @component in terms of width and height.
	SetSize(width int, height int) bool
}

// ComponentIface implements the Component interface.
type ComponentIface struct {
	*externglib.Object
}

var _ Component = (*ComponentIface)(nil)

func wrapComponent(obj *externglib.Object) Component {
	return &ComponentIface{
		Object: obj,
	}
}

func marshalComponent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComponent(obj), nil
}

// Alpha returns the alpha value (i.e. the opacity) for this @component, on a
// scale from 0 (fully transparent) to 1.0 (fully opaque).
func (component *ComponentIface) Alpha() float64 {
	var _arg0 *C.AtkComponent // out
	var _cret C.gdouble       // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))

	_cret = C.atk_component_get_alpha(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Layer gets the layer of the component.
func (component *ComponentIface) Layer() Layer {
	var _arg0 *C.AtkComponent // out
	var _cret C.AtkLayer      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))

	_cret = C.atk_component_get_layer(_arg0)

	var _layer Layer // out

	_layer = (Layer)(_cret)

	return _layer
}

// MDIZOrder gets the zorder of the component. The value G_MININT will be
// returned if the layer of the component is not ATK_LAYER_MDI or
// ATK_LAYER_WINDOW.
func (component *ComponentIface) MDIZOrder() int {
	var _arg0 *C.AtkComponent // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))

	_cret = C.atk_component_get_mdi_zorder(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Size gets the size of the @component in terms of width and height.
//
// If the size can not be obtained (e.g. a non-embedded plug or missing
// support), width and height are set to -1.
//
// Deprecated: Since 2.12. Use atk_component_get_extents() instead.
func (component *ComponentIface) Size() (width int, height int) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))

	C.atk_component_get_size(_arg0, &_arg1, &_arg2)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// GrabFocus grabs focus for this @component.
func (component *ComponentIface) GrabFocus() bool {
	var _arg0 *C.AtkComponent // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))

	_cret = C.atk_component_grab_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveFocusHandler: remove the handler specified by @handler_id from the list
// of functions to be executed when this object receives focus events (in or
// out).
//
// Deprecated: If you need to track when an object gains or lose the focus, use
// the Object::state-change "focused" notification instead.
func (component *ComponentIface) RemoveFocusHandler(handlerId uint) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.guint         // out

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))
	_arg1 = C.guint(handlerId)

	C.atk_component_remove_focus_handler(_arg0, _arg1)
}

// SetSize: set the size of the @component in terms of width and height.
func (component *ComponentIface) SetSize(width int, height int) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(component.Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	_cret = C.atk_component_set_size(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Rectangle: data structure for holding a rectangle. Those coordinates are
// relative to the component top-level parent.
type Rectangle struct {
	native C.AtkRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Rectangle)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
