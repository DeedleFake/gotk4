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
		{T: externglib.Type(C.gtk_filter_list_model_get_type()), F: marshalFilterListModel},
	})
}

// FilterListModel: `GtkFilterListModel` is a list model that filters the
// elements of the underlying model according to a `GtkFilter`.
//
// It hides some elements from the other model according to criteria given by a
// `GtkFilter`.
//
// The model can be set up to do incremental searching, so that filtering long
// lists doesn't block the UI. See [method@Gtk.FilterListModel.set_incremental]
// for details.
type FilterListModel interface {
	gextras.Objector

	// Filter gets the `GtkFilter` currently set on @self.
	Filter() *FilterClass
	// Incremental returns whether incremental filtering is enabled.
	//
	// See [method@Gtk.FilterListModel.set_incremental].
	Incremental() bool
	// Pending returns the number of items that have not been filtered yet.
	//
	// You can use this value to check if @self is busy filtering by comparing
	// the return value to 0 or you can compute the percentage of the filter
	// remaining by dividing the return value by the total number of items in
	// the underlying model:
	//
	// “`c pending = gtk_filter_list_model_get_pending (self); model =
	// gtk_filter_list_model_get_model (self); percentage = pending / (double)
	// g_list_model_get_n_items (model); “`
	//
	// If no filter operation is ongoing - in particular when
	// [property@Gtk.FilterListModel:incremental] is false - this function
	// returns 0.
	Pending() uint
	// SetFilter sets the filter used to filter items.
	SetFilter(filter Filter)
	// SetIncremental sets the filter model to do an incremental sort.
	//
	// When incremental filtering is enabled, the `GtkFilterListModel` will not
	// run filters immediately, but will instead queue an idle handler that
	// incrementally filters the items and adds them to the list. This of course
	// means that items are not instantly added to the list, but only appear
	// incrementally.
	//
	// When your filter blocks the UI while filtering, you might consider
	// turning this on. Depending on your model and filters, this may become
	// interesting around 10,000 to 100,000 items.
	//
	// By default, incremental filtering is disabled.
	//
	// See [method@Gtk.FilterListModel.get_pending] for progress information
	// about an ongoing incremental filtering operation.
	SetIncremental(incremental bool)
}

// FilterListModelClass implements the FilterListModel interface.
type FilterListModelClass struct {
	*externglib.Object
}

var _ FilterListModel = (*FilterListModelClass)(nil)

func wrapFilterListModel(obj *externglib.Object) FilterListModel {
	return &FilterListModelClass{
		Object: obj,
	}
}

func marshalFilterListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilterListModel(obj), nil
}

// Filter gets the `GtkFilter` currently set on @self.
func (self *FilterListModelClass) Filter() *FilterClass {
	var _arg0 *C.GtkFilterListModel // out
	var _cret *C.GtkFilter          // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_filter_list_model_get_filter(_arg0)

	var _filter *FilterClass // out

	_filter = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*FilterClass)

	return _filter
}

// Incremental returns whether incremental filtering is enabled.
//
// See [method@Gtk.FilterListModel.set_incremental].
func (self *FilterListModelClass) Incremental() bool {
	var _arg0 *C.GtkFilterListModel // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_filter_list_model_get_incremental(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Pending returns the number of items that have not been filtered yet.
//
// You can use this value to check if @self is busy filtering by comparing the
// return value to 0 or you can compute the percentage of the filter remaining
// by dividing the return value by the total number of items in the underlying
// model:
//
// “`c pending = gtk_filter_list_model_get_pending (self); model =
// gtk_filter_list_model_get_model (self); percentage = pending / (double)
// g_list_model_get_n_items (model); “`
//
// If no filter operation is ongoing - in particular when
// [property@Gtk.FilterListModel:incremental] is false - this function returns
// 0.
func (self *FilterListModelClass) Pending() uint {
	var _arg0 *C.GtkFilterListModel // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_filter_list_model_get_pending(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetFilter sets the filter used to filter items.
func (self *FilterListModelClass) SetFilter(filter Filter) {
	var _arg0 *C.GtkFilterListModel // out
	var _arg1 *C.GtkFilter          // out

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_filter_list_model_set_filter(_arg0, _arg1)
}

// SetIncremental sets the filter model to do an incremental sort.
//
// When incremental filtering is enabled, the `GtkFilterListModel` will not run
// filters immediately, but will instead queue an idle handler that
// incrementally filters the items and adds them to the list. This of course
// means that items are not instantly added to the list, but only appear
// incrementally.
//
// When your filter blocks the UI while filtering, you might consider turning
// this on. Depending on your model and filters, this may become interesting
// around 10,000 to 100,000 items.
//
// By default, incremental filtering is disabled.
//
// See [method@Gtk.FilterListModel.get_pending] for progress information about
// an ongoing incremental filtering operation.
func (self *FilterListModelClass) SetIncremental(incremental bool) {
	var _arg0 *C.GtkFilterListModel // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(self.Native()))
	if incremental {
		_arg1 = C.TRUE
	}

	C.gtk_filter_list_model_set_incremental(_arg0, _arg1)
}
