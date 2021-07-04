// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_center_box_get_type()), F: marshalCenterBox},
	})
}

// CenterBox: `GtkCenterBox` arranges three children in a row, keeping the
// middle child centered as well as possible.
//
// !An example GtkCenterBox (centerbox.png)
//
// To add children to `GtkCenterBox`, use
// [method@Gtk.CenterBox.set_start_widget],
// [method@Gtk.CenterBox.set_center_widget] and
// [method@Gtk.CenterBox.set_end_widget].
//
// The sizing and positioning of children can be influenced with the align and
// expand properties of the children.
//
//
// GtkCenterBox as GtkBuildable
//
// The `GtkCenterBox` implementation of the `GtkBuildable` interface supports
// placing children in the 3 positions by specifying “start”, “center” or “end”
// as the “type” attribute of a <child> element.
//
//
// CSS nodes
//
// `GtkCenterBox` uses a single CSS node with the name “box”,
//
// The first child of the `GtkCenterBox` will be allocated depending on the text
// direction, i.e. in left-to-right layouts it will be allocated on the left and
// in right-to-left layouts on the right.
//
// In vertical orientation, the nodes of the children are arranged from top to
// bottom.
//
//
// Accessibility
//
// `GtkCenterBox` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type CenterBox interface {
	Widget
	Orientable

	BaselinePosition() BaselinePosition

	CenterWidget() Widget

	EndWidget() Widget

	StartWidget() Widget

	SetBaselinePositionCenterBox(position BaselinePosition)

	SetCenterWidgetCenterBox(child Widget)

	SetEndWidgetCenterBox(child Widget)

	SetStartWidgetCenterBox(child Widget)
}

// centerBox implements the CenterBox class.
type centerBox struct {
	Widget
}

// WrapCenterBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapCenterBox(obj *externglib.Object) CenterBox {
	return centerBox{
		Widget: WrapWidget(obj),
	}
}

func marshalCenterBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterBox(obj), nil
}

func NewCenterBox() CenterBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_center_box_new()

	var _centerBox CenterBox // out

	_centerBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CenterBox)

	return _centerBox
}

func (s centerBox) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterBox       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (s centerBox) CenterWidget() Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_box_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerBox) EndWidget() Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_box_get_end_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerBox) StartWidget() Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_box_get_start_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerBox) SetBaselinePositionCenterBox(position BaselinePosition) {
	var _arg0 *C.GtkCenterBox       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_center_box_set_baseline_position(_arg0, _arg1)
}

func (s centerBox) SetCenterWidgetCenterBox(child Widget) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_center_widget(_arg0, _arg1)
}

func (s centerBox) SetEndWidgetCenterBox(child Widget) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_end_widget(_arg0, _arg1)
}

func (s centerBox) SetStartWidgetCenterBox(child Widget) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_start_widget(_arg0, _arg1)
}

func (s centerBox) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s centerBox) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s centerBox) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s centerBox) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s centerBox) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s centerBox) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s centerBox) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b centerBox) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o centerBox) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o centerBox) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
