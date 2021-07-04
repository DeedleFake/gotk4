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
		{T: externglib.Type(C.gtk_box_layout_get_type()), F: marshalBoxLayout},
	})
}

// BoxLayout: `GtkBoxLayout` is a layout manager that arranges children in a
// single row or column.
//
// Whether it is a row or column depends on the value of its
// [property@Gtk.Orientable:orientation] property. Within the other dimension
// all children all allocated the same size. The `GtkBoxLayout` will respect the
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties of
// each child widget.
//
// If you want all children to be assigned the same size, you can use the
// [property@Gtk.BoxLayout:homogeneous] property.
//
// If you want to specify the amount of space placed between each child, you can
// use the [property@Gtk.BoxLayout:spacing] property.
type BoxLayout interface {
	LayoutManager
	Orientable

	BaselinePosition() BaselinePosition

	Homogeneous() bool

	Spacing() uint

	SetBaselinePositionBoxLayout(position BaselinePosition)

	SetHomogeneousBoxLayout(homogeneous bool)

	SetSpacingBoxLayout(spacing uint)
}

// boxLayout implements the BoxLayout class.
type boxLayout struct {
	LayoutManager
}

// WrapBoxLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapBoxLayout(obj *externglib.Object) BoxLayout {
	return boxLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalBoxLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBoxLayout(obj), nil
}

func NewBoxLayout(orientation Orientation) BoxLayout {
	var _arg1 C.GtkOrientation    // out
	var _cret *C.GtkLayoutManager // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_box_layout_new(_arg1)

	var _boxLayout BoxLayout // out

	_boxLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(BoxLayout)

	return _boxLayout
}

func (b boxLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBoxLayout       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (b boxLayout) Homogeneous() bool {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b boxLayout) Spacing() uint {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b boxLayout) SetBaselinePositionBoxLayout(position BaselinePosition) {
	var _arg0 *C.GtkBoxLayout       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_layout_set_baseline_position(_arg0, _arg1)
}

func (b boxLayout) SetHomogeneousBoxLayout(homogeneous bool) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_layout_set_homogeneous(_arg0, _arg1)
}

func (b boxLayout) SetSpacingBoxLayout(spacing uint) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_box_layout_set_spacing(_arg0, _arg1)
}

func (o boxLayout) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o boxLayout) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
