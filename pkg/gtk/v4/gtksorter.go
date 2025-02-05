// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_Sorter_ConnectChanged(gpointer, GtkSorterChange, guintptr);
// extern GtkSorterOrder _gotk4_gtk4_SorterClass_get_order(GtkSorter*);
// extern GtkOrdering _gotk4_gtk4_SorterClass_compare(GtkSorter*, gpointer, gpointer);
// GtkOrdering _gotk4_gtk4_Sorter_virtual_compare(void* fnptr, GtkSorter* arg0, gpointer arg1, gpointer arg2) {
//   return ((GtkOrdering (*)(GtkSorter*, gpointer, gpointer))(fnptr))(arg0, arg1, arg2);
// };
// GtkSorterOrder _gotk4_gtk4_Sorter_virtual_get_order(void* fnptr, GtkSorter* arg0) {
//   return ((GtkSorterOrder (*)(GtkSorter*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeSorterChange = coreglib.Type(C.gtk_sorter_change_get_type())
	GTypeSorterOrder  = coreglib.Type(C.gtk_sorter_order_get_type())
	GTypeSorter       = coreglib.Type(C.gtk_sorter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSorterChange, F: marshalSorterChange},
		coreglib.TypeMarshaler{T: GTypeSorterOrder, F: marshalSorterOrder},
		coreglib.TypeMarshaler{T: GTypeSorter, F: marshalSorter},
	})
}

// SorterChange describes changes in a sorter in more detail and allows users to
// optimize resorting.
type SorterChange C.gint

const (
	// SorterChangeDifferent: sorter change cannot be described by any of the
	// other enumeration values.
	SorterChangeDifferent SorterChange = iota
	// SorterChangeInverted: sort order was inverted. Comparisons that returned
	// GTK_ORDERING_SMALLER now return GTK_ORDERING_LARGER and vice versa. Other
	// comparisons return the same values as before.
	SorterChangeInverted
	// SorterChangeLessStrict: sorter is less strict: Comparisons may now return
	// GTK_ORDERING_EQUAL that did not do so before.
	SorterChangeLessStrict
	// SorterChangeMoreStrict: sorter is more strict: Comparisons that did
	// return GTK_ORDERING_EQUAL may not do so anymore.
	SorterChangeMoreStrict
)

