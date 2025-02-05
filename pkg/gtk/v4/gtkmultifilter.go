// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeAnyFilter   = coreglib.Type(C.gtk_any_filter_get_type())
	GTypeEveryFilter = coreglib.Type(C.gtk_every_filter_get_type())
	GTypeMultiFilter = coreglib.Type(C.gtk_multi_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAnyFilter, F: marshalAnyFilter},
		coreglib.TypeMarshaler{T: GTypeEveryFilter, F: marshalEveryFilter},
		coreglib.TypeMarshaler{T: GTypeMultiFilter, F: marshalMultiFilter},
	})
}

// AnyFilter: GtkAnyFilter matches an item when at least one of its filters
// matches.
//
// To add filters to a GtkAnyFilter, use gtk.MultiFilter.Append().
type AnyFilter struct {
	_ [0]func() // equal guard
	MultiFilter
}

var (
	_ MultiFilterer = (*AnyFilter)(nil)
)

func wrapAnyFilter(obj *coreglib.Object) *AnyFilter {
	return &AnyFilter{
		MultiFilter: MultiFilter{
			Filter: Filter{
				Object: obj,
			},
			Object: obj,
			ListModel: gio.ListModel{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalAnyFilter(p uintptr) (interface{}, error) {
	return wrapAnyFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAnyFilter creates a new empty "any" filter.
//
// Use gtk.MultiFilter.Append() to add filters to it.
//
// This filter matches an item if any of the filters added to it matches the
// item. In particular, this means that if no filter has been added to it, the
// filter matches no item.
//
// The function returns the following values:
//
//    - anyFilter: new GtkAnyFilter.
//
func NewAnyFilter() *AnyFilter {
	var _cret *C.GtkAnyFilter // in

	_cret = C.gtk_any_filter_new()

	var _anyFilter *AnyFilter // out

	_anyFilter = wrapAnyFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _anyFilter
}

// EveryFilter: GtkEveryFilter matches an item when each of its filters matches.
//
// To add filters to a GtkEveryFilter, use gtk.MultiFilter.Append().
type EveryFilter struct {
	_ [0]func() // equal guard
	MultiFilter
}

var (
	_ MultiFilterer = (*EveryFilter)(nil)
)

func wrapEveryFilter(obj *coreglib.Object) *EveryFilter {
	return &EveryFilter{
		MultiFilter: MultiFilter{
			Filter: Filter{
				Object: obj,
			},
			Object: obj,
			ListModel: gio.ListModel{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalEveryFilter(p uintptr) (interface{}, error) {
	return wrapEveryFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEveryFilter creates a new empty "every" filter.
//
// Use gtk.MultiFilter.Append() to add filters to it.
//
// This filter matches an item if each of the filters added to it matches the
// item. In particular, this means that if no filter has been added to it, the
// filter matches every item.
//
// The function returns the following values:
//
//    - everyFilter: new GtkEveryFilter.
//
func NewEveryFilter() *EveryFilter {
	var _cret *C.GtkEveryFilter // in

	_cret = C.gtk_every_filter_new()

	var _everyFilter *EveryFilter // out

	_everyFilter = wrapEveryFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _everyFilter
}

// MultiFilter: GtkMultiFilter is the base class for filters that combine
// multiple filters.
type MultiFilter struct {
	_ [0]func() // equal guard
	Filter

	*coreglib.Object
	gio.ListModel
	Buildable
}

var (
	_ coreglib.Objector = (*MultiFilter)(nil)
)

// MultiFilterer describes types inherited from class MultiFilter.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MultiFilterer interface {
	coreglib.Objector
	baseMultiFilter() *MultiFilter
}

var _ MultiFilterer = (*MultiFilter)(nil)

func wrapMultiFilter(obj *coreglib.Object) *MultiFilter {
	return &MultiFilter{
		Filter: Filter{
			Object: obj,
		},
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalMultiFilter(p uintptr) (interface{}, error) {
	return wrapMultiFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *MultiFilter) baseMultiFilter() *MultiFilter {
	return self
}

// BaseMultiFilter returns the underlying base object.
func BaseMultiFilter(obj MultiFilterer) *MultiFilter {
	return obj.baseMultiFilter()
}

// Append adds a filter to self to use for matching.
//
// The function takes the following parameters:
//
//    - filter: new filter to use.
//
func (self *MultiFilter) Append(filter *Filter) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 *C.GtkFilter      // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkFilter)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(filter).Native()))

	C.gtk_multi_filter_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filter)
}

// Remove removes the filter at the given position from the list of filters used
// by self.
//
// If position is larger than the number of filters, nothing happens and the
// function returns.
//
// The function takes the following parameters:
//
//    - position of filter to remove.
//
func (self *MultiFilter) Remove(position uint) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_filter_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
