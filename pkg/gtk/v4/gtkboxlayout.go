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
	gextras.Objector

	// BaselinePosition gets the value set by
	// gtk_box_layout_set_baseline_position().
	BaselinePosition() BaselinePosition
	// Homogeneous returns whether the layout is set to be homogeneous.
	Homogeneous() bool
	// Spacing returns the space that @box_layout puts between children.
	Spacing() uint
	// SetHomogeneous sets whether the box layout will allocate the same size to
	// all children.
	SetHomogeneous(homogeneous bool)
	// SetSpacing sets how much spacing to put between children.
	SetSpacing(spacing uint)
}

// BoxLayoutClass implements the BoxLayout interface.
type BoxLayoutClass struct {
	*externglib.Object
	LayoutManagerClass
	OrientableInterface
}

var _ BoxLayout = (*BoxLayoutClass)(nil)

func wrapBoxLayout(obj *externglib.Object) BoxLayout {
	return &BoxLayoutClass{
		Object: obj,
		LayoutManagerClass: LayoutManagerClass{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalBoxLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBoxLayout(obj), nil
}

// BaselinePosition gets the value set by
// gtk_box_layout_set_baseline_position().
func (b *BoxLayoutClass) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBoxLayout       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer((&BoxLayout).Native()))

	_cret = C.gtk_box_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(C.GtkBaselinePosition)

	return _baselinePosition
}

// Homogeneous returns whether the layout is set to be homogeneous.
func (b *BoxLayoutClass) Homogeneous() bool {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer((&BoxLayout).Native()))

	_cret = C.gtk_box_layout_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing returns the space that @box_layout puts between children.
func (b *BoxLayoutClass) Spacing() uint {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer((&BoxLayout).Native()))

	_cret = C.gtk_box_layout_get_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetHomogeneous sets whether the box layout will allocate the same size to all
// children.
func (b *BoxLayoutClass) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer((&BoxLayout).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_layout_set_homogeneous(_arg0, _arg1)
}

// SetSpacing sets how much spacing to put between children.
func (b *BoxLayoutClass) SetSpacing(spacing uint) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer((&BoxLayout).Native()))
	_arg1 = C.guint(spacing)

	C.gtk_box_layout_set_spacing(_arg0, _arg1)
}