func marshalSorterChange(p uintptr) (interface{}, error) {
	return SorterChange(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SorterChange.
func (s SorterChange) String() string {
	switch s {
	case SorterChangeDifferent:
		return "Different"
	case SorterChangeInverted:
		return "Inverted"
	case SorterChangeLessStrict:
		return "LessStrict"
	case SorterChangeMoreStrict:
		return "MoreStrict"
	default:
		return fmt.Sprintf("SorterChange(%d)", s)
	}
}

// SorterOrder describes the type of order that a GtkSorter may produce.
type SorterOrder C.gint

const (
	// SorterOrderPartial: partial order. Any Ordering is possible.
	SorterOrderPartial SorterOrder = iota
	// SorterOrderNone: no order, all elements are considered equal.
	// gtk_sorter_compare() will only return GTK_ORDERING_EQUAL.
	SorterOrderNone
	// SorterOrderTotal: total order. gtk_sorter_compare() will only return
	// GTK_ORDERING_EQUAL if an item is compared with itself. Two different
	// items will never cause this value to be returned.
	SorterOrderTotal
)

func marshalSorterOrder(p uintptr) (interface{}, error) {
	return SorterOrder(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SorterOrder.
func (s SorterOrder) String() string {
	switch s {
	case SorterOrderPartial:
		return "Partial"
	case SorterOrderNone:
		return "None"
	case SorterOrderTotal:
		return "Total"
	default:
		return fmt.Sprintf("SorterOrder(%d)", s)
	}
}

// SorterOverrides contains methods that are overridable.
type SorterOverrides struct {
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
	// return value of gtk.Sorter.GetOrder().
	//
	// The function takes the following parameters:
	//
	//    - item1 (optional): first item to compare.
	//    - item2 (optional): second item to compare.
	//
	// The function returns the following values:
	//
	//    - ordering: GTK_ORDERING_EQUAL if item1 == item2, GTK_ORDERING_SMALLER
	//      if item1 < item2, GTK_ORDERING_LARGER if item1 > item2.
	//
	Compare func(item1, item2 *coreglib.Object) Ordering
	// Order gets the order that self conforms to.
	//
	// See gtk.SorterOrder for details of the possible return values.
	//
	// This function is intended to allow optimizations.
	//
	// The function returns the following values:
	//
	//    - sorterOrder: order.
	//
	Order func() SorterOrder
}

func defaultSorterOverrides(v *Sorter) SorterOverrides {
	return SorterOverrides{
		Compare: v.compare,
		Order:   v.order,
	}
}

// Sorter: GtkSorter is an object to describe sorting criteria.
//
//
// Its primary user is gtk.SortListModel
//
// The model will use a sorter to determine the order in which its items should
// appear by calling gtk.Sorter.Compare() for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime. In that
// case, they will emit the gtk.Sorter::changed signal to notify that the sort
// order is no longer valid and should be updated by calling
// gtk_sorter_compare() again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. gtk.ColumnView has built-in support for sorting lists via the
// gtk.ColumnViewColumn:sorter property, where the user can change the sorting
// by clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// GtkSorter and provide one's own sorter.
type Sorter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Sorter)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Sorter, *SorterClass, SorterOverrides](
		GTypeSorter,
		initSorterClass,
		wrapSorter,
		defaultSorterOverrides,
	)
}

func initSorterClass(gclass unsafe.Pointer, overrides SorterOverrides, classInitFunc func(*SorterClass)) {
	pclass := (*C.GtkSorterClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeSorter))))

	if overrides.Compare != nil {
		pclass.compare = (*[0]byte)(C._gotk4_gtk4_SorterClass_compare)
	}

	if overrides.Order != nil {
		pclass.get_order = (*[0]byte)(C._gotk4_gtk4_SorterClass_get_order)
	}

	if classInitFunc != nil {
		class := (*SorterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSorter(obj *coreglib.Object) *Sorter {
	return &Sorter{
		Object: obj,
	}
}

func marshalSorter(p uintptr) (interface{}, error) {
	return wrapSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted whenever the sorter changed.
//
// Users of the sorter should then update the sort order again via
// gtk_sorter_compare().
//
// gtk.SortListModel handles this signal automatically.
//
// Depending on the change parameter, it may be possible to update the sort
// order without a full resorting. Refer to the gtk.SorterChange documentation
// for details.
func (self *Sorter) ConnectChanged(f func(change SorterChange)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "changed", false, unsafe.Pointer(C._gotk4_gtk4_Sorter_ConnectChanged), f)
}

// Changed emits the gtk.Sorter::changed signal to notify all users of the
// sorter that it has changed.
//
// Users of the sorter should then update the sort order via
// gtk_sorter_compare().
//
// Depending on the change parameter, it may be possible to update the sort
// order without a full resorting. Refer to the gtk.SorterChange documentation
// for details.
//
// This function is intended for implementors of GtkSorter subclasses and should
// not be called from other functions.
//
// The function takes the following parameters:
//
//    - change: how the sorter changed.
//
func (self *Sorter) Changed(change SorterChange) {
	var _arg0 *C.GtkSorter      // out
	var _arg1 C.GtkSorterChange // out

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkSorterChange(change)

	C.gtk_sorter_changed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(change)
}

// Compare compares two given items according to the sort order implemented by
// the sorter.
//
// Sorters implement a partial order:
//
// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
// then a ≤ c
//
// The sorter may signal it conforms to additional constraints via the return
// value of gtk.Sorter.GetOrder().
//
// The function takes the following parameters:
//
//    - item1: first item to compare.
//    - item2: second item to compare.
//
// The function returns the following values:
//
//    - ordering: GTK_ORDERING_EQUAL if item1 == item2, GTK_ORDERING_SMALLER if
//      item1 < item2, GTK_ORDERING_LARGER if item1 > item2.
//
func (self *Sorter) Compare(item1, item2 *coreglib.Object) Ordering {
	var _arg0 *C.GtkSorter  // out
	var _arg1 C.gpointer    // out
	var _arg2 C.gpointer    // out
	var _cret C.GtkOrdering // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item1.Native()))
	_arg2 = C.gpointer(unsafe.Pointer(item2.Native()))

	_cret = C.gtk_sorter_compare(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(item1)
	runtime.KeepAlive(item2)

	var _ordering Ordering // out

	_ordering = Ordering(_cret)

	return _ordering
}

// Order gets the order that self conforms to.
//
// See gtk.SorterOrder for details of the possible return values.
//
// This function is intended to allow optimizations.
//
// The function returns the following values:
//
//    - sorterOrder: order.
//
func (self *Sorter) Order() SorterOrder {
	var _arg0 *C.GtkSorter     // out
	var _cret C.GtkSorterOrder // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_sorter_get_order(_arg0)
	runtime.KeepAlive(self)

	var _sorterOrder SorterOrder // out

	_sorterOrder = SorterOrder(_cret)

	return _sorterOrder
}

// Compare compares two given items according to the sort order implemented by
// the sorter.
//
// Sorters implement a partial order:
//
// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
// then a ≤ c
//
// The sorter may signal it conforms to additional constraints via the return
// value of gtk.Sorter.GetOrder().
//
// The function takes the following parameters:
//
//    - item1 (optional): first item to compare.
//    - item2 (optional): second item to compare.
//
// The function returns the following values:
//
//    - ordering: GTK_ORDERING_EQUAL if item1 == item2, GTK_ORDERING_SMALLER if
//      item1 < item2, GTK_ORDERING_LARGER if item1 > item2.
//
func (self *Sorter) compare(item1, item2 *coreglib.Object) Ordering {
	gclass := (*C.GtkSorterClass)(coreglib.PeekParentClass(self))
	fnarg := gclass.compare

	var _arg0 *C.GtkSorter  // out
	var _arg1 C.gpointer    // out
	var _arg2 C.gpointer    // out
	var _cret C.GtkOrdering // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item1.Native()))
	_arg2 = C.gpointer(unsafe.Pointer(item2.Native()))

	_cret = C._gotk4_gtk4_Sorter_virtual_compare(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(item1)
	runtime.KeepAlive(item2)

	var _ordering Ordering // out

	_ordering = Ordering(_cret)

	return _ordering
}

// Order gets the order that self conforms to.
//
// See gtk.SorterOrder for details of the possible return values.
//
// This function is intended to allow optimizations.
//
// The function returns the following values:
//
//    - sorterOrder: order.
//
func (self *Sorter) order() SorterOrder {
	gclass := (*C.GtkSorterClass)(coreglib.PeekParentClass(self))
	fnarg := gclass.get_order

	var _arg0 *C.GtkSorter     // out
	var _cret C.GtkSorterOrder // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C._gotk4_gtk4_Sorter_virtual_get_order(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(self)

	var _sorterOrder SorterOrder // out

	_sorterOrder = SorterOrder(_cret)

	return _sorterOrder
}

// SorterClass: virtual table for GtkSorter.
//
// An instance of this type is always passed by reference.
type SorterClass struct {
	*sorterClass
}

// sorterClass is the struct that's finalized.
type sorterClass struct {
	native *C.GtkSorterClass
}
