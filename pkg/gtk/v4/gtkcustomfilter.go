// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_CustomFilterFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GType values.
var (
	GTypeCustomFilter = coreglib.Type(C.gtk_custom_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCustomFilter, F: marshalCustomFilter},
	})
}

// CustomFilterFunc: user function that is called to determine if the item
// should be matched.
//
// If the filter matches the item, this function must return TRUE. If the item
// should be filtered out, FALSE must be returned.
type CustomFilterFunc func(item *coreglib.Object) (ok bool)

//export _gotk4_gtk4_CustomFilterFunc
func _gotk4_gtk4_CustomFilterFunc(arg1 C.gpointer, arg2 C.gpointer) (cret C.gboolean) {
	var fn CustomFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CustomFilterFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	ok := fn(_item)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CustomFilterOverrider contains methods that are overridable.
type CustomFilterOverrider interface {
}

// CustomFilter: GtkCustomFilter determines whether to include items with a
// callback.
type CustomFilter struct {
	_ [0]func() // equal guard
	Filter
}

var (
	_ coreglib.Objector = (*CustomFilter)(nil)
)

func classInitCustomFilterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCustomFilter(obj *coreglib.Object) *CustomFilter {
	return &CustomFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalCustomFilter(p uintptr) (interface{}, error) {
	return wrapCustomFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCustomFilter creates a new filter using the given match_func to filter
// items.
//
// If match_func is NULL, the filter matches all items.
//
// If the filter func changes its filtering behavior, gtk_filter_changed() needs
// to be called.
//
// The function takes the following parameters:
//
//    - matchFunc (optional): function to filter items.
//
// The function returns the following values:
//
//    - customFilter: new GtkCustomFilter.
//
func NewCustomFilter(matchFunc CustomFilterFunc) *CustomFilter {
	var _arg1 C.GtkCustomFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret *C.GtkCustomFilter // in

	if matchFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(matchFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_custom_filter_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(matchFunc)

	var _customFilter *CustomFilter // out

	_customFilter = wrapCustomFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _customFilter
}

// SetFilterFunc sets the function used for filtering items.
//
// If match_func is NULL, the filter matches all items.
//
// If the filter func changes its filtering behavior, gtk_filter_changed() needs
// to be called.
//
// If a previous function was set, its user_destroy will be called now.
//
// The function takes the following parameters:
//
//    - matchFunc (optional): function to filter items.
//
func (self *CustomFilter) SetFilterFunc(matchFunc CustomFilterFunc) {
	var _arg0 *C.GtkCustomFilter    // out
	var _arg1 C.GtkCustomFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkCustomFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if matchFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(matchFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_custom_filter_set_filter_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(matchFunc)
}
