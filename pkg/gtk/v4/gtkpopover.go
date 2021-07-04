// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
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
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// Popover: `GtkPopover` is a bubble-like context popup.
//
// !An example GtkPopover (popover.png)
//
// It is primarily meant to provide context-dependent information or options.
// Popovers are attached to a parent widget. By default, they point to the whole
// widget area, although this behavior can be changed with
// [method@Gtk.Popover.set_pointing_to].
//
// The position of a popover relative to the widget it is attached to can also
// be changed with [method@Gtk.Popover.set_position]
//
// By default, `GtkPopover` performs a grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Escape key being
// pressed). If no such modal behavior is desired on a popover,
// [method@Gtk.Popover.set_autohide] may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// `GtkPopover` is often used to replace menus. The best was to do this is to
// use the [class@Gtk.PopoverMenu] subclass which supports being populated from
// a `GMenuModel` with [ctor@Gtk.PopoverMenu.new_from_model].
//
// “`xml <section> <attribute name="display-hint">horizontal-buttons</attribute>
// <item> <attribute name="label">Cut</attribute> <attribute
// name="action">app.cut</attribute> <attribute
// name="verb-icon">edit-cut-symbolic</attribute> </item> <item> <attribute
// name="label">Copy</attribute> <attribute name="action">app.copy</attribute>
// <attribute name="verb-icon">edit-copy-symbolic</attribute> </item> <item>
// <attribute name="label">Paste</attribute> <attribute
// name="action">app.paste</attribute> <attribute
// name="verb-icon">edit-paste-symbolic</attribute> </item> </section> “`
//
//
// CSS nodes
//
// “` popover[.menu] ├── arrow ╰── contents.background ╰── <child> “`
//
// The contents child node always gets the .background style class and the
// popover itself gets the .menu style class if the popover is menu-like (i.e.
// `GtkPopoverMenu`).
//
// Particular uses of `GtkPopover`, such as touch selection popups or magnifiers
// in `GtkEntry` or `GtkTextView` get style classes like .touch-selection or
// .magnifier to differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually not have any
// background. The visible part of the popover can have a shadow. To specify it
// in CSS, set the box-shadow of the contents node.
//
// Note that, in order to accomplish appropriate arrow visuals, `GtkPopover`
// uses custom drawing for the arrow node. This makes it possible for the arrow
// to change its shape dynamically, but it also limits the possibilities of
// styling it using CSS. In particular, the arrow gets drawn over the content
// node's border and shadow, so they look like one shape, which means that the
// border width of the content node and the arrow node should be the same. The
// arrow also does not support any border shape other than solid, no
// border-radius, only one border width (border-bottom-width is used) and no
// box-shadow.
type Popover interface {
	Native
	ShortcutManager

	Autohide() bool

	CascadePopdown() bool

	Child() Widget

	HasArrow() bool

	MnemonicsVisible() bool

	Offset() (xOffset int, yOffset int)

	PointingTo() (gdk.Rectangle, bool)

	Position() PositionType

	PopdownPopover()

	PopupPopover()

	PresentPopover()

	SetAutohidePopover(autohide bool)

	SetCascadePopdownPopover(cascadePopdown bool)

	SetChildPopover(child Widget)

	SetDefaultWidgetPopover(widget Widget)

	SetHasArrowPopover(hasArrow bool)

	SetMnemonicsVisiblePopover(mnemonicsVisible bool)

	SetOffsetPopover(xOffset int, yOffset int)

	SetPointingToPopover(rect *gdk.Rectangle)

	SetPositionPopover(position PositionType)
}

// popover implements the Popover class.
type popover struct {
	Widget
}

// WrapPopover wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopover(obj *externglib.Object) Popover {
	return popover{
		Widget: WrapWidget(obj),
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopover(obj), nil
}

func NewPopover() Popover {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_popover_new()

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Popover)

	return _popover
}

func (p popover) Autohide() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_autohide(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) CascadePopdown() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_cascade_popdown(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) Child() Widget {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (p popover) HasArrow() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_has_arrow(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) MnemonicsVisible() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_mnemonics_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) Offset() (xOffset int, yOffset int) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.int         // in
	var _arg2 C.int         // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_get_offset(_arg0, &_arg1, &_arg2)

	var _xOffset int // out
	var _yOffset int // out

	_xOffset = int(_arg1)
	_yOffset = int(_arg2)

	return _xOffset, _yOffset
}

func (p popover) PointingTo() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkPopover  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_pointing_to(_arg0, &_arg1)

	var _rect gdk.Rectangle // out
	var _ok bool            // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rect = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

func (p popover) Position() PositionType {
	var _arg0 *C.GtkPopover     // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_position(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

func (p popover) PopdownPopover() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(_arg0)
}

func (p popover) PopupPopover() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(_arg0)
}

func (p popover) PresentPopover() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_present(_arg0)
}

func (p popover) SetAutohidePopover(autohide bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if autohide {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_autohide(_arg0, _arg1)
}

func (p popover) SetCascadePopdownPopover(cascadePopdown bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if cascadePopdown {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_cascade_popdown(_arg0, _arg1)
}

func (p popover) SetChildPopover(child Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_popover_set_child(_arg0, _arg1)
}

func (p popover) SetDefaultWidgetPopover(widget Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(_arg0, _arg1)
}

func (p popover) SetHasArrowPopover(hasArrow bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if hasArrow {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_has_arrow(_arg0, _arg1)
}

func (p popover) SetMnemonicsVisiblePopover(mnemonicsVisible bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if mnemonicsVisible {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_mnemonics_visible(_arg0, _arg1)
}

func (p popover) SetOffsetPopover(xOffset int, yOffset int) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(xOffset)
	_arg2 = C.int(yOffset)

	C.gtk_popover_set_offset(_arg0, _arg1, _arg2)
}

func (p popover) SetPointingToPopover(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
}

func (p popover) SetPositionPopover(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_popover_set_position(_arg0, _arg1)
}

func (s popover) Renderer() gsk.Renderer {
	return WrapNative(gextras.InternObject(s)).Renderer()
}

func (s popover) Surface() gdk.Surface {
	return WrapNative(gextras.InternObject(s)).Surface()
}

func (s popover) SurfaceTransform() (x float64, y float64) {
	return WrapNative(gextras.InternObject(s)).SurfaceTransform()
}

func (s popover) Realize() {
	WrapNative(gextras.InternObject(s)).Realize()
}

func (s popover) Unrealize() {
	WrapNative(gextras.InternObject(s)).Unrealize()
}

func (s popover) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s popover) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s popover) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s popover) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s popover) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s popover) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s popover) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b popover) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
