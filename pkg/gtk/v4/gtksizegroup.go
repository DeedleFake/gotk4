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
		{T: externglib.Type(C.gtk_size_group_get_type()), F: marshalSizeGroup},
	})
}

// SizeGroup: `GtkSizeGroup` groups widgets together so they all request the
// same size.
//
// This is typically useful when you want a column of widgets to have the same
// size, but you can’t use a `GtkGrid`.
//
// In detail, the size requested for each widget in a `GtkSizeGroup` is the
// maximum of the sizes that would have been requested for each widget in the
// size group if they were not in the size group. The mode of the size group
// (see [method@Gtk.SizeGroup.set_mode]) determines whether this applies to the
// horizontal size, the vertical size, or both sizes.
//
// Note that size groups only affect the amount of space requested, not the size
// that the widgets finally receive. If you want the widgets in a `GtkSizeGroup`
// to actually be the same size, you need to pack them in such a way that they
// get the size they request and not more.
//
// `GtkSizeGroup` objects are referenced by each widget in the size group, so
// once you have added all widgets to a `GtkSizeGroup`, you can drop the initial
// reference to the size group with g_object_unref(). If the widgets in the size
// group are subsequently destroyed, then they will be removed from the size
// group and drop their references on the size group; when all widgets have been
// removed, the size group will be freed.
//
// Widgets can be part of multiple size groups; GTK will compute the horizontal
// size of a widget from the horizontal requisition of all widgets that can be
// reached from the widget by a chain of size groups of type
// GTK_SIZE_GROUP_HORIZONTAL or GTK_SIZE_GROUP_BOTH, and the vertical size from
// the vertical requisition of all widgets that can be reached from the widget
// by a chain of size groups of type GTK_SIZE_GROUP_VERTICAL or
// GTK_SIZE_GROUP_BOTH.
//
// Note that only non-contextual sizes of every widget are ever consulted by
// size groups (since size groups have no knowledge of what size a widget will
// be allocated in one dimension, it cannot derive how much height a widget will
// receive for a given width). When grouping widgets that trade height for width
// in mode GTK_SIZE_GROUP_VERTICAL or GTK_SIZE_GROUP_BOTH: the height for the
// minimum width will be the requested height for all widgets in the group. The
// same is of course true when horizontally grouping width for height widgets.
//
// Widgets that trade height-for-width should set a reasonably large minimum
// width by way of [property@Gtk.Label:width-chars] for instance. Widgets with
// static sizes as well as widgets that grow (such as ellipsizing text) need no
// such considerations.
//
//
// GtkSizeGroup as GtkBuildable
//
// Size groups can be specified in a UI definition by placing an <object>
// element with `class="GtkSizeGroup"` somewhere in the UI definition. The
// widgets that belong to the size group are specified by a <widgets> element
// that may contain multiple <widget> elements, one for each member of the size
// group. The ”name” attribute gives the id of the widget.
//
// An example of a UI definition fragment with `GtkSizeGroup`: “`xml <object
// class="GtkSizeGroup"> <property name="mode">horizontal</property> <widgets>
// <widget name="radio1"/> <widget name="radio2"/> </widgets> </object> “`
type SizeGroup interface {
	gextras.Objector

	// AddWidget adds a widget to a `GtkSizeGroup`.
	//
	// In the future, the requisition of the widget will be determined as the
	// maximum of its requisition and the requisition of the other widgets in
	// the size group. Whether this applies horizontally, vertically, or in both
	// directions depends on the mode of the size group. See
	// [method@Gtk.SizeGroup.set_mode].
	//
	// When the widget is destroyed or no longer referenced elsewhere, it will
	// be removed from the size group.
	AddWidget(widget Widget)
	// Mode gets the current mode of the size group.
	Mode() SizeGroupMode
	// RemoveWidget removes a widget from a `GtkSizeGroup`.
	RemoveWidget(widget Widget)
}

// SizeGroupClass implements the SizeGroup interface.
type SizeGroupClass struct {
	*externglib.Object
	BuildableInterface
}

var _ SizeGroup = (*SizeGroupClass)(nil)

func wrapSizeGroup(obj *externglib.Object) SizeGroup {
	return &SizeGroupClass{
		Object: obj,
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSizeGroup(obj), nil
}

// AddWidget adds a widget to a `GtkSizeGroup`.
//
// In the future, the requisition of the widget will be determined as the
// maximum of its requisition and the requisition of the other widgets in the
// size group. Whether this applies horizontally, vertically, or in both
// directions depends on the mode of the size group. See
// [method@Gtk.SizeGroup.set_mode].
//
// When the widget is destroyed or no longer referenced elsewhere, it will be
// removed from the size group.
func (s *SizeGroupClass) AddWidget(widget Widget) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer((&SizeGroup).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_size_group_add_widget(_arg0, _arg1)
}

// Mode gets the current mode of the size group.
func (s *SizeGroupClass) Mode() SizeGroupMode {
	var _arg0 *C.GtkSizeGroup    // out
	var _cret C.GtkSizeGroupMode // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer((&SizeGroup).Native()))

	_cret = C.gtk_size_group_get_mode(_arg0)

	var _sizeGroupMode SizeGroupMode // out

	_sizeGroupMode = (SizeGroupMode)(C.GtkSizeGroupMode)

	return _sizeGroupMode
}

// RemoveWidget removes a widget from a `GtkSizeGroup`.
func (s *SizeGroupClass) RemoveWidget(widget Widget) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer((&SizeGroup).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_size_group_remove_widget(_arg0, _arg1)
}
