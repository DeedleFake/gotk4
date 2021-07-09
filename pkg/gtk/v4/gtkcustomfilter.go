// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.gtk_custom_filter_get_type()), F: marshalCustomFilter},
	})
}

// CustomFilterFunc: user function that is called to determine if the @item
// should be matched.
//
// If the filter matches the item, this function must return true. If the item
// should be filtered out, false must be returned.
type CustomFilterFunc func(item *externglib.Object, userData interface{}) (ok bool)

//export gotk4_CustomFilterFunc
func gotk4_CustomFilterFunc(arg0 C.gpointer, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData interface{}    // out

	item = gextras.CastObject(
		externglib.Take(unsafe.Pointer(&arg0))).(*externglib.Object)
	userData = box.Get(uintptr(arg1))

	fn := v.(CustomFilterFunc)
	ok := fn(item, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CustomFilter: `GtkCustomFilter` determines whether to include items with a
// callback.
type CustomFilter interface {
	gextras.Objector

	privateCustomFilterClass()
}

// CustomFilterClass implements the CustomFilter interface.
type CustomFilterClass struct {
	FilterClass
}

var _ CustomFilter = (*CustomFilterClass)(nil)

func wrapCustomFilter(obj *externglib.Object) CustomFilter {
	return &CustomFilterClass{
		FilterClass: FilterClass{
			Object: obj,
		},
	}
}

func marshalCustomFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCustomFilter(obj), nil
}

func (*CustomFilterClass) privateCustomFilterClass() {}
