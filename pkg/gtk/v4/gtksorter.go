// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_sorter_get_type()), F: marshalSorter},
	})
}

// Sorter: `GtkSorter` is an object to describe sorting criteria.
//
// Its primary user is [class@Gtk.SortListModel]
//
// The model will use a sorter to determine the order in which its items should
// appear by calling [method@Gtk.Sorter.compare] for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime. In that
// case, they will emit the [signal@Gtk.Sorter::changed] signal to notify that
// the sort order is no longer valid and should be updated by calling
// gtk_sorter_compare() again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. [class@Gtk.ColumnView] has built-in support for sorting lists via
// the [property@Gtk.ColumnViewColumn:sorter] property, where the user can
// change the sorting by clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// `GtkSorter` and provide one's own sorter.
type Sorter interface {
	gextras.Objector

	// Changed emits the [signal@Gtk.Sorter::changed] signal to notify all users
	// of the sorter that it has changed.
	//
	// Users of the sorter should then update the sort order via
	// gtk_sorter_compare().
	//
	// Depending on the @change parameter, it may be possible to update the sort
	// order without a full resorting. Refer to the [enum@Gtk.SorterChange]
	// documentation for details.
	//
	// This function is intended for implementors of `GtkSorter` subclasses and
	// should not be called from other functions.
	Changed(change SorterChange)
	// Compare compares two given items according to the sort order implemented
	// by the sorter.
	//
	// Sorters implement a partial order:
	//
	// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
	// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
	// then a ≤ c
	//
	// The sorter may signal it conforms to additional constraints via the
	// return value of [method@Gtk.Sorter.get_order].
	Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering
	// Order gets the order that @self conforms to.
	//
	// See [enum@Gtk.SorterOrder] for details of the possible return values.
	//
	// This function is intended to allow optimizations.
	Order() SorterOrder
}

// sorter implements the Sorter interface.
type sorter struct {
	gextras.Objector
}

var _ Sorter = (*sorter)(nil)

// WrapSorter wraps a GObject to the right type. It is
// primarily used internally.
func WrapSorter(obj *externglib.Object) Sorter {
	return Sorter{
		Objector: obj,
	}
}

func marshalSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSorter(obj), nil
}

// Changed emits the [signal@Gtk.Sorter::changed] signal to notify all users
// of the sorter that it has changed.
//
// Users of the sorter should then update the sort order via
// gtk_sorter_compare().
//
// Depending on the @change parameter, it may be possible to update the sort
// order without a full resorting. Refer to the [enum@Gtk.SorterChange]
// documentation for details.
//
// This function is intended for implementors of `GtkSorter` subclasses and
// should not be called from other functions.
func (s sorter) Changed(change SorterChange) {
	var _arg0 *C.GtkSorter
	var _arg1 C.GtkSorterChange

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkSorterChange)(change)

	C.gtk_sorter_changed(_arg0, _arg1)
}

// Compare compares two given items according to the sort order implemented
// by the sorter.
//
// Sorters implement a partial order:
//
// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
// then a ≤ c
//
// The sorter may signal it conforms to additional constraints via the
// return value of [method@Gtk.Sorter.get_order].
func (s sorter) Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering {
	var _arg0 *C.GtkSorter
	var _arg1 C.gpointer
	var _arg2 C.gpointer

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(item1.Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer(item2.Native()))

	var _cret C.GtkOrdering

	cret = C.gtk_sorter_compare(_arg0, _arg1, _arg2)

	var _ordering Ordering

	_ordering = Ordering(_cret)

	return _ordering
}

// Order gets the order that @self conforms to.
//
// See [enum@Gtk.SorterOrder] for details of the possible return values.
//
// This function is intended to allow optimizations.
func (s sorter) Order() SorterOrder {
	var _arg0 *C.GtkSorter

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))

	var _cret C.GtkSorterOrder

	cret = C.gtk_sorter_get_order(_arg0)

	var _sorterOrder SorterOrder

	_sorterOrder = SorterOrder(_cret)

	return _sorterOrder
}
