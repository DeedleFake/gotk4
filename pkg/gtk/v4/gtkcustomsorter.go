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
		{T: externglib.Type(C.gtk_custom_sorter_get_type()), F: marshalCustomSorter},
	})
}

// CustomSorter: `GtkCustomSorter` is a `GtkSorter` implementation that sorts
// via a callback function.
type CustomSorter interface {
	Sorter

	// AsSorter casts the class to the Sorter interface.
	AsSorter() Sorter

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
	//
	// This method is inherited from Sorter
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
	//
	// This method is inherited from Sorter
	Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering
	// GetOrder gets the order that @self conforms to.
	//
	// See [enum@Gtk.SorterOrder] for details of the possible return values.
	//
	// This function is intended to allow optimizations.
	//
	// This method is inherited from Sorter
	GetOrder() SorterOrder
}

// customSorter implements the CustomSorter interface.
type customSorter struct {
	*externglib.Object
}

var _ CustomSorter = (*customSorter)(nil)

// WrapCustomSorter wraps a GObject to a type that implements
// interface CustomSorter. It is primarily used internally.
func WrapCustomSorter(obj *externglib.Object) CustomSorter {
	return customSorter{obj}
}

func marshalCustomSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCustomSorter(obj), nil
}

func (c customSorter) AsSorter() Sorter {
	return WrapSorter(gextras.InternObject(c))
}

func (s customSorter) Changed(change SorterChange) {
	WrapSorter(gextras.InternObject(s)).Changed(change)
}

func (s customSorter) Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering {
	return WrapSorter(gextras.InternObject(s)).Compare(item1, item2)
}

func (s customSorter) GetOrder() SorterOrder {
	return WrapSorter(gextras.InternObject(s)).GetOrder()
}
